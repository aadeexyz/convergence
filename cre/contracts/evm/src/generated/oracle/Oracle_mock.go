// Code generated — DO NOT EDIT.

//go:build !wasip1

package oracle

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

// OracleMock is a mock implementation of Oracle for testing.
type OracleMock struct {
	CurrentRoundId             func() (*big.Int, error)
	Decimals                   func() (uint8, error)
	GetExpectedAuthor          func() (common.Address, error)
	GetExpectedWorkflowId      func() ([32]byte, error)
	GetExpectedWorkflowName    func() ([10]byte, error)
	GetForwarderAddress        func() (common.Address, error)
	GetLatestRound             func() (IOracleRound, error)
	GetRound                   func(GetRoundInput) (IOracleRound, error)
	Keyword                    func() (string, error)
	Owner                      func() (common.Address, error)
	OwnershipHandoverExpiresAt func(OwnershipHandoverExpiresAtInput) (*big.Int, error)
	RollingEMAWindow           func() ([]*big.Int, error)
	SupportsInterface          func(SupportsInterfaceInput) (bool, error)
}

// NewOracleMock creates a new OracleMock for testing.
func NewOracleMock(address common.Address, clientMock *evmmock.ClientCapability) *OracleMock {
	mock := &OracleMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["currentRoundId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.CurrentRoundId == nil {
				return nil, errors.New("currentRoundId method not mocked")
			}
			result, err := mock.CurrentRoundId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["currentRoundId"].Outputs.Pack(result)
		},
		string(abi.Methods["decimals"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Decimals == nil {
				return nil, errors.New("decimals method not mocked")
			}
			result, err := mock.Decimals()
			if err != nil {
				return nil, err
			}
			return abi.Methods["decimals"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedAuthor"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedAuthor == nil {
				return nil, errors.New("getExpectedAuthor method not mocked")
			}
			result, err := mock.GetExpectedAuthor()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedAuthor"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedWorkflowId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedWorkflowId == nil {
				return nil, errors.New("getExpectedWorkflowId method not mocked")
			}
			result, err := mock.GetExpectedWorkflowId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedWorkflowId"].Outputs.Pack(result)
		},
		string(abi.Methods["getExpectedWorkflowName"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetExpectedWorkflowName == nil {
				return nil, errors.New("getExpectedWorkflowName method not mocked")
			}
			result, err := mock.GetExpectedWorkflowName()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getExpectedWorkflowName"].Outputs.Pack(result)
		},
		string(abi.Methods["getForwarderAddress"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetForwarderAddress == nil {
				return nil, errors.New("getForwarderAddress method not mocked")
			}
			result, err := mock.GetForwarderAddress()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getForwarderAddress"].Outputs.Pack(result)
		},
		string(abi.Methods["getLatestRound"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetLatestRound == nil {
				return nil, errors.New("getLatestRound method not mocked")
			}
			result, err := mock.GetLatestRound()
			if err != nil {
				return nil, err
			}
			return abi.Methods["getLatestRound"].Outputs.Pack(result)
		},
		string(abi.Methods["getRound"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetRound == nil {
				return nil, errors.New("getRound method not mocked")
			}
			inputs := abi.Methods["getRound"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetRoundInput{
				Id: values[0].(*big.Int),
			}

			result, err := mock.GetRound(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getRound"].Outputs.Pack(result)
		},
		string(abi.Methods["keyword"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Keyword == nil {
				return nil, errors.New("keyword method not mocked")
			}
			result, err := mock.Keyword()
			if err != nil {
				return nil, err
			}
			return abi.Methods["keyword"].Outputs.Pack(result)
		},
		string(abi.Methods["owner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Owner == nil {
				return nil, errors.New("owner method not mocked")
			}
			result, err := mock.Owner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["owner"].Outputs.Pack(result)
		},
		string(abi.Methods["ownershipHandoverExpiresAt"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.OwnershipHandoverExpiresAt == nil {
				return nil, errors.New("ownershipHandoverExpiresAt method not mocked")
			}
			inputs := abi.Methods["ownershipHandoverExpiresAt"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := OwnershipHandoverExpiresAtInput{
				PendingOwner: values[0].(common.Address),
			}

			result, err := mock.OwnershipHandoverExpiresAt(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["ownershipHandoverExpiresAt"].Outputs.Pack(result)
		},
		string(abi.Methods["rollingEMAWindow"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.RollingEMAWindow == nil {
				return nil, errors.New("rollingEMAWindow method not mocked")
			}
			result, err := mock.RollingEMAWindow()
			if err != nil {
				return nil, err
			}
			return abi.Methods["rollingEMAWindow"].Outputs.Pack(result)
		},
		string(abi.Methods["supportsInterface"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.SupportsInterface == nil {
				return nil, errors.New("supportsInterface method not mocked")
			}
			inputs := abi.Methods["supportsInterface"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := SupportsInterfaceInput{
				InterfaceId: values[0].([4]byte),
			}

			result, err := mock.SupportsInterface(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["supportsInterface"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
