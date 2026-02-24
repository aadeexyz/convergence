//go:build wasip1

package main

import (
	"fmt"
	"log/slog"
	"math/big"

	"cre-workflow/wam-workflow/fof"
	"cre-workflow/wam-workflow/mf"

	fomf "cre-workflow/contracts/evm/src/generated/factory_of_market_factories"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

type ExecutionResult struct {
	Result string
}

// Workflow configuration loaded from the config.json file
type Config struct {
	FoF      fof.FoFConfig `json:"fof"`
	GasLimit uint64        `json:"gasLimit"`
}

// Workflow implementation with a list of capability triggers
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	logTrigger, err := fof.NewMarketFactoryCreatedTrigger(&config.FoF)
	if err != nil {
		return nil, fmt.Errorf("failed to create MarketFactoryCreated trigger: %w", err)
	}

	// Cron trigger: settle all factories every 24 hours
	cronTrigger := cron.Trigger(&cron.Config{Schedule: "0 0 * * *"})

	return cre.Workflow[*Config]{
		cre.Handler(logTrigger, onMarketFactoryCreated),
		cre.Handler(cronTrigger, onSettlementCron),
	}, nil
}

func onMarketFactoryCreated(config *Config, runtime cre.Runtime, payload *bindings.DecodedLog[fomf.MarketFactoryCreatedDecoded]) (*ExecutionResult, error) {
	logger := runtime.Logger()

	factoryAddress := payload.Data.MarketFactory
	factoryName := payload.Data.Name
	factorySymbol := payload.Data.Symbol

	logger.Info("New market factory created",
		"address", factoryAddress.Hex(),
		"name", factoryName,
		"symbol", factorySymbol,
	)

	// Read the collateral token address from FoF
	collateralToken, err := fof.GetCollateralToken(&config.FoF, runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to read collateral token: %w", err)
	}
	logger.Info("Collateral token", "address", collateralToken.Hex())

	// Read the factory's collateral token balance
	balance, err := mf.GetCollateralBalance(runtime, config.FoF.ChainName, collateralToken, factoryAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to read factory balance: %w", err)
	}
	logger.Info("Factory collateral balance", "factory", factoryAddress.Hex(), "balance", balance.String())

	if balance.Sign() == 0 {
		return &ExecutionResult{
			Result: fmt.Sprintf("Skipped: factory %s (%s) has zero collateral balance", factoryName, factoryAddress.Hex()),
		}, nil
	}

	// Mock attention index: 0.5 = 50000000 (8 decimals)
	index := big.NewInt(50_000_000)
	ema := big.NewInt(50_000_000)

	err = mf.SubmitAttentionIndex(runtime, config.FoF.ChainName, factoryAddress, index, ema, config.GasLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to submit attention index: %w", err)
	}

	return &ExecutionResult{
		Result: fmt.Sprintf("Submitted index %s to factory %s (%s) with balance %s", index.String(), factoryName, factoryAddress.Hex(), balance.String()),
	}, nil
}

func onSettlementCron(config *Config, runtime cre.Runtime, trigger *cron.Payload) (*ExecutionResult, error) {
	logger := runtime.Logger()
	scheduledTime := trigger.ScheduledExecutionTime.AsTime()
	logger.Info("Settlement cron fired", "scheduledTime", scheduledTime)

	// Read collateral token address
	collateralToken, err := fof.GetCollateralToken(&config.FoF, runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to read collateral token: %w", err)
	}

	// Read all market factories
	factories, err := fof.GetMarketFactories(&config.FoF, runtime)
	if err != nil {
		return nil, fmt.Errorf("failed to read market factories: %w", err)
	}

	if len(factories) == 0 {
		return &ExecutionResult{Result: "No market factories found"}, nil
	}

	var settled, skipped, failed int

	for _, factoryAddr := range factories {
		log := logger.With("factory", factoryAddr.Hex())

		// Read factory's collateral balance
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

		log.Info("Settling factory", "balance", balance.String())

		// TODO: replace with real attention index from external data source
		index := big.NewInt(50_000_000)
		ema := big.NewInt(50_000_000)

		err = mf.SubmitAttentionIndex(runtime, config.FoF.ChainName, factoryAddr, index, ema, config.GasLimit)
		if err != nil {
			log.Error("Failed to submit index, skipping", "error", err)
			failed++
			continue
		}

		log.Info("Settlement submitted", "index", index.String())
		settled++
	}

	result := fmt.Sprintf("Settlement complete: %d settled, %d skipped (zero balance), %d failed out of %d factories",
		settled, skipped, failed, len(factories))
	logger.Info(result)

	return &ExecutionResult{Result: result}, nil
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
