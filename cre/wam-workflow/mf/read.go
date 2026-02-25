package mf

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	mfBinding "cre-workflow/contracts/evm/src/generated/market_factory"
	oracleBinding "cre-workflow/contracts/evm/src/generated/oracle"
)

func GetFactoryName(
	runtime cre.Runtime,
	chainName string,
	factoryAddress common.Address,
) (string, error) {
	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return "", fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contract, err := mfBinding.NewMarketFactory(evmClient, factoryAddress, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create MarketFactory binding: %w", err)
	}

	name, err := contract.Name(runtime, big.NewInt(-2)).Await()
	if err != nil {
		return "", fmt.Errorf("failed to read factory name: %w", err)
	}

	return name, nil
}

func GetOracleAddress(
	runtime cre.Runtime,
	chainName string,
	factoryAddress common.Address,
) (common.Address, error) {
	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contract, err := mfBinding.NewMarketFactory(evmClient, factoryAddress, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create MarketFactory binding: %w", err)
	}

	oracleAddr, err := contract.Oracle(runtime, big.NewInt(-2)).Await()
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to read oracle address: %w", err)
	}

	return oracleAddr, nil
}

func GetRollingEMAWindow(
	runtime cre.Runtime,
	chainName string,
	oracleAddress common.Address,
) ([]*big.Int, error) {
	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return nil, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contract, err := oracleBinding.NewOracle(evmClient, oracleAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Oracle binding: %w", err)
	}

	window, err := contract.RollingEMAWindow(runtime, big.NewInt(-2)).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to read rollingEMAWindow: %w", err)
	}

	return window, nil
}
