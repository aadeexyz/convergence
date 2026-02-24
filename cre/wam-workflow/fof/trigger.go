package fof

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	fomf "cre-workflow/contracts/evm/src/generated/factory_of_market_factories"
)

func NewMarketFactoryCreatedTrigger(config *FoFConfig) (cre.Trigger[*evm.Log, *bindings.DecodedLog[fomf.MarketFactoryCreatedDecoded]], error) {
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
		return nil, fmt.Errorf("failed to create FoF binding: %w", err)
	}

	trigger, err := contract.LogTriggerMarketFactoryCreatedLog(
		chainSelector,
		evm.ConfidenceLevel_CONFIDENCE_LEVEL_FINALIZED,
		[]fomf.MarketFactoryCreatedTopics{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create MarketFactoryCreated log trigger: %w", err)
	}

	return trigger, nil
}
