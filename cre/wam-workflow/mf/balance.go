package mf

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/cre"

	erc20Binding "cre-workflow/contracts/evm/src/generated/erc20"
)

func GetCollateralBalance(
	runtime cre.Runtime,
	chainName string,
	tokenAddress common.Address,
	account common.Address,
) (*big.Int, error) {
	chainSelector, err := evm.ChainSelectorFromName(chainName)
	if err != nil {
		return nil, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	token, err := erc20Binding.NewERC20(evmClient, tokenAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ERC20 binding: %w", err)
	}

	balance, err := token.BalanceOf(runtime, erc20Binding.BalanceOfInput{
		Account: account,
	}, big.NewInt(-2)).Await()
	if err != nil {
		return nil, fmt.Errorf("failed to read balanceOf: %w", err)
	}

	return balance, nil
}
