package lens

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var LensMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getWorkflowData\",\"inputs\":[{\"name\":\"fof_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"factories\",\"type\":\"tuple[]\",\"internalType\":\"structILens.WorkflowFactoryData[]\",\"components\":[{\"name\":\"factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"latestMarket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rollingEMAWindow\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"}]",
}

type WorkflowFactoryData struct {
	Factory           common.Address
	Name              string
	Oracle            common.Address
	LatestMarket      common.Address
	CollateralBalance *big.Int
	RollingEMAWindow  []*big.Int
}

type GetWorkflowDataInput struct {
	Fof common.Address
}

type GetWorkflowDataOutput struct {
	CollateralToken common.Address
	Factories       []WorkflowFactoryData
}

type Lens struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   LensCodec
}

type LensCodec interface {
	EncodeGetWorkflowDataMethodCall(in GetWorkflowDataInput) ([]byte, error)
	DecodeGetWorkflowDataMethodOutput(data []byte) (GetWorkflowDataOutput, error)
}

type Codec struct {
	abi *abi.ABI
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

func NewCodec() (LensCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(LensMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
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
		return GetWorkflowDataOutput{}, fmt.Errorf("expected 2 outputs, got %d", len(vals))
	}

	var result GetWorkflowDataOutput

	collateralTokenJSON, err := json.Marshal(vals[0])
	if err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to marshal collateralToken: %w", err)
	}
	if err := json.Unmarshal(collateralTokenJSON, &result.CollateralToken); err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to unmarshal collateralToken: %w", err)
	}

	factoriesJSON, err := json.Marshal(vals[1])
	if err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to marshal factories: %w", err)
	}
	if err := json.Unmarshal(factoriesJSON, &result.Factories); err != nil {
		return GetWorkflowDataOutput{}, fmt.Errorf("failed to unmarshal factories: %w", err)
	}

	return result, nil
}

func (c Lens) GetWorkflowData(
	runtime cre.Runtime,
	args GetWorkflowDataInput,
	blockNumber *big.Int,
) cre.Promise[GetWorkflowDataOutput] {
	calldata, err := c.Codec.EncodeGetWorkflowDataMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult(GetWorkflowDataOutput{}, err)
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
