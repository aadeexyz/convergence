// Code generated — DO NOT EDIT.

package oracle

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

var OracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"keyword_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"currentRoundId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOracle.Round\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRound\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIOracle.Round\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyword\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rollingEMAWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRound\",\"inputs\":[{\"name\":\"index_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ema_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AnswerSubmitted\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"ema\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// Structs
type IOracleRound struct {
	Id        *big.Int
	Timestamp *big.Int
	Index     *big.Int
}

// Contract Method Inputs
type CompleteOwnershipHandoverInput struct {
	PendingOwner common.Address
}

type GetRoundInput struct {
	Id *big.Int
}

type OwnershipHandoverExpiresAtInput struct {
	PendingOwner common.Address
}

type SubmitRoundInput struct {
	Index *big.Int
	Ema   *big.Int
}

type TransferOwnershipInput struct {
	NewOwner common.Address
}

// Contract Method Outputs

// Errors
type AlreadyInitialized struct {
}

type NewOwnerIsZeroAddress struct {
}

type NoHandoverRequest struct {
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

type AnswerSubmittedTopics struct {
	RoundId   *big.Int
	Timestamp *big.Int
}

type AnswerSubmittedDecoded struct {
	RoundId   *big.Int
	Timestamp *big.Int
	Index     *big.Int
	Ema       *big.Int
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

// Main Binding Type for Oracle
type Oracle struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   OracleCodec
}

type OracleCodec interface {
	EncodeCancelOwnershipHandoverMethodCall() ([]byte, error)
	EncodeCompleteOwnershipHandoverMethodCall(in CompleteOwnershipHandoverInput) ([]byte, error)
	EncodeCurrentRoundIdMethodCall() ([]byte, error)
	DecodeCurrentRoundIdMethodOutput(data []byte) (*big.Int, error)
	EncodeDecimalsMethodCall() ([]byte, error)
	DecodeDecimalsMethodOutput(data []byte) (uint8, error)
	EncodeGetLatestRoundMethodCall() ([]byte, error)
	DecodeGetLatestRoundMethodOutput(data []byte) (IOracleRound, error)
	EncodeGetRoundMethodCall(in GetRoundInput) ([]byte, error)
	DecodeGetRoundMethodOutput(data []byte) (IOracleRound, error)
	EncodeKeywordMethodCall() ([]byte, error)
	DecodeKeywordMethodOutput(data []byte) (string, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeOwnershipHandoverExpiresAtMethodCall(in OwnershipHandoverExpiresAtInput) ([]byte, error)
	DecodeOwnershipHandoverExpiresAtMethodOutput(data []byte) (*big.Int, error)
	EncodeRenounceOwnershipMethodCall() ([]byte, error)
	EncodeRequestOwnershipHandoverMethodCall() ([]byte, error)
	EncodeRollingEMAWindowMethodCall() ([]byte, error)
	DecodeRollingEMAWindowMethodOutput(data []byte) ([]*big.Int, error)
	EncodeSubmitRoundMethodCall(in SubmitRoundInput) ([]byte, error)
	EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error)
	EncodeIOracleRoundStruct(in IOracleRound) ([]byte, error)
	AnswerSubmittedLogHash() []byte
	EncodeAnswerSubmittedTopics(evt abi.Event, values []AnswerSubmittedTopics) ([]*evm.TopicValues, error)
	DecodeAnswerSubmitted(log *evm.Log) (*AnswerSubmittedDecoded, error)
	OwnershipHandoverCanceledLogHash() []byte
	EncodeOwnershipHandoverCanceledTopics(evt abi.Event, values []OwnershipHandoverCanceledTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipHandoverCanceled(log *evm.Log) (*OwnershipHandoverCanceledDecoded, error)
	OwnershipHandoverRequestedLogHash() []byte
	EncodeOwnershipHandoverRequestedTopics(evt abi.Event, values []OwnershipHandoverRequestedTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipHandoverRequested(log *evm.Log) (*OwnershipHandoverRequestedDecoded, error)
	OwnershipTransferredLogHash() []byte
	EncodeOwnershipTransferredTopics(evt abi.Event, values []OwnershipTransferredTopics) ([]*evm.TopicValues, error)
	DecodeOwnershipTransferred(log *evm.Log) (*OwnershipTransferredDecoded, error)
}

func NewOracle(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &Oracle{
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

func NewCodec() (OracleCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeCancelOwnershipHandoverMethodCall() ([]byte, error) {
	return c.abi.Pack("cancelOwnershipHandover")
}

func (c *Codec) EncodeCompleteOwnershipHandoverMethodCall(in CompleteOwnershipHandoverInput) ([]byte, error) {
	return c.abi.Pack("completeOwnershipHandover", in.PendingOwner)
}

func (c *Codec) EncodeCurrentRoundIdMethodCall() ([]byte, error) {
	return c.abi.Pack("currentRoundId")
}

func (c *Codec) DecodeCurrentRoundIdMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["currentRoundId"].Outputs.Unpack(data)
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

func (c *Codec) EncodeGetLatestRoundMethodCall() ([]byte, error) {
	return c.abi.Pack("getLatestRound")
}

func (c *Codec) DecodeGetLatestRoundMethodOutput(data []byte) (IOracleRound, error) {
	vals, err := c.abi.Methods["getLatestRound"].Outputs.Unpack(data)
	if err != nil {
		return *new(IOracleRound), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(IOracleRound), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result IOracleRound
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(IOracleRound), fmt.Errorf("failed to unmarshal to IOracleRound: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetRoundMethodCall(in GetRoundInput) ([]byte, error) {
	return c.abi.Pack("getRound", in.Id)
}

func (c *Codec) DecodeGetRoundMethodOutput(data []byte) (IOracleRound, error) {
	vals, err := c.abi.Methods["getRound"].Outputs.Unpack(data)
	if err != nil {
		return *new(IOracleRound), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(IOracleRound), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result IOracleRound
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(IOracleRound), fmt.Errorf("failed to unmarshal to IOracleRound: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeKeywordMethodCall() ([]byte, error) {
	return c.abi.Pack("keyword")
}

func (c *Codec) DecodeKeywordMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["keyword"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(string), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result string
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(string), fmt.Errorf("failed to unmarshal to string: %w", err)
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

func (c *Codec) EncodeRenounceOwnershipMethodCall() ([]byte, error) {
	return c.abi.Pack("renounceOwnership")
}

func (c *Codec) EncodeRequestOwnershipHandoverMethodCall() ([]byte, error) {
	return c.abi.Pack("requestOwnershipHandover")
}

func (c *Codec) EncodeRollingEMAWindowMethodCall() ([]byte, error) {
	return c.abi.Pack("rollingEMAWindow")
}

func (c *Codec) DecodeRollingEMAWindowMethodOutput(data []byte) ([]*big.Int, error) {
	vals, err := c.abi.Methods["rollingEMAWindow"].Outputs.Unpack(data)
	if err != nil {
		return *new([]*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []*big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]*big.Int), fmt.Errorf("failed to unmarshal to []*big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeSubmitRoundMethodCall(in SubmitRoundInput) ([]byte, error) {
	return c.abi.Pack("submitRound", in.Index, in.Ema)
}

func (c *Codec) EncodeTransferOwnershipMethodCall(in TransferOwnershipInput) ([]byte, error) {
	return c.abi.Pack("transferOwnership", in.NewOwner)
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

func (c *Codec) AnswerSubmittedLogHash() []byte {
	return c.abi.Events["AnswerSubmitted"].ID.Bytes()
}

func (c *Codec) EncodeAnswerSubmittedTopics(
	evt abi.Event,
	values []AnswerSubmittedTopics,
) ([]*evm.TopicValues, error) {
	var roundIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.RoundId).IsZero() {
			roundIdRule = append(roundIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.RoundId)
		if err != nil {
			return nil, err
		}
		roundIdRule = append(roundIdRule, fieldVal)
	}
	var timestampRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Timestamp).IsZero() {
			timestampRule = append(timestampRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Timestamp)
		if err != nil {
			return nil, err
		}
		timestampRule = append(timestampRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		roundIdRule,
		timestampRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeAnswerSubmitted decodes a log into a AnswerSubmitted struct.
func (c *Codec) DecodeAnswerSubmitted(log *evm.Log) (*AnswerSubmittedDecoded, error) {
	event := new(AnswerSubmittedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "AnswerSubmitted", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["AnswerSubmitted"].Inputs {
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

func (c Oracle) CurrentRoundId(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeCurrentRoundIdMethodCall()
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
		return c.Codec.DecodeCurrentRoundIdMethodOutput(response.Data)
	})

}

func (c Oracle) Decimals(
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

func (c Oracle) GetLatestRound(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[IOracleRound] {
	calldata, err := c.Codec.EncodeGetLatestRoundMethodCall()
	if err != nil {
		return cre.PromiseFromResult[IOracleRound](*new(IOracleRound), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (IOracleRound, error) {
		return c.Codec.DecodeGetLatestRoundMethodOutput(response.Data)
	})

}

func (c Oracle) GetRound(
	runtime cre.Runtime,
	args GetRoundInput,
	blockNumber *big.Int,
) cre.Promise[IOracleRound] {
	calldata, err := c.Codec.EncodeGetRoundMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[IOracleRound](*new(IOracleRound), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (IOracleRound, error) {
		return c.Codec.DecodeGetRoundMethodOutput(response.Data)
	})

}

func (c Oracle) Keyword(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[string] {
	calldata, err := c.Codec.EncodeKeywordMethodCall()
	if err != nil {
		return cre.PromiseFromResult[string](*new(string), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) (string, error) {
		return c.Codec.DecodeKeywordMethodOutput(response.Data)
	})

}

func (c Oracle) Owner(
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

func (c Oracle) OwnershipHandoverExpiresAt(
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

func (c Oracle) RollingEMAWindow(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[[]*big.Int] {
	calldata, err := c.Codec.EncodeRollingEMAWindowMethodCall()
	if err != nil {
		return cre.PromiseFromResult[[]*big.Int](*new([]*big.Int), err)
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
	return cre.Then(promise, func(response *evm.CallContractReply) ([]*big.Int, error) {
		return c.Codec.DecodeRollingEMAWindowMethodOutput(response.Data)
	})

}

func (c Oracle) WriteReportFromIOracleRound(
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

func (c Oracle) WriteReport(
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
func (c *Oracle) DecodeAlreadyInitializedError(data []byte) (*AlreadyInitialized, error) {
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

// DecodeNewOwnerIsZeroAddressError decodes a NewOwnerIsZeroAddress error from revert data.
func (c *Oracle) DecodeNewOwnerIsZeroAddressError(data []byte) (*NewOwnerIsZeroAddress, error) {
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
func (c *Oracle) DecodeNoHandoverRequestError(data []byte) (*NoHandoverRequest, error) {
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

// DecodeUnauthorizedError decodes a Unauthorized error from revert data.
func (c *Oracle) DecodeUnauthorizedError(data []byte) (*Unauthorized, error) {
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

func (c *Oracle) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["AlreadyInitialized"].ID.Bytes()[:4]):
		return c.DecodeAlreadyInitializedError(data)
	case common.Bytes2Hex(c.ABI.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]):
		return c.DecodeNewOwnerIsZeroAddressError(data)
	case common.Bytes2Hex(c.ABI.Errors["NoHandoverRequest"].ID.Bytes()[:4]):
		return c.DecodeNoHandoverRequestError(data)
	case common.Bytes2Hex(c.ABI.Errors["Unauthorized"].ID.Bytes()[:4]):
		return c.DecodeUnauthorizedError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// AnswerSubmittedTrigger wraps the raw log trigger and provides decoded AnswerSubmittedDecoded data
type AnswerSubmittedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]         // Embed the raw trigger
	contract                        *Oracle // Keep reference for decoding
}

// Adapt method that decodes the log into AnswerSubmitted data
func (t *AnswerSubmittedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[AnswerSubmittedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeAnswerSubmitted(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode AnswerSubmitted log: %w", err)
	}

	return &bindings.DecodedLog[AnswerSubmittedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *Oracle) LogTriggerAnswerSubmittedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []AnswerSubmittedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[AnswerSubmittedDecoded]], error) {
	event := c.ABI.Events["AnswerSubmitted"]
	topics, err := c.Codec.EncodeAnswerSubmittedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for AnswerSubmitted: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &AnswerSubmittedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *Oracle) FilterLogsAnswerSubmitted(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.AnswerSubmittedLogHash()}},
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
	contract                        *Oracle // Keep reference for decoding
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

func (c *Oracle) LogTriggerOwnershipHandoverCanceledLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipHandoverCanceledTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipHandoverCanceledDecoded]], error) {
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

func (c *Oracle) FilterLogsOwnershipHandoverCanceled(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
	contract                        *Oracle // Keep reference for decoding
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

func (c *Oracle) LogTriggerOwnershipHandoverRequestedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipHandoverRequestedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipHandoverRequestedDecoded]], error) {
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

func (c *Oracle) FilterLogsOwnershipHandoverRequested(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
	contract                        *Oracle // Keep reference for decoding
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

func (c *Oracle) LogTriggerOwnershipTransferredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []OwnershipTransferredTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[OwnershipTransferredDecoded]], error) {
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

func (c *Oracle) FilterLogsOwnershipTransferred(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
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
