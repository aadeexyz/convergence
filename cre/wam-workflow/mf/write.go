package mf

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	mfBinding "cre-workflow/contracts/evm/src/generated/market_factory"
	oracleBinding "cre-workflow/contracts/evm/src/generated/oracle"
)

func SubmitAttentionIndex(
	runtime cre.Runtime,
	chainName string,
	factoryAddress common.Address,
	index *big.Int,
	ema *big.Int,
	gasLimit uint64,
) error {
	logger := runtime.Logger()

	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contract, err := mfBinding.NewMarketFactory(evmClient, factoryAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to create MarketFactory binding: %w", err)
	}

	// ABI-encode (uint256 index, uint256 ema) matching _processReport(bytes calldata report_)
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return fmt.Errorf("failed to create uint256 type: %w", err)
	}

	args := abi.Arguments{
		{Type: uint256Type},
		{Type: uint256Type},
	}

	encodedPayload, err := args.Pack(index, ema)
	if err != nil {
		return fmt.Errorf("failed to encode report payload: %w", err)
	}

	logger.Info("Generating report",
		"factory", factoryAddress.Hex(),
		"index", index.String(),
		"ema", ema.String(),
	)

	report, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: encodedPayload,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await()
	if err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}

	gasConfig := &evm.GasConfig{GasLimit: gasLimit}
	resp, err := contract.WriteReport(runtime, report, gasConfig).Await()
	if err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	if resp.TxStatus != evm.TxStatus_TX_STATUS_SUCCESS {
		errorMsg := "unknown error"
		if resp.ErrorMessage != nil {
			errorMsg = *resp.ErrorMessage
		}
		return fmt.Errorf("transaction failed with status %v: %s", resp.TxStatus, errorMsg)
	}

	txHash := fmt.Sprintf("0x%x", resp.TxHash)
	logger.Info("Report submitted successfully",
		"factory", factoryAddress.Hex(),
		"txHash", txHash,
		"index", index.String(),
		"ema", ema.String(),
	)

	return nil
}

func SubmitOracleIndex(
	runtime cre.Runtime,
	chainName string,
	oracleAddress common.Address,
	index *big.Int,
	ema *big.Int,
	gasLimit uint64,
) error {
	logger := runtime.Logger()

	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	contract, err := oracleBinding.NewOracle(evmClient, oracleAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to create Oracle binding: %w", err)
	}

	// ABI-encode (uint256 index, uint256 ema) matching _processReport(bytes calldata report_)
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return fmt.Errorf("failed to create uint256 type: %w", err)
	}

	args := abi.Arguments{
		{Type: uint256Type},
		{Type: uint256Type},
	}

	encodedPayload, err := args.Pack(index, ema)
	if err != nil {
		return fmt.Errorf("failed to encode report payload: %w", err)
	}

	logger.Info("Generating oracle report",
		"oracle", oracleAddress.Hex(),
		"index", index.String(),
		"ema", ema.String(),
	)

	report, err := runtime.GenerateReport(&cre.ReportRequest{
		EncodedPayload: encodedPayload,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}).Await()
	if err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}

	gasConfig := &evm.GasConfig{GasLimit: gasLimit}
	resp, err := contract.WriteReport(runtime, report, gasConfig).Await()
	if err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	if resp.TxStatus != evm.TxStatus_TX_STATUS_SUCCESS {
		errorMsg := "unknown error"
		if resp.ErrorMessage != nil {
			errorMsg = *resp.ErrorMessage
		}
		return fmt.Errorf("transaction failed with status %v: %s", resp.TxStatus, errorMsg)
	}

	txHash := fmt.Sprintf("0x%x", resp.TxHash)
	logger.Info("Oracle report submitted successfully",
		"oracle", oracleAddress.Hex(),
		"txHash", txHash,
		"index", index.String(),
		"ema", ema.String(),
	)

	return nil
}
