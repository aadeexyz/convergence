package lens

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
)

type WorkflowFactoryData struct {
	Factory           common.Address
	Name              string
	Oracle            common.Address
	LatestMarket      common.Address
	CollateralBalance *big.Int
	RollingEMAWindow  []*big.Int
}

type WorkflowData struct {
	CollateralToken common.Address
	Factories       []WorkflowFactoryData
}

const lensABIJSON = `[{
	"type": "function",
	"name": "getWorkflowData",
	"inputs": [{"name": "fof_", "type": "address"}],
	"outputs": [
		{"name": "collateralToken", "type": "address"},
		{
			"name": "factories",
			"type": "tuple[]",
			"components": [
				{"name": "factory", "type": "address"},
				{"name": "name", "type": "string"},
				{"name": "oracle", "type": "address"},
				{"name": "latestMarket", "type": "address"},
				{"name": "collateralBalance", "type": "uint256"},
				{"name": "rollingEMAWindow", "type": "uint256[]"}
			]
		}
	],
	"stateMutability": "view"
}]`

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

	parsed, err := abi.JSON(strings.NewReader(lensABIJSON))
	if err != nil {
		return cre.PromiseFromResult[*WorkflowData](nil, fmt.Errorf("failed to parse Lens ABI: %w", err))
	}

	calldata, err := parsed.Pack("getWorkflowData", fofAddress)
	if err != nil {
		return cre.PromiseFromResult[*WorkflowData](nil, fmt.Errorf("failed to encode getWorkflowData call: %w", err))
	}

	client := &evm.Client{ChainSelector: chainSelector}

	promise := client.CallContract(runtime, &evm.CallContractRequest{
		Call:        &evm.CallMsg{To: lensAddress.Bytes(), Data: calldata},
		BlockNumber: pb.NewBigIntFromInt(big.NewInt(-2)),
	})

	return cre.Then(promise, func(response *evm.CallContractReply) (*WorkflowData, error) {
		return decodeWorkflowData(&parsed, response.Data)
	})
}

func decodeWorkflowData(parsed *abi.ABI, data []byte) (*WorkflowData, error) {
	method := parsed.Methods["getWorkflowData"]
	values, err := method.Outputs.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getWorkflowData output: %w", err)
	}

	if len(values) != 2 {
		return nil, fmt.Errorf("expected 2 outputs, got %d", len(values))
	}

	collateralToken, ok := values[0].(common.Address)
	if !ok {
		return nil, fmt.Errorf("failed to cast collateralToken to address")
	}

	// Use reflection to decode the tuple array since go-ethereum returns
	// an internal anonymous struct type that won't match inline assertions.
	rawSlice := reflect.ValueOf(values[1])
	if rawSlice.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected slice for factories, got %s", rawSlice.Kind())
	}

	factories := make([]WorkflowFactoryData, rawSlice.Len())
	for i := 0; i < rawSlice.Len(); i++ {
		elem := rawSlice.Index(i)
		factories[i] = WorkflowFactoryData{
			Factory:           elem.FieldByName("Factory").Interface().(common.Address),
			Name:              elem.FieldByName("Name").Interface().(string),
			Oracle:            elem.FieldByName("Oracle").Interface().(common.Address),
			LatestMarket:      elem.FieldByName("LatestMarket").Interface().(common.Address),
			CollateralBalance: elem.FieldByName("CollateralBalance").Interface().(*big.Int),
			RollingEMAWindow:  elem.FieldByName("RollingEMAWindow").Interface().([]*big.Int),
		}
	}

	return &WorkflowData{
		CollateralToken: collateralToken,
		Factories:       factories,
	}, nil
}
