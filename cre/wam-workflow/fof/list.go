package fof

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	fomf "cre-workflow/contracts/evm/src/generated/factory_of_market_factories"
)

type FoFConfig struct {
	ContractAddress string `json:"contractAddress"`
	ChainName       string `json:"chainName"`
	LensAddress     string `json:"lensAddress"`
}

func GetCollateralToken(config *FoFConfig, runtime cre.Runtime) (common.Address, error) {
	chainSelector, err := evm.ChainSelectorFromName(config.ChainName)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contractAddress := common.HexToAddress(config.ContractAddress)

	contract, err := fomf.NewFactoryOfMarketFactories(evmClient, contractAddress, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create FoF instance: %w", err)
	}

	token, err := contract.CollateralToken(runtime, big.NewInt(-2)).Await()
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to read collateralToken: %w", err)
	}

	return token, nil
}

func GetMarketFactories(config *FoFConfig, runtime cre.Runtime) ([]common.Address, error) {
	chainSelector, err := evm.ChainSelectorFromName(config.ChainName)
	if err != nil {
		return nil, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contractAddress := common.HexToAddress(config.ContractAddress)

	contract, err := fomf.NewFactoryOfMarketFactories(evmClient, contractAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create FactoryOfMarketFactories instance: %w", err)
	}

	factories, err := contract.MarketFactories(runtime, big.NewInt(-2)).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to read marketFactories: %w", err)
	}

	runtime.Logger().Info("Successfully read marketFactories",
		"contract", config.ContractAddress,
		"count", len(factories),
	)

	return factories, nil
}
