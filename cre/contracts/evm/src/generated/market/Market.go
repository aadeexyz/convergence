// Code generated — DO NOT EDIT.

package market

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

var MarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"collateralToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"isLong_\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"positionTokenAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"collateralToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"longPositionToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPositionToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"isLong_\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"collateralAmount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"price\",\"inputs\":[{\"name\":\"isLong_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"isLong_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"seed\",\"inputs\":[{\"name\":\"longCollateral_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shortCollateral_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settle\",\"inputs\":[{\"name\":\"settlementRoundId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"settlementRoundId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shortPositionToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPositionToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"MarketSeeded\",\"inputs\":[{\"name\":\"seeder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"longCollateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shortCollateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"positionTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketSettled\",\"inputs\":[{\"name\":\"settlementRoundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionBurned\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"positionTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionMinted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"positionTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PositionRedeemed\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"positionTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySeeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientCollateral\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPositionTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// Structs

// Contract Method Inputs
type BurnInput struct {
	IsLong              bool
	PositionTokenAmount *big.Int
}

type CompleteOwnershipHandoverInput struct {
	PendingOwner common.Address
}

type MintInput struct {
	IsLong           bool
	CollateralAmount *big.Int
}

type OwnershipHandoverExpiresAtInput struct {
	PendingOwner common.Address
}

type PriceInput struct {
	IsLong bool
}

type RedeemInput struct {
	IsLong bool
}

type SeedInput struct {
	LongCollateral  *big.Int
	ShortCollateral *big.Int
}

type SettleInput struct {
	SettlementRoundId *big.Int
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

// Contract Method Outputs

// Errors
type AlreadyInitialized struct {
}

type AlreadySeeded struct {
}

type AlreadySettled struct {
}

type InsufficientCollateral struct {
}

type InsufficientPositionTokens struct {
}

type NewOwnerIsZeroAddress struct {
}

type NoHandoverRequest struct {
}

type NotSettled struct {
}

type Unauthorized struct {
}

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type MarketSeededTopics struct {
	Seeder common.Address
}

type MarketSeededDecoded struct {
	Seeder              common.Address
	LongCollateral      *big.Int
	ShortCollateral     *big.Int
	PositionTokenAmount *big.Int
}

type MarketSettledTopics struct {
	SettlementRoundId *big.Int
}

type MarketSettledDecoded struct {
	SettlementRoundId *big.Int
}

type OwnershipHandoverCanceledTopics struct {
	PendingOwner common.Address
}

type OwnershipHandoverCanceledDecoded struct {
	PendingOwner common.Address
}

type OwnershipHandoverRequestedTopics struct {
	PendingOwner common.Address
}

type OwnershipHandoverRequestedDecoded struct {
	PendingOwner common.Address
}

type OwnershipTransferredTopics struct {
	OldOwner common.Address
	NewOwner common.Address
}

type OwnershipTransferredDecoded struct {
	OldOwner common.Address
	NewOwner common.Address
}

type PositionBurnedTopics struct {
	Account common.Address
	IsLong  bool
}

type PositionBurnedDecoded struct {
	Account             common.Address
	IsLong              bool
	PositionTokenAmount *big.Int
	CollateralAmount    *big.Int
}

type PositionMintedTopics struct {
	Account common.Address
	IsLong  bool
}

type PositionMintedDecoded struct {
	Account             common.Address
	IsLong              bool
	CollateralAmount    *big.Int
	PositionTokenAmount *big.Int
}

type PositionRedeemedTopics struct {
	Account common.Address
	IsLong  bool
}

type PositionRedeemedDecoded struct {
	Account             common.Address
	IsLong              bool
	PositionTokenAmount *big.Int
	CollateralAmount    *big.Int
}

// Main Binding Type for Market
type Market struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   MarketCodec
}

type MarketCodec interface {
	EncodeBurnMethodCall(in BurnInput) ([]byte, error)
	EncodeCancelOwnershipHandoverMethodCall() ([]byte, error)
	EncodeCollateralTokenMethodCall() ([]byte, error)
	DecodeCollateralTokenMethodOutput(data []byte) (common.Address, error)
	EncodeCompleteOwnershipHandoverMethodCall(in CompleteOwnershipHandoverInput) ([]byte, error)
	EncodeDecimalsMethodCall() ([]byte, error)
	DecodeDecimalsMethodOutput(data []byte) (uint8, error)
	EncodeLongPositionTokenMethodCall() ([]byte, error)
	DecodeLongPositionTokenMethodOutput(data []byte) (common.Address, error)
	EncodeMintMethodCall(in MintInput) ([]byte, error)
	EncodeOracleMethodCall() ([]byte, error)
	DecodeOracleMethodOutput(data []byte) (common.Address, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeOwnershipHandoverExpiresAtMethodCall(in OwnershipHandoverExpiresAtInput) ([]byte, error)
	DecodeOwnershipHandoverExpiresAtMethodOutput(data []byte) (*big.Int, error)
	EncodePriceMethodCall(in PriceInput) ([]byte, error)
	DecodePriceMethodOutput(data []byte) (*big.Int, error)
	EncodeRedeemMethodCall(in RedeemInput) ([]byte, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeRequestOwnershipHandoverMethodCall() ([]byte, error)
	EncodeSeedMethodCall(in SeedInput) ([]byte, error)
	EncodeSettleMethodCall(in SettleInput) ([]byte, error)
	EncodeSettledMethodCall() ([]byte, error)
	DecodeSettledMethodOutput(data []byte) (bool, error)
	EncodeSettlementRoundIdMethodCall() ([]byte, error)
	DecodeSettlementRoundIdMethodOutput(data []byte) (*big.Int, error)
	EncodeShortPositionTokenMethodCall() ([]byte, error)
	DecodeShortPositionTokenMethodOutput(data []byte) (common.Address, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	MarketSeededLogHash() []byte
	EncodeMarketSeededTopics(evt abi.Event, values []MarketSeededTopics) ([]*evm.TopicValues, error)
	DecodeMarketSeeded(log *evm.Log) (*MarketSeededDecoded, error)
	MarketSettledLogHash() []byte
	EncodeMarketSettledTopics(evt abi.Event, values []MarketSettledTopics) ([]*evm.TopicValues, error)
	DecodeMarketSettled(log *evm.Log) (*MarketSettledDecoded, error)
	OwnershipHandoverCanceledLogHash() []byte
	EncodeOwnershipHandoverCanceledTopics(evt abi.Event, values []OwnershipHandoverCanceledTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipHandoverCanceled(log *evm.Log) (*OwnershipHandoverCanceledDecoded, error)
	OwnershipHandoverRequestedLogHash() []byte
	EncodeOwnershipHandoverRequestedTopics(evt abi.Event, values []OwnershipHandoverRequestedTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipHandoverRequested(log *evm.Log) (*OwnershipHandoverRequestedDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
	PositionBurnedLogHash() []byte
	EncodePositionBurnedTopics(evt abi.Event, values []PositionBurnedTopics) ([]*evm.TopicValues, error)
	DecodePositionBurned(log *evm.Log) (*PositionBurnedDecoded, error)
	PositionMintedLogHash() []byte
	EncodePositionMintedTopics(evt abi.Event, values []PositionMintedTopics) ([]*evm.TopicValues, error)
	DecodePositionMinted(log *evm.Log) (*PositionMintedDecoded, error)
	PositionRedeemedLogHash() []byte
	EncodePositionRedeemedTopics(evt abi.Event, values []PositionRedeemedTopics) ([]*evm.TopicValues, error)
	DecodePositionRedeemed(log *evm.Log) (*PositionRedeemedDecoded, error)
}

func NewMarket(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &Market{
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

func NewCodec() (MarketCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeBurnMethodCall(in BurnInput) ([]byte, error) {
	return c.abi.Pack("burn", in.IsLong, in.PositionTokenAmount)
}

func (c *Codec) EncodeCancelOwnershipHandoverMethodCall() ([]byte, error) {
	return c.abi.Pack("cancelOwnershipHandover")
}

func (c *Codec) EncodeCollateralTokenMethodCall() ([]byte, error) {
	return c.abi.Pack("collateralToken")
}

func (c *Codec) DecodeCollateralTokenMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["collateralToken"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeCompleteOwnershipHandoverMethodCall(in CompleteOwnershipHandoverInput) ([]byte, error) {
	return c.abi.Pack("completeOwnershipHandover", in.PendingOwner)
}

func (c *Codec) EncodeDecimalsMethodCall() ([]byte, error) {
	return c.abi.Pack("decimals")
}

func (c *Codec) DecodeDecimalsMethodOutput(data []byte) (uint8, error) {
	vals, err := c.abi.Methods["decimals"].Outputs.Unpack(data)
	if err != nil {
		return *new(uint8), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(uint8), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result uint8
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(uint8), fmt.Errorf("failed to unmarshal to uint8: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeLongPositionTokenMethodCall() ([]byte, error) {
	return c.abi.Pack("longPositionToken")
}

func (c *Codec) DecodeLongPositionTokenMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["longPositionToken"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeMintMethodCall(in MintInput) ([]byte, error) {
	return c.abi.Pack("mint", in.IsLong, in.CollateralAmount)
}

func (c *Codec) EncodeOracleMethodCall() ([]byte, error) {
	return c.abi.Pack("oracle")
}

func (c *Codec) DecodeOracleMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["oracle"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOwnerMethodCall() ([]byte, error) {
	return c.abi.Pack("owner")
}

func (c *Codec) DecodeOwnerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["owner"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOwnershipHandoverExpiresAtMethodCall(in OwnershipHandoverExpiresAtInput) ([]byte, error) {
	return c.abi.Pack("ownershipHandoverExpiresAt", in.PendingOwner)
}

func (c *Codec) DecodeOwnershipHandoverExpiresAtMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["ownershipHandoverExpiresAt"].Outputs.Unpack(data)
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

func (c *Codec) EncodePriceMethodCall(in PriceInput) ([]byte, error) {
	return c.abi.Pack("price", in.IsLong)
}

func (c *Codec) DecodePriceMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["price"].Outputs.Unpack(data)
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

func (c *Codec) EncodeRedeemMethodCall(in RedeemInput) ([]byte, error) {
	return c.abi.Pack("redeem", in.IsLong)
}

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeRequestOwnershipHandoverMethodCall() ([]byte, error) {
	return c.abi.Pack("requestOwnershipHandover")
}

func (c *Codec) EncodeSeedMethodCall(in SeedInput) ([]byte, error) {
	return c.abi.Pack("seed", in.LongCollateral, in.ShortCollateral)
}

func (c *Codec) EncodeSettleMethodCall(in SettleInput) ([]byte, error) {
	return c.abi.Pack("settle", in.SettlementRoundId)
}

func (c *Codec) EncodeSettledMethodCall() ([]byte, error) {
	return c.abi.Pack("settled")
}

func (c *Codec) DecodeSettledMethodOutput(data []byte) (bool, error) {
	vals, err := c.abi.Methods["settled"].Outputs.Unpack(data)
	if err != nil {
		return *new(bool), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(bool), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result bool
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(bool), fmt.Errorf("failed to unmarshal to bool: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeSettlementRoundIdMethodCall() ([]byte, error) {
	return c.abi.Pack("settlementRoundId")
}

func (c *Codec) DecodeSettlementRoundIdMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["settlementRoundId"].Outputs.Unpack(data)
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

func (c *Codec) EncodeShortPositionTokenMethodCall() ([]byte, error) {
	return c.abi.Pack("shortPositionToken")
}

func (c *Codec) DecodeShortPositionTokenMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["shortPositionToken"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
}

func (c *Codec) MarketSeededLogHash() []byte {
	return c.abi.Events["MarketSeeded"].ID.Bytes()
}

func (c *Codec) EncodeMarketSeededTopics(
	evt abi.Event,
	values []MarketSeededTopics,
) ([]*evm.TopicValues, error) {
	var seederRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Seeder).IsZero() {
			seederRule = append(seederRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Seeder)
		if err != nil {
			return nil, err
		}
		seederRule = append(seederRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		seederRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMarketSeeded decodes a log into a MarketSeeded struct.
func (c *Codec) DecodeMarketSeeded(log *evm.Log) (*MarketSeededDecoded, error) {
	event := new(MarketSeededDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MarketSeeded", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MarketSeeded"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) MarketSettledLogHash() []byte {
	return c.abi.Events["MarketSettled"].ID.Bytes()
}

func (c *Codec) EncodeMarketSettledTopics(
	evt abi.Event,
	values []MarketSettledTopics,
) ([]*evm.TopicValues, error) {
	var settlementRoundIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.SettlementRoundId).IsZero() {
			settlementRoundIdRule = append(settlementRoundIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.SettlementRoundId)
		if err != nil {
			return nil, err
		}
		settlementRoundIdRule = append(settlementRoundIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		settlementRoundIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeMarketSettled decodes a log into a MarketSettled struct.
func (c *Codec) DecodeMarketSettled(log *evm.Log) (*MarketSettledDecoded, error) {
	event := new(MarketSettledDecoded)
	if err := c.abi.UnpackIntoInterface(event, "MarketSettled", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["MarketSettled"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipHandoverCanceledLogHash() []byte {
	return c.abi.Events["OwnershipHandoverCanceled"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipHandoverCanceledTopics(
	evt abi.Event,
	values []OwnershipHandoverCanceledTopics,
) ([]*evm.TopicValues, error) {
	var pendingOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.PendingOwner).IsZero() {
			pendingOwnerRule = append(pendingOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.PendingOwner)
		if err != nil {
			return nil, err
		}
		pendingOwnerRule = append(pendingOwnerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		pendingOwnerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipHandoverCanceled decodes a log into a OwnershipHandoverCanceled struct.
func (c *Codec) DecodeOwnershipHandoverCanceled(log *evm.Log) (*OwnershipHandoverCanceledDecoded, error) {
	event := new(OwnershipHandoverCanceledDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipHandoverCanceled", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipHandoverCanceled"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipHandoverRequestedLogHash() []byte {
	return c.abi.Events["OwnershipHandoverRequested"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipHandoverRequestedTopics(
	evt abi.Event,
	values []OwnershipHandoverRequestedTopics,
) ([]*evm.TopicValues, error) {
	var pendingOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.PendingOwner).IsZero() {
			pendingOwnerRule = append(pendingOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.PendingOwner)
		if err != nil {
			return nil, err
		}
		pendingOwnerRule = append(pendingOwnerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		pendingOwnerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipHandoverRequested decodes a log into a OwnershipHandoverRequested struct.
func (c *Codec) DecodeOwnershipHandoverRequested(log *evm.Log) (*OwnershipHandoverRequestedDecoded, error) {
	event := new(OwnershipHandoverRequestedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipHandoverRequested", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipHandoverRequested"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) OwnershipTransferredLogHash() []byte {
	return c.abi.Events["OwnershipTransferred"].ID.Bytes()
}

func (c *Codec) EncodeOwnershipTransferredTopics(
	evt abi.Event,
	values []OwnershipTransferredTopics,
) ([]*evm.TopicValues, error) {
	var oldOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.OldOwner).IsZero() {
			oldOwnerRule = append(oldOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.OldOwner)
		if err != nil {
			return nil, err
		}
		oldOwnerRule = append(oldOwnerRule, fieldVal)
	}
	var newOwnerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.NewOwner).IsZero() {
			newOwnerRule = append(newOwnerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.NewOwner)
		if err != nil {
			return nil, err
		}
		newOwnerRule = append(newOwnerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		oldOwnerRule,
		newOwnerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeOwnershipTransferred decodes a log into a OwnershipTransferred struct.
func (c *Codec) DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error) {
	event := new(OwnershipTransferredDecoded)
	if err := c.abi.UnpackIntoInterface(event, "OwnershipTransferred", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["OwnershipTransferred"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) PositionBurnedLogHash() []byte {
	return c.abi.Events["PositionBurned"].ID.Bytes()
}

func (c *Codec) EncodePositionBurnedTopics(
	evt abi.Event,
	values []PositionBurnedTopics,
) ([]*evm.TopicValues, error) {
	var accountRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Account).IsZero() {
			accountRule = append(accountRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Account)
		if err != nil {
			return nil, err
		}
		accountRule = append(accountRule, fieldVal)
	}
	var isLongRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.IsLong).IsZero() {
			isLongRule = append(isLongRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.IsLong)
		if err != nil {
			return nil, err
		}
		isLongRule = append(isLongRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		accountRule,
		isLongRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodePositionBurned decodes a log into a PositionBurned struct.
func (c *Codec) DecodePositionBurned(log *evm.Log) (*PositionBurnedDecoded, error) {
	event := new(PositionBurnedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "PositionBurned", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["PositionBurned"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) PositionMintedLogHash() []byte {
	return c.abi.Events["PositionMinted"].ID.Bytes()
}

func (c *Codec) EncodePositionMintedTopics(
	evt abi.Event,
	values []PositionMintedTopics,
) ([]*evm.TopicValues, error) {
	var accountRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Account).IsZero() {
			accountRule = append(accountRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Account)
		if err != nil {
			return nil, err
		}
		accountRule = append(accountRule, fieldVal)
	}
	var isLongRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.IsLong).IsZero() {
			isLongRule = append(isLongRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.IsLong)
		if err != nil {
			return nil, err
		}
		isLongRule = append(isLongRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		accountRule,
		isLongRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodePositionMinted decodes a log into a PositionMinted struct.
func (c *Codec) DecodePositionMinted(log *evm.Log) (*PositionMintedDecoded, error) {
	event := new(PositionMintedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "PositionMinted", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["PositionMinted"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) PositionRedeemedLogHash() []byte {
	return c.abi.Events["PositionRedeemed"].ID.Bytes()
}

func (c *Codec) EncodePositionRedeemedTopics(
	evt abi.Event,
	values []PositionRedeemedTopics,
) ([]*evm.TopicValues, error) {
	var accountRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Account).IsZero() {
			accountRule = append(accountRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Account)
		if err != nil {
			return nil, err
		}
		accountRule = append(accountRule, fieldVal)
	}
	var isLongRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.IsLong).IsZero() {
			isLongRule = append(isLongRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.IsLong)
		if err != nil {
			return nil, err
		}
		isLongRule = append(isLongRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		accountRule,
		isLongRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodePositionRedeemed decodes a log into a PositionRedeemed struct.
func (c *Codec) DecodePositionRedeemed(log *evm.Log) (*PositionRedeemedDecoded, error) {
	event := new(PositionRedeemedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "PositionRedeemed", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["PositionRedeemed"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c Market) CollateralToken(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeCollateralTokenMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeCollateralTokenMethodOutput(response.Data)
	})

}

func (c Market) Decimals(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[uint8] {
	calldata, err := c.Codec.EncodeDecimalsMethodCall()
	if err != nil {
		return cre.PromiseFromResult[uint8](*new(uint8), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (uint8, error) {
		return c.Codec.DecodeDecimalsMethodOutput(response.Data)
	})

}

func (c Market) LongPositionToken(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeLongPositionTokenMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeLongPositionTokenMethodOutput(response.Data)
	})

}

func (c Market) Oracle(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOracleMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOracleMethodOutput(response.Data)
	})

}

func (c Market) Owner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOwnerMethodOutput(response.Data)
	})

}

func (c Market) OwnershipHandoverExpiresAt(
	runtime cre.Runtime,
	args OwnershipHandoverExpiresAtInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeOwnershipHandoverExpiresAtMethodCall(args)
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
		return c.Codec.DecodeOwnershipHandoverExpiresAtMethodOutput(response.Data)
	})

}

func (c Market) Price(
	runtime cre.Runtime,
	args PriceInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodePriceMethodCall(args)
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
		return c.Codec.DecodePriceMethodOutput(response.Data)
	})

}

func (c Market) Settled(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[bool] {
	calldata, err := c.Codec.EncodeSettledMethodCall()
	if err != nil {
		return cre.PromiseFromResult[bool](*new(bool), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (bool, error) {
		return c.Codec.DecodeSettledMethodOutput(response.Data)
	})

}

func (c Market) SettlementRoundId(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeSettlementRoundIdMethodCall()
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
		return c.Codec.DecodeSettlementRoundIdMethodOutput(response.Data)
	})

}

func (c Market) ShortPositionToken(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeShortPositionTokenMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeShortPositionTokenMethodOutput(response.Data)
	})

}

func (c Market) WriteReport(
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

// DecodeAlreadyInitializedError decodes a AlreadyInitialized error from revert data.
func (c *Market) DecodeAlreadyInitializedError(data []byte) (*AlreadyInitialized, error) {
	args := c.ABI.Errors["AlreadyInitialized"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &AlreadyInitialized{}, nil
}

// Error implements the error interface for AlreadyInitialized.
func (e *AlreadyInitialized) Error() string {
	return fmt.Sprintf("AlreadyInitialized error:")
}

// DecodeAlreadySeededError decodes a AlreadySeeded error from revert data.
func (c *Market) DecodeAlreadySeededError(data []byte) (*AlreadySeeded, error) {
	args := c.ABI.Errors["AlreadySeeded"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &AlreadySeeded{}, nil
}

// Error implements the error interface for AlreadySeeded.
func (e *AlreadySeeded) Error() string {
	return fmt.Sprintf("AlreadySeeded error:")
}

// DecodeAlreadySettledError decodes a AlreadySettled error from revert data.
func (c *Market) DecodeAlreadySettledError(data []byte) (*AlreadySettled, error) {
	args := c.ABI.Errors["AlreadySettled"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &AlreadySettled{}, nil
}

// Error implements the error interface for AlreadySettled.
func (e *AlreadySettled) Error() string {
	return fmt.Sprintf("AlreadySettled error:")
}

// DecodeInsufficientCollateralError decodes a InsufficientCollateral error from revert data.
func (c *Market) DecodeInsufficientCollateralError(data []byte) (*InsufficientCollateral, error) {
	args := c.ABI.Errors["InsufficientCollateral"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &InsufficientCollateral{}, nil
}

// Error implements the error interface for InsufficientCollateral.
func (e *InsufficientCollateral) Error() string {
	return fmt.Sprintf("InsufficientCollateral error:")
}

// DecodeInsufficientPositionTokensError decodes a InsufficientPositionTokens error from revert data.
func (c *Market) DecodeInsufficientPositionTokensError(data []byte) (*InsufficientPositionTokens, error) {
	args := c.ABI.Errors["InsufficientPositionTokens"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &InsufficientPositionTokens{}, nil
}

// Error implements the error interface for InsufficientPositionTokens.
func (e *InsufficientPositionTokens) Error() string {
	return fmt.Sprintf("InsufficientPositionTokens error:")
}

// DecodeNewOwnerIsZeroAddressError decodes a NewOwnerIsZeroAddress error from revert data.
func (c *Market) DecodeNewOwnerIsZeroAddressError(data []byte) (*NewOwnerIsZeroAddress, error) {
	args := c.ABI.Errors["NewOwnerIsZeroAddress"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NewOwnerIsZeroAddress{}, nil
}

// Error implements the error interface for NewOwnerIsZeroAddress.
func (e *NewOwnerIsZeroAddress) Error() string {
	return fmt.Sprintf("NewOwnerIsZeroAddress error:")
}

// DecodeNoHandoverRequestError decodes a NoHandoverRequest error from revert data.
func (c *Market) DecodeNoHandoverRequestError(data []byte) (*NoHandoverRequest, error) {
	args := c.ABI.Errors["NoHandoverRequest"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NoHandoverRequest{}, nil
}

// Error implements the error interface for NoHandoverRequest.
func (e *NoHandoverRequest) Error() string {
	return fmt.Sprintf("NoHandoverRequest error:")
}

// DecodeNotSettledError decodes a NotSettled error from revert data.
func (c *Market) DecodeNotSettledError(data []byte) (*NotSettled, error) {
	args := c.ABI.Errors["NotSettled"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &NotSettled{}, nil
}

// Error implements the error interface for NotSettled.
func (e *NotSettled) Error() string {
	return fmt.Sprintf("NotSettled error:")
}

// DecodeUnauthorizedError decodes a Unauthorized error from revert data.
func (c *Market) DecodeUnauthorizedError(data []byte) (*Unauthorized, error) {
	args := c.ABI.Errors["Unauthorized"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 0 {
		return nil, fmt.Errorf("expected 0 values, got %d", len(values))
	}

	return &Unauthorized{}, nil
}

// Error implements the error interface for Unauthorized.
func (e *Unauthorized) Error() string {
	return fmt.Sprintf("Unauthorized error:")
}

func (c *Market) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["AlreadyInitialized"].ID.Bytes()[:4]):
		return c.DecodeAlreadyInitializedError(data)
	case common.Bytes2Hex(c.ABI.Errors["AlreadySeeded"].ID.Bytes()[:4]):
		return c.DecodeAlreadySeededError(data)
	case common.Bytes2Hex(c.ABI.Errors["AlreadySettled"].ID.Bytes()[:4]):
		return c.DecodeAlreadySettledError(data)
	case common.Bytes2Hex(c.ABI.Errors["InsufficientCollateral"].ID.Bytes()[:4]):
		return c.DecodeInsufficientCollateralError(data)
	case common.Bytes2Hex(c.ABI.Errors["InsufficientPositionTokens"].ID.Bytes()[:4]):
		return c.DecodeInsufficientPositionTokensError(data)
	case common.Bytes2Hex(c.ABI.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]):
		return c.DecodeNewOwnerIsZeroAddressError(data)
	case common.Bytes2Hex(c.ABI.Errors["NoHandoverRequest"].ID.Bytes()[:4]):
		return c.DecodeNoHandoverRequestError(data)
	case common.Bytes2Hex(c.ABI.Errors["NotSettled"].ID.Bytes()[:4]):
		return c.DecodeNotSettledError(data)
	case common.Bytes2Hex(c.ABI.Errors["Unauthorized"].ID.Bytes()[:4]):
		return c.DecodeUnauthorizedError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// MarketSeededTrigger wraps the raw log trigger and provides decoded MarketSeededDecoded data
type MarketSeededTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into MarketSeeded data
func (t *MarketSeededTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MarketSeededDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMarketSeeded(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MarketSeeded log: %w", err)
	}

	return &bindings.DecodedLog[MarketSeededDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerMarketSeededLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MarketSeededTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MarketSeededDecoded]], error) {
	event := c.ABI.Events["MarketSeeded"]
	topics, err := c.Codec.EncodeMarketSeededTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MarketSeeded: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MarketSeededTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsMarketSeeded(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MarketSeededLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// MarketSettledTrigger wraps the raw log trigger and provides decoded MarketSettledDecoded data
type MarketSettledTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into MarketSettled data
func (t *MarketSettledTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[MarketSettledDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeMarketSettled(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode MarketSettled log: %w", err)
	}

	return &bindings.DecodedLog[MarketSettledDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerMarketSettledLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []MarketSettledTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[MarketSettledDecoded]], error) {
	event := c.ABI.Events["MarketSettled"]
	topics, err := c.Codec.EncodeMarketSettledTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for MarketSettled: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &MarketSettledTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsMarketSettled(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.MarketSettledLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipHandoverCanceledTrigger wraps the raw log trigger and provides decoded OwnershipHandoverCanceledDecoded data
type OwnershipHandoverCanceledTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipHandoverCanceled data
func (t *OwnershipHandoverCanceledTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipHandoverCanceledDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipHandoverCanceled(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipHandoverCanceled log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipHandoverCanceledDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerOwnershipHandoverCanceledLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipHandoverCanceledTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipHandoverCanceledDecoded]], error) {
	event := c.ABI.Events["OwnershipHandoverCanceled"]
	topics, err := c.Codec.EncodeOwnershipHandoverCanceledTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipHandoverCanceled: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipHandoverCanceledTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsOwnershipHandoverCanceled(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipHandoverCanceledLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipHandoverRequestedTrigger wraps the raw log trigger and provides decoded OwnershipHandoverRequestedDecoded data
type OwnershipHandoverRequestedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipHandoverRequested data
func (t *OwnershipHandoverRequestedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipHandoverRequestedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipHandoverRequested(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipHandoverRequested log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipHandoverRequestedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerOwnershipHandoverRequestedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipHandoverRequestedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipHandoverRequestedDecoded]], error) {
	event := c.ABI.Events["OwnershipHandoverRequested"]
	topics, err := c.Codec.EncodeOwnershipHandoverRequestedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipHandoverRequested: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipHandoverRequestedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsOwnershipHandoverRequested(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipHandoverRequestedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// OwnershipTransferredTrigger wraps the raw log trigger and provides decoded OwnershipTransferredDecoded data
type OwnershipTransferredTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into OwnershipTransferred data
func (t *OwnershipTransferredTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[OwnershipTransferredDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeOwnershipTransferred(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode OwnershipTransferred log: %w", err)
	}

	return &bindings.DecodedLog[OwnershipTransferredDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
	event := c.ABI.Events["OwnershipTransferred"]
	topics, err := c.Codec.EncodeOwnershipTransferredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for OwnershipTransferred: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &OwnershipTransferredTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.OwnershipTransferredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// PositionBurnedTrigger wraps the raw log trigger and provides decoded PositionBurnedDecoded data
type PositionBurnedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into PositionBurned data
func (t *PositionBurnedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[PositionBurnedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodePositionBurned(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PositionBurned log: %w", err)
	}

	return &bindings.DecodedLog[PositionBurnedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerPositionBurnedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []PositionBurnedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[PositionBurnedDecoded]], error) {
	event := c.ABI.Events["PositionBurned"]
	topics, err := c.Codec.EncodePositionBurnedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for PositionBurned: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &PositionBurnedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsPositionBurned(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.PositionBurnedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// PositionMintedTrigger wraps the raw log trigger and provides decoded PositionMintedDecoded data
type PositionMintedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into PositionMinted data
func (t *PositionMintedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[PositionMintedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodePositionMinted(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PositionMinted log: %w", err)
	}

	return &bindings.DecodedLog[PositionMintedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerPositionMintedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []PositionMintedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[PositionMintedDecoded]], error) {
	event := c.ABI.Events["PositionMinted"]
	topics, err := c.Codec.EncodePositionMintedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for PositionMinted: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &PositionMintedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsPositionMinted(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.PositionMintedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// PositionRedeemedTrigger wraps the raw log trigger and provides decoded PositionRedeemedDecoded data
type PositionRedeemedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Market // Keep reference for decoding
}

// Adapt method that decodes the log into PositionRedeemed data
func (t *PositionRedeemedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[PositionRedeemedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodePositionRedeemed(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PositionRedeemed log: %w", err)
	}

	return &bindings.DecodedLog[PositionRedeemedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Market) LogTriggerPositionRedeemedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []PositionRedeemedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[PositionRedeemedDecoded]], error) {
	event := c.ABI.Events["PositionRedeemed"]
	topics, err := c.Codec.EncodePositionRedeemedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for PositionRedeemed: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &PositionRedeemedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Market) FilterLogsPositionRedeemed(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.PositionRedeemedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}
