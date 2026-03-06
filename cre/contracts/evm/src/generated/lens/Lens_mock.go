// Code generated — DO NOT EDIT.

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
	GetAllMarkets          func(GetAllMarketsInput) ([]ILensMarketView, error)
	GetAllRounds           func(GetAllRoundsInput) ([]IOracleRound, error)
	GetFactory             func(GetFactoryInput) (ILensFactoryView, error)
	GetLatestMarket        func(GetLatestMarketInput) (ILensMarketView, error)
	GetMarkPriceHistory    func(GetMarkPriceHistoryInput) ([]IMarketPriceSnapshot, error)
	GetMarket              func(GetMarketInput) (ILensMarketView, error)
	GetMarketPrice         func(GetMarketPriceInput) (*big.Int, error)
	GetOracle              func(GetOracleInput) (ILensOracleView, error)
	GetOraclePrice         func(GetOraclePriceInput) (GetOraclePriceOutput, error)
	GetRedeemable          func(GetRedeemableInput) ([]ILensRedeemablePosition, error)
	GetRedeemableForMarket func(GetRedeemableForMarketInput) (*big.Int, error)
	GetWorkflowData        func(GetWorkflowDataInput) (GetWorkflowDataOutput, error)
}

// NewLensMock creates a new LensMock for testing.
func NewLensMock(address common.Address, clientMock *evmmock.ClientCapability) *LensMock {
	mock := &LensMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["getAllMarkets"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetAllMarkets == nil {
				return nil, errors.New("getAllMarkets method not mocked")
			}
			inputs := abi.Methods["getAllMarkets"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetAllMarketsInput{
				Factory: values[0].(common.Address),
			}

			result, err := mock.GetAllMarkets(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getAllMarkets"].Outputs.Pack(result)
		},
		string(abi.Methods["getAllRounds"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetAllRounds == nil {
				return nil, errors.New("getAllRounds method not mocked")
			}
			inputs := abi.Methods["getAllRounds"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetAllRoundsInput{
				Factory: values[0].(common.Address),
			}

			result, err := mock.GetAllRounds(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getAllRounds"].Outputs.Pack(result)
		},
		string(abi.Methods["getFactory"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetFactory == nil {
				return nil, errors.New("getFactory method not mocked")
			}
			inputs := abi.Methods["getFactory"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetFactoryInput{
				Factory: values[0].(common.Address),
			}

			result, err := mock.GetFactory(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getFactory"].Outputs.Pack(result)
		},
		string(abi.Methods["getLatestMarket"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetLatestMarket == nil {
				return nil, errors.New("getLatestMarket method not mocked")
			}
			inputs := abi.Methods["getLatestMarket"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetLatestMarketInput{
				Factory: values[0].(common.Address),
			}

			result, err := mock.GetLatestMarket(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getLatestMarket"].Outputs.Pack(result)
		},
		string(abi.Methods["getMarkPriceHistory"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetMarkPriceHistory == nil {
				return nil, errors.New("getMarkPriceHistory method not mocked")
			}
			inputs := abi.Methods["getMarkPriceHistory"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 2 {
				return nil, errors.New("expected 2 input values")
			}

			args := GetMarkPriceHistoryInput{
				Factory: values[0].(common.Address),
				Count:   values[1].(*big.Int),
			}

			result, err := mock.GetMarkPriceHistory(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getMarkPriceHistory"].Outputs.Pack(result)
		},
		string(abi.Methods["getMarket"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetMarket == nil {
				return nil, errors.New("getMarket method not mocked")
			}
			inputs := abi.Methods["getMarket"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetMarketInput{
				Market: values[0].(common.Address),
			}

			result, err := mock.GetMarket(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getMarket"].Outputs.Pack(result)
		},
		string(abi.Methods["getMarketPrice"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetMarketPrice == nil {
				return nil, errors.New("getMarketPrice method not mocked")
			}
			inputs := abi.Methods["getMarketPrice"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 2 {
				return nil, errors.New("expected 2 input values")
			}

			args := GetMarketPriceInput{
				Factory: values[0].(common.Address),
				IsLong:  values[1].(bool),
			}

			result, err := mock.GetMarketPrice(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getMarketPrice"].Outputs.Pack(result)
		},
		string(abi.Methods["getOracle"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetOracle == nil {
				return nil, errors.New("getOracle method not mocked")
			}
			inputs := abi.Methods["getOracle"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetOracleInput{
				Factory: values[0].(common.Address),
			}

			result, err := mock.GetOracle(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getOracle"].Outputs.Pack(result)
		},
		string(abi.Methods["getOraclePrice"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetOraclePrice == nil {
				return nil, errors.New("getOraclePrice method not mocked")
			}
			inputs := abi.Methods["getOraclePrice"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetOraclePriceInput{
				Factory: values[0].(common.Address),
			}

			result, err := mock.GetOraclePrice(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getOraclePrice"].Outputs.Pack(
				result.Index,
				result.Timestamp,
			)
		},
		string(abi.Methods["getRedeemable"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetRedeemable == nil {
				return nil, errors.New("getRedeemable method not mocked")
			}
			inputs := abi.Methods["getRedeemable"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 2 {
				return nil, errors.New("expected 2 input values")
			}

			args := GetRedeemableInput{
				Factory: values[0].(common.Address),
				Account: values[1].(common.Address),
			}

			result, err := mock.GetRedeemable(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getRedeemable"].Outputs.Pack(result)
		},
		string(abi.Methods["getRedeemableForMarket"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetRedeemableForMarket == nil {
				return nil, errors.New("getRedeemableForMarket method not mocked")
			}
			inputs := abi.Methods["getRedeemableForMarket"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 3 {
				return nil, errors.New("expected 3 input values")
			}

			args := GetRedeemableForMarketInput{
				Market:  values[0].(common.Address),
				Account: values[1].(common.Address),
				IsLong:  values[2].(bool),
			}

			result, err := mock.GetRedeemableForMarket(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getRedeemableForMarket"].Outputs.Pack(result)
		},
		string(abi.Methods["getWorkflowData"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetWorkflowData == nil {
				return nil, errors.New("getWorkflowData method not mocked")
			}
			inputs := abi.Methods["getWorkflowData"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetWorkflowDataInput{
				Fof: values[0].(common.Address),
			}

			result, err := mock.GetWorkflowData(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getWorkflowData"].Outputs.Pack(
				result.CollateralToken,
				result.Factories,
			)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
