// Code generated — DO NOT EDIT.

//go:build !wasip1

package factory_of_market_factories

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

// FactoryOfMarketFactoriesMock is a mock implementation of FactoryOfMarketFactories for testing.
type FactoryOfMarketFactoriesMock struct {
	CollateralToken            func() (common.Address, error)
	CreationFee                func() (*big.Int, error)
	ForwarderAddress           func() (common.Address, error)
	LiquidityFee               func() (*big.Int, error)
	MarketFactories            func() ([]common.Address, error)
	MarketFactory              func(MarketFactoryInput) (common.Address, error)
	OracleDecimals             func() (uint8, error)
	Owner                      func() (common.Address, error)
	OwnershipHandoverExpiresAt func(OwnershipHandoverExpiresAtInput) (*big.Int, error)
	ProtocolFee                func() (*big.Int, error)
	TotalMarketFactories       func() (*big.Int, error)
}

// NewFactoryOfMarketFactoriesMock creates a new FactoryOfMarketFactoriesMock for testing.
func NewFactoryOfMarketFactoriesMock(address common.Address, clientMock *evmmock.ClientCapability) *FactoryOfMarketFactoriesMock {
	mock := &FactoryOfMarketFactoriesMock{}

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
		string(abi.Methods["creationFee"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.CreationFee == nil {
				return nil, errors.New("creationFee method not mocked")
			}
			result, err := mock.CreationFee()
			if err != nil {
				return nil, err
			}
			return abi.Methods["creationFee"].Outputs.Pack(result)
		},
		string(abi.Methods["forwarderAddress"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ForwarderAddress == nil {
				return nil, errors.New("forwarderAddress method not mocked")
			}
			result, err := mock.ForwarderAddress()
			if err != nil {
				return nil, err
			}
			return abi.Methods["forwarderAddress"].Outputs.Pack(result)
		},
		string(abi.Methods["liquidityFee"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.LiquidityFee == nil {
				return nil, errors.New("liquidityFee method not mocked")
			}
			result, err := mock.LiquidityFee()
			if err != nil {
				return nil, err
			}
			return abi.Methods["liquidityFee"].Outputs.Pack(result)
		},
		string(abi.Methods["marketFactories"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MarketFactories == nil {
				return nil, errors.New("marketFactories method not mocked")
			}
			result, err := mock.MarketFactories()
			if err != nil {
				return nil, err
			}
			return abi.Methods["marketFactories"].Outputs.Pack(result)
		},
		string(abi.Methods["marketFactory"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.MarketFactory == nil {
				return nil, errors.New("marketFactory method not mocked")
			}
			inputs := abi.Methods["marketFactory"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := MarketFactoryInput{
				Index: values[0].(*big.Int),
			}

			result, err := mock.MarketFactory(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["marketFactory"].Outputs.Pack(result)
		},
		string(abi.Methods["oracleDecimals"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.OracleDecimals == nil {
				return nil, errors.New("oracleDecimals method not mocked")
			}
			result, err := mock.OracleDecimals()
			if err != nil {
				return nil, err
			}
			return abi.Methods["oracleDecimals"].Outputs.Pack(result)
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
		string(abi.Methods["protocolFee"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ProtocolFee == nil {
				return nil, errors.New("protocolFee method not mocked")
			}
			result, err := mock.ProtocolFee()
			if err != nil {
				return nil, err
			}
			return abi.Methods["protocolFee"].Outputs.Pack(result)
		},
		string(abi.Methods["totalMarketFactories"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.TotalMarketFactories == nil {
				return nil, errors.New("totalMarketFactories method not mocked")
			}
			result, err := mock.TotalMarketFactories()
			if err != nil {
				return nil, err
			}
			return abi.Methods["totalMarketFactories"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
