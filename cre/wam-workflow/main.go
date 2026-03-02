//go:build wasip1

package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/big"

	"cre-workflow/wam-workflow/fof"
	"cre-workflow/wam-workflow/mf"

	"github.com/ethereum/go-ethereum/common"
	protos "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"

	fomf "cre-workflow/contracts/evm/src/generated/factory_of_market_factories"
)

const oracleDecimals = 8

type ExecutionResult struct {
	Result string
}

// Workflow configuration loaded from the config.json file
type Config struct {
	FoF                 fof.FoFConfig `json:"fof"`
	Schedule            string        `json:"schedule"`
	GasLimit            uint64        `json:"gasLimit"`
	TwitterPageDelaySec int           `json:"twitterPageDelaySec"`
}

// Workflow implementation with a list of capability triggers
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	cronTrigger := cron.Trigger(&cron.Config{Schedule: config.Schedule})

	// Set up log trigger for MarketFactoryCreated events
	chainSelector, err := evm.ChainSelectorFromName(config.FoF.ChainName)
	if err != nil {
		return nil, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{ChainSelector: chainSelector}
	contractAddress := common.HexToAddress(config.FoF.ContractAddress)
	contract, err := fomf.NewFactoryOfMarketFactories(evmClient, contractAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create FoF instance: %w", err)
	}

	factoryCreatedTrigger, err := contract.LogTriggerMarketFactoryCreatedLog(
		chainSelector, evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED, nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create log trigger: %w", err)
	}

	return cre.Workflow[*Config]{
		cre.Handler(cronTrigger, onCronTrigger),
		cre.Handler(factoryCreatedTrigger, onFactoryCreated),
	}, nil
}

// fetchAPIKeys retrieves API keys from secrets (non-fatal if missing)
func fetchAPIKeys(runtime cre.Runtime) (googleKey, twitterKey, serpKey string) {
	logger := runtime.Logger()

	googleSecret, err := runtime.GetSecret(&protos.SecretRequest{Id: "GOOGLE_API_KEY"}).Await()
	if err != nil {
		logger.Warn("GOOGLE_API_KEY secret not found", "error", err)
	} else {
		googleKey = googleSecret.Value
	}
	twitterSecret, err := runtime.GetSecret(&protos.SecretRequest{Id: "TWITTER_API_KEY"}).Await()
	if err != nil {
		logger.Warn("TWITTER_API_KEY secret not found", "error", err)
	} else {
		twitterKey = twitterSecret.Value
	}
	serpSecret, err := runtime.GetSecret(&protos.SecretRequest{Id: "SERP_API_KEY"}).Await()
	if err != nil {
		logger.Warn("SERP_API_KEY secret not found", "error", err)
	} else {
		serpKey = serpSecret.Value
	}
	return
}

// processFactory computes and submits the attention index for a single factory.
// When settlement is true, the report goes to MarketFactory (settle + seed + oracle update).
// When settlement is false, the report goes directly to the Oracle (oracle-only update).
// Returns: "settled", "updated", "skipped", or an error.
func processFactory(
	config *Config,
	runtime cre.Runtime,
	factoryAddr common.Address,
	collateralToken common.Address,
	settlement bool,
	googleKey, twitterKey, serpKey string,
) (string, error) {
	log := runtime.Logger().With("factory", factoryAddr.Hex(), "settlement", settlement)
	scale := math.Pow(10, oracleDecimals)

	// Check collateral balance: look at the latest market (where collateral lives during active trading)
	latestMarket, err := mf.GetLatestMarket(runtime, config.FoF.ChainName, factoryAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read latest market: %w", err)
	}

	// Check factory balance first, then fall back to latest market balance
	balanceAccount := factoryAddr
	if latestMarket != (common.Address{}) {
		balanceAccount = latestMarket
	}
	balance, err := mf.GetCollateralBalance(runtime, config.FoF.ChainName, collateralToken, balanceAccount)
	if err != nil {
		return "", fmt.Errorf("failed to read balance: %w", err)
	}
	log.Info("Balance check", "account", balanceAccount.Hex(), "balance", balance.String())
	if balance.Sign() == 0 {
		log.Info("Zero balance, skipping")
		return "skipped", nil
	}

	// Read the factory's keyword (its name)
	keyword, err := mf.GetFactoryName(runtime, config.FoF.ChainName, factoryAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read factory name: %w", err)
	}
	log = log.With("keyword", keyword)

	// Read oracle address and rolling EMA window
	oracleAddr, err := mf.GetOracleAddress(runtime, config.FoF.ChainName, factoryAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read oracle address: %w", err)
	}

	emaWindow, err := mf.GetRollingEMAWindow(runtime, config.FoF.ChainName, oracleAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read EMA window: %w", err)
	}

	// Compute attention scores for this keyword
	scores, err := GetRawIndex(keyword, runtime, googleKey, twitterKey, serpKey, config.TwitterPageDelaySec)
	if err != nil {
		return "", fmt.Errorf("failed to compute attention scores: %w", err)
	}

	log.Info("Attention scores computed",
		"youtubeScore", scores.YouTube,
		"twitterScore", scores.Twitter,
		"googleTrendsScore", scores.GoogleTrends,
		"rawIndex", scores.RawIndex,
	)

	// Sigmoid → scale to 8 decimals
	rawIndex := scores.RawIndex
	normalized := Sigmoid(rawIndex)
	scaledIndex := new(big.Int).SetUint64(uint64(normalized * scale))

	// Compute EMA using median of non-zero on-chain values as previous EMA
	var scaledEMA *big.Int
	medianEMA := MedianNonZero(emaWindow)
	if medianEMA != nil {
		prevEMAFloat := float64(medianEMA.Uint64()) / scale
		emaFloat := EMA(normalized, &prevEMAFloat)
		scaledEMA = new(big.Int).SetUint64(uint64(emaFloat * scale))
	} else {
		scaledEMA = new(big.Int).Set(scaledIndex)
	}

	log.Info("Index computed",
		"rawIndex", rawIndex,
		"sigmoid", normalized,
		"scaledIndex", scaledIndex.String(),
		"scaledEMA", scaledEMA.String(),
		"medianPrevEMA", medianEMA,
	)

	if settlement {
		// Push to MarketFactory: oracle update + settle old market + seed new market
		err = mf.SubmitAttentionIndex(runtime, config.FoF.ChainName, factoryAddr, scaledIndex, scaledEMA, config.GasLimit)
		if err != nil {
			return "", fmt.Errorf("failed to submit settlement report: %w", err)
		}
		log.Info("Settlement report submitted")
		return "settled", nil
	}

	// Push directly to Oracle: oracle-only update
	err = mf.SubmitOracleIndex(runtime, config.FoF.ChainName, oracleAddr, scaledIndex, scaledEMA, config.GasLimit)
	if err != nil {
		return "", fmt.Errorf("failed to submit oracle report: %w", err)
	}
	log.Info("Oracle report submitted")
	return "updated", nil
}

func onCronTrigger(config *Config, runtime cre.Runtime, trigger *cron.Payload) (*ExecutionResult, error) {
	logger := runtime.Logger()
	scheduledTime := trigger.ScheduledExecutionTime.AsTime()
	logger.Info("Cron trigger fired", "scheduledTime", scheduledTime)

	googleKey, twitterKey, serpKey := fetchAPIKeys(runtime)

	collateralToken, err := fof.GetCollateralToken(&config.FoF, runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to read collateral token: %w", err)
	}

	factories, err := fof.GetMarketFactories(&config.FoF, runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to read market factories: %w", err)
	}

	if len(factories) == 0 {
		return &ExecutionResult{Result: "No market factories found"}, nil
	}

	// Settle at midnight UTC; oracle-only update otherwise
	settlement := scheduledTime.UTC().Hour() == 0
	settlement = true // testing only
	logger.Info("Settlement check", "utcHour", scheduledTime.UTC().Hour(), "settlement", settlement)

	var settled, updated, skipped, failed int

	for _, factoryAddr := range factories {

		status, err := processFactory(config, runtime, factoryAddr, collateralToken, settlement, googleKey, twitterKey, serpKey)
		if err != nil {
			logger.Error("Factory processing failed, skipping", "factory", factoryAddr.Hex(), "error", err)
			failed++
			continue
		}
		switch status {
		case "settled":
			settled++
		case "updated":
			updated++
		case "skipped":
			skipped++
		}
	}

	result := fmt.Sprintf("%d settled, %d updated, %d skipped, %d failed out of %d factories",
		settled, updated, skipped, failed, len(factories))
	logger.Info(result)

	return &ExecutionResult{Result: result}, nil
}

func onFactoryCreated(config *Config, runtime cre.Runtime, trigger *bindings.DecodedLog[fomf.MarketFactoryCreatedDecoded]) (*ExecutionResult, error) {
	logger := runtime.Logger()
	factoryAddr := trigger.Data.MarketFactory
	logger.Info("New MarketFactory created, computing first index",
		"factory", factoryAddr.Hex(),
		"name", trigger.Data.Name,
		"symbol", trigger.Data.Symbol,
	)

	googleKey, twitterKey, serpKey := fetchAPIKeys(runtime)

	collateralToken, err := fof.GetCollateralToken(&config.FoF, runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to read collateral token: %w", err)
	}

	// New factory always gets a settlement report (first market creation)
	status, err := processFactory(config, runtime, factoryAddr, collateralToken, true, googleKey, twitterKey, serpKey)
	if err != nil {
		return nil, fmt.Errorf("failed to process new factory %s: %w", factoryAddr.Hex(), err)
	}

	result := fmt.Sprintf("New factory %s: %s", factoryAddr.Hex(), status)
	logger.Info(result)

	return &ExecutionResult{Result: result}, nil
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
