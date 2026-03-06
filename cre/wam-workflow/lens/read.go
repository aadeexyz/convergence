package lens

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	lensBinding "cre-workflow/contracts/evm/src/generated/lens"
)

type WorkflowFactoryData = lensBinding.WorkflowFactoryData

type WorkflowData struct {
	CollateralToken common.Address
	Factories       []WorkflowFactoryData
}

func GetWorkflowData(
	runtime cre.Runtime,
	chainName string,
	lensAddress common.Address,
	fofAddress common.Address,
) cre.Promise[*WorkflowData] {
	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return cre.PromiseFromResult[*WorkflowData](nil, fmt.Errorf("invalid chain name: %w", err))
	}

	contract, err := lensBinding.NewLens(&evm.Client{ChainSelector: chainSelector}, lensAddress, nil)
	if err != nil {
		return cre.PromiseFromResult[*WorkflowData](nil, fmt.Errorf("failed to create Lens binding: %w", err))
	}

	promise := contract.GetWorkflowData(runtime, lensBinding.GetWorkflowDataInput{
		Fof: fofAddress,
	}, big.NewInt(-2))

	return cre.Then(promise, func(response lensBinding.GetWorkflowDataOutput) (*WorkflowData, error) {
		return &WorkflowData{
			CollateralToken: response.CollateralToken,
			Factories:       response.Factories,
		}, nil
	})
}
