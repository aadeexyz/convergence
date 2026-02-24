// Code generated — DO NOT EDIT.

//go:build !wasip1

package market

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

// MarketMock is a mock implementation of Market for testing.
type MarketMock struct {
	CollateralToken            func() (common.Address, error)
	Decimals                   func() (uint8, error)
	LongPositionToken          func() (common.Address, error)
	Oracle                     func() (common.Address, error)
	Owner                      func() (common.Address, error)
	OwnershipHandoverExpiresAt func(OwnershipHandoverExpiresAtInput) (*big.Int, error)
	Price                      func(PriceInput) (*big.Int, error)
	Settled                    func() (bool, error)
	SettlementRoundId          func() (*big.Int, error)
	ShortPositionToken         func() (common.Address, error)
}

// NewMarketMock creates a new MarketMock for testing.
func NewMarketMock(address common.Address, clientMock *evmmock.ClientCapability) *MarketMock {
	mock := &MarketMock{}

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
		string(abi.Methods["longPositionToken"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.LongPositionToken == nil {
				return nil, errors.New("longPositionToken method not mocked")
			}
			result, err := mock.LongPositionToken()
			if err != nil {
				return nil, err
			}
			return abi.Methods["longPositionToken"].Outputs.Pack(result)
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
		string(abi.Methods["price"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Price == nil {
				return nil, errors.New("price method not mocked")
			}
			inputs := abi.Methods["price"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := PriceInput{
				IsLong: values[0].(bool),
			}

			result, err := mock.Price(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["price"].Outputs.Pack(result)
		},
		string(abi.Methods["settled"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Settled == nil {
				return nil, errors.New("settled method not mocked")
			}
			result, err := mock.Settled()
			if err != nil {
				return nil, err
			}
			return abi.Methods["settled"].Outputs.Pack(result)
		},
		string(abi.Methods["settlementRoundId"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.SettlementRoundId == nil {
				return nil, errors.New("settlementRoundId method not mocked")
			}
			result, err := mock.SettlementRoundId()
			if err != nil {
				return nil, err
			}
			return abi.Methods["settlementRoundId"].Outputs.Pack(result)
		},
		string(abi.Methods["shortPositionToken"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.ShortPositionToken == nil {
				return nil, errors.New("shortPositionToken method not mocked")
			}
			result, err := mock.ShortPositionToken()
			if err != nil {
				return nil, err
			}
			return abi.Methods["shortPositionToken"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
