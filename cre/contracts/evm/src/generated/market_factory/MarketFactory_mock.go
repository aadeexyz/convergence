// Code generated — DO NOT EDIT.

//go:build !wasip1

package market_factory

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

// MarketFactoryMock is a mock implementation of MarketFactory for testing.
type MarketFactoryMock struct {
	CollateralToken            func() (common.Address, error)
	Decimals                   func() (uint8, error)
	GetExpectedAuthor          func() (common.Address, error)
	GetExpectedWorkflowId      func() ([32]byte, error)
	GetExpectedWorkflowName    func() ([10]byte, error)
	GetForwarderAddress        func() (common.Address, error)
	LatestMarket               func() (common.Address, error)
	Markets                    func(MarketsInput) (common.Address, error)
	Name                       func() (string, error)
	Oracle                     func() (common.Address, error)
	Owner                      func() (common.Address, error)
	OwnershipHandoverExpiresAt func(OwnershipHandoverExpiresAtInput) (*big.Int, error)
	State                      func() (uint8, error)
	SupportsInterface          func(SupportsInterfaceInput) (bool, error)
	Symbol                     func() (string, error)
	TotalMarkets               func() (*big.Int, error)
}

// NewMarketFactoryMock creates a new MarketFactoryMock for testing.
func NewMarketFactoryMock(address common.Address, clientMock *evmmock.ClientCapability) *MarketFactoryMock {
	mock := &MarketFactoryMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["collateralToken"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.CollateralToken == nil {
				return nil, errors.New("collateralToken method not mocked")
			}
			result, err := mock.CollateralToken()
			if err != nil {
				return nil, err
			}
			return abi.Methods["collateralToken"].Outputs.Pack(result)
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
		string(abi.Methods["latestMarket"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.LatestMarket == nil {
				return nil, errors.New("latestMarket method not mocked")
			}
			result, err := mock.LatestMarket()
			if err != nil {
				return nil, err
			}
			return abi.Methods["latestMarket"].Outputs.Pack(result)
		},
		string(abi.Methods["markets"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Markets == nil {
				return nil, errors.New("markets method not mocked")
			}
			inputs := abi.Methods["markets"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := MarketsInput{
				Arg0: values[0].(*big.Int),
			}

			result, err := mock.Markets(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["markets"].Outputs.Pack(result)
		},
		string(abi.Methods["name"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Name == nil {
				return nil, errors.New("name method not mocked")
			}
			result, err := mock.Name()
			if err != nil {
				return nil, err
			}
			return abi.Methods["name"].Outputs.Pack(result)
		},
		string(abi.Methods["oracle"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Oracle == nil {
				return nil, errors.New("oracle method not mocked")
			}
			result, err := mock.Oracle()
			if err != nil {
				return nil, err
			}
			return abi.Methods["oracle"].Outputs.Pack(result)
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
		string(abi.Methods["state"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.State == nil {
				return nil, errors.New("state method not mocked")
			}
			result, err := mock.State()
			if err != nil {
				return nil, err
			}
			return abi.Methods["state"].Outputs.Pack(result)
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
		string(abi.Methods["symbol"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Symbol == nil {
				return nil, errors.New("symbol method not mocked")
			}
			result, err := mock.Symbol()
			if err != nil {
				return nil, err
			}
			return abi.Methods["symbol"].Outputs.Pack(result)
		},
		string(abi.Methods["totalMarkets"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TotalMarkets == nil {
				return nil, errors.New("totalMarkets method not mocked")
			}
			result, err := mock.TotalMarkets()
			if err != nil {
				return nil, err
			}
			return abi.Methods["totalMarkets"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
