//go:build !wasip1

package lens

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// LensMock is a mock implementation of Lens for testing.
type LensMock struct {
	GetWorkflowData func(GetWorkflowDataInput) (GetWorkflowDataOutput, error)
}

// NewLensMock creates a new LensMock for testing.
func NewLensMock(address common.Address, clientMock *evmmock.ClientCapability) *LensMock {
	mock := &LensMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["getWorkflowData"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetWorkflowData == nil {
				return nil, errors.New("getWorkflowData method not mocked")
			}

			values, err := abi.Methods["getWorkflowData"].Inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			result, err := mock.GetWorkflowData(GetWorkflowDataInput{
				Fof: values[0].(common.Address),
			})
			if err != nil {
				return nil, err
			}

			return abi.Methods["getWorkflowData"].Outputs.Pack(result.CollateralToken, result.Factories)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
