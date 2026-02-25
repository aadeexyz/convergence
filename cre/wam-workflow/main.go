//go:build wasip1

package main

import (
	"fmt"
	"log/slog"
	"math"
	"math/big"

	"cre-workflow/wam-workflow/fof"
	"cre-workflow/wam-workflow/mf"

	protos "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

const oracleDecimals = 8

type ExecutionResult struct {
	Result string
}

// Workflow configuration loaded from the config.json file
type Config struct {
	FoF      fof.FoFConfig `json:"fof"`
	Schedule string        `json:"schedule"`
	GasLimit uint64        `json:"gasLimit"`
}

// Workflow implementation with a list of capability triggers
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	cronTrigger := cron.Trigger(&cron.Config{Schedule: config.Schedule})

	return cre.Workflow[*Config]{
		cre.Handler(cronTrigger, onCronTrigger),
	}, nil
}

func onCronTrigger(config *Config, runtime cre.Runtime, trigger *cron.Payload) (*ExecutionResult, error) {
	logger := runtime.Logger()
	scheduledTime := trigger.ScheduledExecutionTime.AsTime()
	logger.Info("Cron trigger fired", "scheduledTime", scheduledTime)

	// ── 1. Fetch API keys from secrets (non-fatal for simulation) ───
	var googleKey, twitterKey, serpKey string

	googleSecret, err := runtime.GetSecret(&protos.SecretRequest{Id: "GOOGLE_API_KEY"}).Await()
	if err != nil {
		logger.Warn("GOOGLE_API_KEY secret not found, API calls will use fallback", "error", err)
	} else {
		googleKey = googleSecret.Value
	}
	twitterSecret, err := runtime.GetSecret(&protos.SecretRequest{Id: "TWITTER_API_KEY"}).Await()
	if err != nil {
		logger.Warn("TWITTER_API_KEY secret not found, API calls will use fallback", "error", err)
	} else {
		twitterKey = twitterSecret.Value
	}
	serpSecret, err := runtime.GetSecret(&protos.SecretRequest{Id: "SERP_API_KEY"}).Await()
	if err != nil {
		logger.Warn("SERP_API_KEY secret not found, API calls will use fallback", "error", err)
	} else {
		serpKey = serpSecret.Value
	}

	// ── 2. Read on-chain state ──────────────────────────────────────
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

	// ── 3. Process each factory with its own keyword ────────────────
	var settled, skipped, failed int
	scale := math.Pow(10, oracleDecimals)

	for _, factoryAddr := range factories {
		log := logger.With("factory", factoryAddr.Hex())

		// Check collateral balance first (cheap read)
		balance, err := mf.GetCollateralBalance(runtime, config.FoF.ChainName, collateralToken, factoryAddr)
		if err != nil {
			log.Error("Failed to read balance, skipping", "error", err)
			failed++
			continue
		}

		if balance.Sign() == 0 {
			log.Info("Zero balance, skipping")
			skipped++
			continue
		}

		// Read the factory's keyword (its name)
		keyword, err := mf.GetFactoryName(runtime, config.FoF.ChainName, factoryAddr)
		if err != nil {
			log.Error("Failed to read factory name, skipping", "error", err)
			failed++
			continue
		}
		log = log.With("keyword", keyword)

		// Read oracle address and rolling EMA window
		oracleAddr, err := mf.GetOracleAddress(runtime, config.FoF.ChainName, factoryAddr)
		if err != nil {
			log.Error("Failed to read oracle address, skipping", "error", err)
			failed++
			continue
		}

		emaWindow, err := mf.GetRollingEMAWindow(runtime, config.FoF.ChainName, oracleAddr)
		if err != nil {
			log.Error("Failed to read EMA window, skipping", "error", err)
			failed++
			continue
		}

		// Compute attention scores for this keyword
		scores, err := GetRawIndex(keyword, runtime, googleKey, twitterKey, serpKey)
		if err != nil {
			log.Error("Failed to compute attention scores, skipping", "error", err)
			failed++
			continue
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
			// Convert on-chain median (8-decimal uint256) back to float for EMA calc
			prevEMAFloat := float64(medianEMA.Uint64()) / scale
			emaFloat := EMA(normalized, &prevEMAFloat)
			scaledEMA = new(big.Int).SetUint64(uint64(emaFloat * scale))
		} else {
			// No previous data — seed EMA with current index
			scaledEMA = new(big.Int).Set(scaledIndex)
		}

		log.Info("Index computed",
			"rawIndex", rawIndex,
			"sigmoid", normalized,
			"scaledIndex", scaledIndex.String(),
			"scaledEMA", scaledEMA.String(),
			"medianPrevEMA", medianEMA,
		)

		// Submit to chain
		err = mf.SubmitAttentionIndex(runtime, config.FoF.ChainName, factoryAddr, scaledIndex, scaledEMA, config.GasLimit)
		if err != nil {
			log.Error("Failed to submit index, skipping", "error", err)
			failed++
			continue
		}

		log.Info("Index submitted successfully")
		settled++
	}

	result := fmt.Sprintf("%d settled, %d skipped, %d failed out of %d factories",
		settled, skipped, failed, len(factories))
	logger.Info(result)

	return &ExecutionResult{Result: result}, nil
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
