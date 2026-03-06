// Code generated — DO NOT EDIT.

package lens

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var LensMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getAllMarkets\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structILens.MarketView[]\",\"components\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longPositionToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shortPositionToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shortPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"settled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"settlementRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllRounds\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIOracle.Round[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFactory\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILens.FactoryView\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumIMarketFactory.State\"},{\"name\":\"totalMarkets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"latestMarket\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestMarket\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILens.MarketView\",\"components\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longPositionToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shortPositionToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shortPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"settled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"settlementRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarkPriceHistory\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"count_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIMarket.PriceSnapshot[]\",\"components\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"longPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarket\",\"inputs\":[{\"name\":\"market_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILens.MarketView\",\"components\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longPositionToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shortPositionToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shortPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"settled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"settlementRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketPrice\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isLong_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOracle\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structILens.OracleView\",\"components\":[{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"latestRound\",\"type\":\"tuple\",\"internalType\":\"structIOracle.Round\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"keyword\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rollingEMAWindow\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOraclePrice\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedeemable\",\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structILens.RedeemablePosition[]\",\"components\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"marketIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"positionTokenBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"redeemableCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedeemableForMarket\",\"inputs\":[{\"name\":\"market_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isLong_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowData\",\"inputs\":[{\"name\":\"fof_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"factories\",\"type\":\"tuple[]\",\"internalType\":\"structILens.WorkflowFactoryData[]\",\"components\":[{\"name\":\"factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"latestMarket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rollingEMAWindow\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"}]",
}

// Structs
type ILensFactoryView struct {
	Name            string
	Symbol          string
	Oracle          common.Address
	CollateralToken common.Address
	Decimals        uint8
	State           uint8
	TotalMarkets    *big.Int
	LatestMarket    common.Address
}

type ILensMarketView struct {
	Market             common.Address
	LongPositionToken  common.Address
	ShortPositionToken common.Address
	LongPrice          *big.Int
	ShortPrice         *big.Int
	TotalLiquidity     *big.Int
	Decimals           uint8
	Settled            bool
	SettlementRoundId  *big.Int
}

type ILensOracleView struct {
	CurrentRoundId   *big.Int
	LatestRound      IOracleRound
	Decimals         uint8
	Keyword          string
	RollingEMAWindow []*big.Int
}

type ILensRedeemablePosition struct {
	Market               common.Address
	MarketIndex          *big.Int
	IsLong               bool
	PositionTokenBalance *big.Int
	RedeemableCollateral *big.Int
}

type ILensWorkflowFactoryData struct {
	Factory           common.Address
	Name              string
	Oracle            common.Address
	LatestMarket      common.Address
	CollateralBalance *big.Int
	RollingEMAWindow  []*big.Int
}

type IMarketPriceSnapshot struct {
	Timestamp *big.Int
	LongPrice *big.Int
}

type IOracleRound struct {
	Id        *big.Int
	Timestamp *big.Int
	Index     *big.Int
}

// Contract Method Inputs
type GetAllMarketsInput struct {
	Factory common.Address
}

type GetAllRoundsInput struct {
	Factory common.Address
}

type GetFactoryInput struct {
	Factory common.Address
}

type GetLatestMarketInput struct {
	Factory common.Address
}

type GetMarkPriceHistoryInput struct {
	Factory common.Address
	Count   *big.Int
}

type GetMarketInput struct {
	Market common.Address
}

type GetMarketPriceInput struct {
	Factory common.Address
	IsLong  bool
}

type GetOracleInput struct {
	Factory common.Address
}

type GetOraclePriceInput struct {
	Factory common.Address
}

type GetRedeemableInput struct {
	Factory common.Address
	Account common.Address
}

type GetRedeemableForMarketInput struct {
	Market  common.Address
	Account common.Address
	IsLong  bool
}

type GetWorkflowDataInput struct {
	Fof common.Address
}

// Contract Method Outputs
type GetOraclePriceOutput struct {
	Index     *big.Int
	Timestamp *big.Int
}

type GetWorkflowDataOutput struct {
	CollateralToken common.Address
	Factories       []ILensWorkflowFactoryData
}

// Errors

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

// Main Binding Type for Lens
type Lens struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   LensCodec
}

type LensCodec interface {
	EncodeGetAllMarketsMethodCall(in GetAllMarketsInput) ([]byte, error)
	DecodeGetAllMarketsMethodOutput(data []byte) ([]ILensMarketView, error)
	EncodeGetAllRoundsMethodCall(in GetAllRoundsInput) ([]byte, error)
	DecodeGetAllRoundsMethodOutput(data []byte) ([]IOracleRound, error)
	EncodeGetFactoryMethodCall(in GetFactoryInput) ([]byte, error)
	DecodeGetFactoryMethodOutput(data []byte) (ILensFactoryView, error)
	EncodeGetLatestMarketMethodCall(in GetLatestMarketInput) ([]byte, error)
	DecodeGetLatestMarketMethodOutput(data []byte) (ILensMarketView, error)
	EncodeGetMarkPriceHistoryMethodCall(in GetMarkPriceHistoryInput) ([]byte, error)
	DecodeGetMarkPriceHistoryMethodOutput(data []byte) ([]IMarketPriceSnapshot, error)
	EncodeGetMarketMethodCall(in GetMarketInput) ([]byte, error)
	DecodeGetMarketMethodOutput(data []byte) (ILensMarketView, error)
	EncodeGetMarketPriceMethodCall(in GetMarketPriceInput) ([]byte, error)
	DecodeGetMarketPriceMethodOutput(data []byte) (*big.Int, error)
	EncodeGetOracleMethodCall(in GetOracleInput) ([]byte, error)
	DecodeGetOracleMethodOutput(data []byte) (ILensOracleView, error)
	EncodeGetOraclePriceMethodCall(in GetOraclePriceInput) ([]byte, error)
	DecodeGetOraclePriceMethodOutput(data []byte) (GetOraclePriceOutput, error)
	EncodeGetRedeemableMethodCall(in GetRedeemableInput) ([]byte, error)
	DecodeGetRedeemableMethodOutput(data []byte) ([]ILensRedeemablePosition, error)
	EncodeGetRedeemableForMarketMethodCall(in GetRedeemableForMarketInput) ([]byte, error)
	DecodeGetRedeemableForMarketMethodOutput(data []byte) (*big.Int, error)
	EncodeGetWorkflowDataMethodCall(in GetWorkflowDataInput) ([]byte, error)
	DecodeGetWorkflowDataMethodOutput(data []byte) (GetWorkflowDataOutput, error)
	EncodeILensFactoryViewStruct(in ILensFactoryView) ([]byte, error)
	EncodeILensMarketViewStruct(in ILensMarketView) ([]byte, error)
	EncodeILensOracleViewStruct(in ILensOracleView) ([]byte, error)
	EncodeILensRedeemablePositionStruct(in ILensRedeemablePosition) ([]byte, error)
	EncodeILensWorkflowFactoryDataStruct(in ILensWorkflowFactoryData) ([]byte, error)
	EncodeIMarketPriceSnapshotStruct(in IMarketPriceSnapshot) ([]byte, error)
	EncodeIOracleRoundStruct(in IOracleRound) ([]byte, error)
}

func NewLens(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*Lens, error) {
	parsed, err := abi.JSON(strings.NewReader(LensMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &Lens{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (LensCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(LensMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeGetAllMarketsMethodCall(in GetAllMarketsInput) ([]byte, error) {
	return c.abi.Pack("getAllMarkets", in.Factory)
}

func (c *Codec) DecodeGetAllMarketsMethodOutput(data []byte) ([]ILensMarketView, error) {
	vals, err := c.abi.Methods["getAllMarkets"].Outputs.Unpack(data)
	if err != nil {
		return *new([]ILensMarketView), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]ILensMarketView), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []ILensMarketView
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]ILensMarketView), fmt.Errorf("failed to unmarshal to []ILensMarketView: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetAllRoundsMethodCall(in GetAllRoundsInput) ([]byte, error) {
	return c.abi.Pack("getAllRounds", in.Factory)
}

func (c *Codec) DecodeGetAllRoundsMethodOutput(data []byte) ([]IOracleRound, error) {
	vals, err := c.abi.Methods["getAllRounds"].Outputs.Unpack(data)
	if err != nil {
		return *new([]IOracleRound), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]IOracleRound), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []IOracleRound
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]IOracleRound), fmt.Errorf("failed to unmarshal to []IOracleRound: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetFactoryMethodCall(in GetFactoryInput) ([]byte, error) {
	return c.abi.Pack("getFactory", in.Factory)
}

func (c *Codec) DecodeGetFactoryMethodOutput(data []byte) (ILensFactoryView, error) {
	vals, err := c.abi.Methods["getFactory"].Outputs.Unpack(data)
	if err != nil {
		return *new(ILensFactoryView), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(ILensFactoryView), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result ILensFactoryView
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(ILensFactoryView), fmt.Errorf("failed to unmarshal to ILensFactoryView: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetLatestMarketMethodCall(in GetLatestMarketInput) ([]byte, error) {
	return c.abi.Pack("getLatestMarket", in.Factory)
}

func (c *Codec) DecodeGetLatestMarketMethodOutput(data []byte) (ILensMarketView, error) {
	vals, err := c.abi.Methods["getLatestMarket"].Outputs.Unpack(data)
	if err != nil {
		return *new(ILensMarketView), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(ILensMarketView), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result ILensMarketView
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(ILensMarketView), fmt.Errorf("failed to unmarshal to ILensMarketView: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetMarkPriceHistoryMethodCall(in GetMarkPriceHistoryInput) ([]byte, error) {
	return c.abi.Pack("getMarkPriceHistory", in.Factory, in.Count)
}

func (c *Codec) DecodeGetMarkPriceHistoryMethodOutput(data []byte) ([]IMarketPriceSnapshot, error) {
	vals, err := c.abi.Methods["getMarkPriceHistory"].Outputs.Unpack(data)
	if err != nil {
		return *new([]IMarketPriceSnapshot), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]IMarketPriceSnapshot), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []IMarketPriceSnapshot
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]IMarketPriceSnapshot), fmt.Errorf("failed to unmarshal to []IMarketPriceSnapshot: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetMarketMethodCall(in GetMarketInput) ([]byte, error) {
	return c.abi.Pack("getMarket", in.Market)
}

func (c *Codec) DecodeGetMarketMethodOutput(data []byte) (ILensMarketView, error) {
	vals, err := c.abi.Methods["getMarket"].Outputs.Unpack(data)
	if err != nil {
		return *new(ILensMarketView), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(ILensMarketView), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result ILensMarketView
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(ILensMarketView), fmt.Errorf("failed to unmarshal to ILensMarketView: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetMarketPriceMethodCall(in GetMarketPriceInput) ([]byte, error) {
	return c.abi.Pack("getMarketPrice", in.Factory, in.IsLong)
}

func (c *Codec) DecodeGetMarketPriceMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getMarketPrice"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetOracleMethodCall(in GetOracleInput) ([]byte, error) {
	return c.abi.Pack("getOracle", in.Factory)
}

func (c *Codec) DecodeGetOracleMethodOutput(data []byte) (ILensOracleView, error) {
	vals, err := c.abi.Methods["getOracle"].Outputs.Unpack(data)
	if err != nil {
		return *new(ILensOracleView), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(ILensOracleView), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result ILensOracleView
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(ILensOracleView), fmt.Errorf("failed to unmarshal to ILensOracleView: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetOraclePriceMethodCall(in GetOraclePriceInput) ([]byte, error) {
	return c.abi.Pack("getOraclePrice", in.Factory)
}

func (c *Codec) DecodeGetOraclePriceMethodOutput(data []byte) (GetOraclePriceOutput, error) {
	vals, err := c.abi.Methods["getOraclePrice"].Outputs.Unpack(data)
	if err != nil {
		return GetOraclePriceOutput{}, err
	}
	if len(vals) != 2 {
		return GetOraclePriceOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return GetOraclePriceOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 *big.Int
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return GetOraclePriceOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return GetOraclePriceOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 *big.Int
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return GetOraclePriceOutput{}, fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return GetOraclePriceOutput{
		Index:     result0,
		Timestamp: result1,
	}, nil
}

func (c *Codec) EncodeGetRedeemableMethodCall(in GetRedeemableInput) ([]byte, error) {
	return c.abi.Pack("getRedeemable", in.Factory, in.Account)
}

func (c *Codec) DecodeGetRedeemableMethodOutput(data []byte) ([]ILensRedeemablePosition, error) {
	vals, err := c.abi.Methods["getRedeemable"].Outputs.Unpack(data)
	if err != nil {
		return *new([]ILensRedeemablePosition), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]ILensRedeemablePosition), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []ILensRedeemablePosition
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]ILensRedeemablePosition), fmt.Errorf("failed to unmarshal to []ILensRedeemablePosition: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetRedeemableForMarketMethodCall(in GetRedeemableForMarketInput) ([]byte, error) {
	return c.abi.Pack("getRedeemableForMarket", in.Market, in.Account, in.IsLong)
}

func (c *Codec) DecodeGetRedeemableForMarketMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getRedeemableForMarket"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetWorkflowDataMethodCall(in GetWorkflowDataInput) ([]byte, error) {
	return c.abi.Pack("getWorkflowData", in.Fof)
}

func (c *Codec) DecodeGetWorkflowDataMethodOutput(data []byte) (GetWorkflowDataOutput, error) {
	vals, err := c.abi.Methods["getWorkflowData"].Outputs.Unpack(data)
	if err != nil {
		return GetWorkflowDataOutput{}, err
	}
	if len(vals) != 2 {
		return GetWorkflowDataOutput{}, fmt.Errorf("expected 2 values, got %d", len(vals))
	}
	jsonData0, err := json.Marshal(vals[0])
	if err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to marshal ABI result 0: %w", err)
	}

	var result0 common.Address
	if err := json.Unmarshal(jsonData0, &result0); err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}
	jsonData1, err := json.Marshal(vals[1])
	if err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to marshal ABI result 1: %w", err)
	}

	var result1 []ILensWorkflowFactoryData
	if err := json.Unmarshal(jsonData1, &result1); err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to unmarshal to []ILensWorkflowFactoryData: %w", err)
	}

	return GetWorkflowDataOutput{
		CollateralToken: result0,
		Factories:       result1,
	}, nil
}

func (c *Codec) EncodeILensFactoryViewStruct(in ILensFactoryView) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "name", Type: "string"},
			{Name: "symbol", Type: "string"},
			{Name: "oracle", Type: "address"},
			{Name: "collateralToken", Type: "address"},
			{Name: "decimals", Type: "uint8"},
			{Name: "state", Type: "uint8"},
			{Name: "totalMarkets", Type: "uint256"},
			{Name: "latestMarket", Type: "address"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for ILensFactoryView: %w", err)
	}
	args := abi.Arguments{
		{Name: "iLensFactoryView", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeILensMarketViewStruct(in ILensMarketView) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "market", Type: "address"},
			{Name: "longPositionToken", Type: "address"},
			{Name: "shortPositionToken", Type: "address"},
			{Name: "longPrice", Type: "uint256"},
			{Name: "shortPrice", Type: "uint256"},
			{Name: "totalLiquidity", Type: "uint256"},
			{Name: "decimals", Type: "uint8"},
			{Name: "settled", Type: "bool"},
			{Name: "settlementRoundId", Type: "uint256"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for ILensMarketView: %w", err)
	}
	args := abi.Arguments{
		{Name: "iLensMarketView", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeILensOracleViewStruct(in ILensOracleView) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "currentRoundId", Type: "uint256"},
			{Name: "latestRound", Type: "(uint256,uint256,uint256)"},
			{Name: "decimals", Type: "uint8"},
			{Name: "keyword", Type: "string"},
			{Name: "rollingEMAWindow", Type: "uint256[]"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for ILensOracleView: %w", err)
	}
	args := abi.Arguments{
		{Name: "iLensOracleView", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeILensRedeemablePositionStruct(in ILensRedeemablePosition) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "market", Type: "address"},
			{Name: "marketIndex", Type: "uint256"},
			{Name: "isLong", Type: "bool"},
			{Name: "positionTokenBalance", Type: "uint256"},
			{Name: "redeemableCollateral", Type: "uint256"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for ILensRedeemablePosition: %w", err)
	}
	args := abi.Arguments{
		{Name: "iLensRedeemablePosition", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeILensWorkflowFactoryDataStruct(in ILensWorkflowFactoryData) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "factory", Type: "address"},
			{Name: "name", Type: "string"},
			{Name: "oracle", Type: "address"},
			{Name: "latestMarket", Type: "address"},
			{Name: "collateralBalance", Type: "uint256"},
			{Name: "rollingEMAWindow", Type: "uint256[]"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for ILensWorkflowFactoryData: %w", err)
	}
	args := abi.Arguments{
		{Name: "iLensWorkflowFactoryData", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeIMarketPriceSnapshotStruct(in IMarketPriceSnapshot) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "timestamp", Type: "uint256"},
			{Name: "longPrice", Type: "uint256"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for IMarketPriceSnapshot: %w", err)
	}
	args := abi.Arguments{
		{Name: "iMarketPriceSnapshot", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *Codec) EncodeIOracleRoundStruct(in IOracleRound) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "id", Type: "uint256"},
			{Name: "timestamp", Type: "uint256"},
			{Name: "index", Type: "uint256"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for IOracleRound: %w", err)
	}
	args := abi.Arguments{
		{Name: "iOracleRound", Type: tupleType},
	}

	return args.Pack(in)
}

func (c Lens) GetAllMarkets(
	runtime cre.Runtime,
	args GetAllMarketsInput,
	blockNumber *big.Int,
) cre.Promise[[]ILensMarketView] {
	calldata, err := c.Codec.EncodeGetAllMarketsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[]ILensMarketView](*new([]ILensMarketView), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([]ILensMarketView, error) {
		return c.Codec.DecodeGetAllMarketsMethodOutput(response.Data)
	})

}

func (c Lens) GetAllRounds(
	runtime cre.Runtime,
	args GetAllRoundsInput,
	blockNumber *big.Int,
) cre.Promise[[]IOracleRound] {
	calldata, err := c.Codec.EncodeGetAllRoundsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[]IOracleRound](*new([]IOracleRound), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([]IOracleRound, error) {
		return c.Codec.DecodeGetAllRoundsMethodOutput(response.Data)
	})

}

func (c Lens) GetFactory(
	runtime cre.Runtime,
	args GetFactoryInput,
	blockNumber *big.Int,
) cre.Promise[ILensFactoryView] {
	calldata, err := c.Codec.EncodeGetFactoryMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[ILensFactoryView](*new(ILensFactoryView), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (ILensFactoryView, error) {
		return c.Codec.DecodeGetFactoryMethodOutput(response.Data)
	})

}

func (c Lens) GetLatestMarket(
	runtime cre.Runtime,
	args GetLatestMarketInput,
	blockNumber *big.Int,
) cre.Promise[ILensMarketView] {
	calldata, err := c.Codec.EncodeGetLatestMarketMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[ILensMarketView](*new(ILensMarketView), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (ILensMarketView, error) {
		return c.Codec.DecodeGetLatestMarketMethodOutput(response.Data)
	})

}

func (c Lens) GetMarkPriceHistory(
	runtime cre.Runtime,
	args GetMarkPriceHistoryInput,
	blockNumber *big.Int,
) cre.Promise[[]IMarketPriceSnapshot] {
	calldata, err := c.Codec.EncodeGetMarkPriceHistoryMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[]IMarketPriceSnapshot](*new([]IMarketPriceSnapshot), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([]IMarketPriceSnapshot, error) {
		return c.Codec.DecodeGetMarkPriceHistoryMethodOutput(response.Data)
	})

}

func (c Lens) GetMarket(
	runtime cre.Runtime,
	args GetMarketInput,
	blockNumber *big.Int,
) cre.Promise[ILensMarketView] {
	calldata, err := c.Codec.EncodeGetMarketMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[ILensMarketView](*new(ILensMarketView), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (ILensMarketView, error) {
		return c.Codec.DecodeGetMarketMethodOutput(response.Data)
	})

}

func (c Lens) GetMarketPrice(
	runtime cre.Runtime,
	args GetMarketPriceInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetMarketPriceMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeGetMarketPriceMethodOutput(response.Data)
	})

}

func (c Lens) GetOracle(
	runtime cre.Runtime,
	args GetOracleInput,
	blockNumber *big.Int,
) cre.Promise[ILensOracleView] {
	calldata, err := c.Codec.EncodeGetOracleMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[ILensOracleView](*new(ILensOracleView), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (ILensOracleView, error) {
		return c.Codec.DecodeGetOracleMethodOutput(response.Data)
	})

}

func (c Lens) GetOraclePrice(
	runtime cre.Runtime,
	args GetOraclePriceInput,
	blockNumber *big.Int,
) cre.Promise[GetOraclePriceOutput] {
	calldata, err := c.Codec.EncodeGetOraclePriceMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[GetOraclePriceOutput](GetOraclePriceOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (GetOraclePriceOutput, error) {
		return c.Codec.DecodeGetOraclePriceMethodOutput(response.Data)
	})

}

func (c Lens) GetRedeemable(
	runtime cre.Runtime,
	args GetRedeemableInput,
	blockNumber *big.Int,
) cre.Promise[[]ILensRedeemablePosition] {
	calldata, err := c.Codec.EncodeGetRedeemableMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[]ILensRedeemablePosition](*new([]ILensRedeemablePosition), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([]ILensRedeemablePosition, error) {
		return c.Codec.DecodeGetRedeemableMethodOutput(response.Data)
	})

}

func (c Lens) GetRedeemableForMarket(
	runtime cre.Runtime,
	args GetRedeemableForMarketInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetRedeemableForMarketMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeGetRedeemableForMarketMethodOutput(response.Data)
	})

}

func (c Lens) GetWorkflowData(
	runtime cre.Runtime,
	args GetWorkflowDataInput,
	blockNumber *big.Int,
) cre.Promise[GetWorkflowDataOutput] {
	calldata, err := c.Codec.EncodeGetWorkflowDataMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[GetWorkflowDataOutput](GetWorkflowDataOutput{}, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (GetWorkflowDataOutput, error) {
		return c.Codec.DecodeGetWorkflowDataMethodOutput(response.Data)
	})

}

func (c Lens) WriteReportFromILensFactoryView(
	runtime cre.Runtime,
	input ILensFactoryView,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeILensFactoryViewStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReportFromILensMarketView(
	runtime cre.Runtime,
	input ILensMarketView,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeILensMarketViewStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReportFromILensOracleView(
	runtime cre.Runtime,
	input ILensOracleView,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeILensOracleViewStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReportFromILensRedeemablePosition(
	runtime cre.Runtime,
	input ILensRedeemablePosition,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeILensRedeemablePositionStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReportFromILensWorkflowFactoryData(
	runtime cre.Runtime,
	input ILensWorkflowFactoryData,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeILensWorkflowFactoryDataStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReportFromIMarketPriceSnapshot(
	runtime cre.Runtime,
	input IMarketPriceSnapshot,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeIMarketPriceSnapshotStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReportFromIOracleRound(
	runtime cre.Runtime,
	input IOracleRound,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeIOracleRoundStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c Lens) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

func (c *Lens) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}
