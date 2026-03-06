//go:build wasip1

package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/big"

	"cre-workflow/wam-workflow/fof"
	"cre-workflow/wam-workflow/lens"
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

type Config struct {
	FoF                 fof.FoFConfig `json:"fof"`
	Schedule            string        `json:"schedule"`
	GasLimit            uint64        `json:"gasLimit"`
	TwitterPageDelaySec int           `json:"twitterPageDelaySec"`
}

func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	cronTrigger := cron.Trigger(&cron.Config{Schedule: config.Schedule})

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

// Returns: "settled", "updated", "skipped", or an error.
func processFactoryFromData(
	config *Config,
	runtime cre.Runtime,
	fd lens.WorkflowFactoryData,
	settlement bool,
	googleKey, twitterKey, serpKey string,
) (string, error) {
	log := runtime.Logger().With("factory", fd.Factory.Hex(), "settlement", settlement)
	scale := math.Pow(10, oracleDecimals)

	log.Info("Balance check", "balance", fd.CollateralBalance.String())
	if fd.CollateralBalance.Sign() == 0 {
		log.Info("Zero balance, skipping")
		return "skipped", nil
	}

	keyword := fd.Name
	log = log.With("keyword", keyword)

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

	rawIndex := scores.RawIndex
	normalized := Sigmoid(rawIndex)
	scaledIndex := new(big.Int).SetUint64(uint64(normalized * scale))

	var scaledEMA *big.Int
	medianEMA := MedianNonZero(fd.RollingEMAWindow)
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

		err = mf.SubmitAttentionIndex(runtime, config.FoF.ChainName, fd.Factory, scaledIndex, scaledEMA, config.GasLimit)
		if err != nil {
			return "", fmt.Errorf("failed to submit settlement report: %w", err)
		}
		log.Info("Settlement report submitted")
		return "settled", nil
	}

	err = mf.SubmitOracleIndex(runtime, config.FoF.ChainName, fd.Oracle, scaledIndex, scaledEMA, config.GasLimit)
	if err != nil {
		return "", fmt.Errorf("failed to submit oracle report: %w", err)
	}
	log.Info("Oracle report submitted")
	return "updated", nil
}

func processFactory(
	config *Config,
	runtime cre.Runtime,
	factoryAddr common.Address,
	collateralToken common.Address,
	settlement bool,
	googleKey, twitterKey, serpKey string,
) (string, error) {
	log := runtime.Logger().With("factory", factoryAddr.Hex(), "settlement", settlement)

	latestMarket, err := mf.GetLatestMarket(runtime, config.FoF.ChainName, factoryAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read latest market: %w", err)
	}

	balanceAccount := factoryAddr
	if latestMarket != (common.Address{}) {
		balanceAccount = latestMarket
	}
	balance, err := mf.GetCollateralBalance(runtime, config.FoF.ChainName, collateralToken, balanceAccount)
	if err != nil {
		return "", fmt.Errorf("failed to read balance: %w", err)
	}

	keyword, err := mf.GetFactoryName(runtime, config.FoF.ChainName, factoryAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read factory name: %w", err)
	}

	oracleAddr, err := mf.GetOracleAddress(runtime, config.FoF.ChainName, factoryAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read oracle address: %w", err)
	}

	emaWindow, err := mf.GetRollingEMAWindow(runtime, config.FoF.ChainName, oracleAddr)
	if err != nil {
		return "", fmt.Errorf("failed to read EMA window: %w", err)
	}

	fd := lens.WorkflowFactoryData{
		Factory:           factoryAddr,
		Name:              keyword,
		Oracle:            oracleAddr,
		LatestMarket:      latestMarket,
		CollateralBalance: balance,
		RollingEMAWindow:  emaWindow,
	}

	log.Info("Individual reads complete", "keyword", keyword, "oracle", oracleAddr.Hex())
	return processFactoryFromData(config, runtime, fd, settlement, googleKey, twitterKey, serpKey)
}

func onCronTrigger(config *Config, runtime cre.Runtime, trigger *cron.Payload) (*ExecutionResult, error) {
	logger := runtime.Logger()
	scheduledTime := trigger.ScheduledExecutionTime.AsTime()
	logger.Info("Cron trigger fired", "scheduledTime", scheduledTime)

	googleKey, twitterKey, serpKey := fetchAPIKeys(runtime)

	lensAddr := common.HexToAddress(config.FoF.LensAddress)
	fofAddr := common.HexToAddress(config.FoF.ContractAddress)
	workflowData, err := lens.GetWorkflowData(runtime, config.FoF.ChainName, lensAddr, fofAddr).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to batch read workflow data: %w", err)
	}

	n := len(workflowData.Factories)
	if n == 0 {
		return &ExecutionResult{Result: "No market factories found"}, nil
	}

	logger.Info("Batch read complete",
		"collateralToken", workflowData.CollateralToken.Hex(),
		"factoryCount", n,
	)

	const cronIntervalSec = 30 * 60 // 30 minutes
	slot := scheduledTime.Unix() / cronIntervalSec
	idx := int(slot % int64(n))
	fd := workflowData.Factories[idx]

	logger.Info("Round-robin selection",
		"slot", slot,
		"index", idx,
		"factory", fd.Factory.Hex(),
		"keyword", fd.Name,
	)

	status, err := processFactoryFromData(config, runtime, fd, true, googleKey, twitterKey, serpKey)
	if err != nil {
		return nil, fmt.Errorf("factory %s failed: %w", fd.Factory.Hex(), err)
	}

	result := fmt.Sprintf("Factory %s (%s): %s [slot %d, index %d/%d]",
		fd.Factory.Hex(), fd.Name, status, slot, idx, n)
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
