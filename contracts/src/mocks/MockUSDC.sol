// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ERC20} from "solady/tokens/ERC20.sol";

/// @title MockUSDC
/// @author @aadeexyz
/// @notice Mock USDC token for testing purposes
contract MockUSDC is ERC20 {
    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the token name
    function name() public pure override returns (string memory) {
        return "USD Coin";
    }

    /// @notice Returns the token symbol
    function symbol() public pure override returns (string memory) {
        return "USDC";
    }

    /// @notice Returns the token decimals
    function decimals() public pure override returns (uint8) {
        return 6;
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Mints tokens to an address (open mint for testing)
    /// @param to_ The address to mint tokens to
    /// @param amount_ The amount of tokens to mint
    function mint(address to_, uint256 amount_) external {
        _mint(to_, amount_);
    }
}
