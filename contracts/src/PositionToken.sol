// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ERC20} from "solady/tokens/ERC20.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {LibClone} from "solady/utils/LibClone.sol";

/// @title PositionToken
/// @author @aadeexyz
/// @notice ERC20 token representing a long or short position in an attention market
/// @dev Deployed as a clone with immutable args: abi.encode(string name, string symbol, uint8 decimals)
contract PositionToken is ERC20, Ownable {
    using LibClone for address;
    
    /*//////////////////////////////////////////////////////////////
                              INITIALIZER
    //////////////////////////////////////////////////////////////*/

    /// @notice Initializes the clone (can only be called once)
    /// @param owner_ The owner of this position token (the Market contract)
    function initialize(address owner_) external {
        _initializeOwner(owner_);
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the token name from immutable args
    function name() public view override returns (string memory) {
        (string memory name_,,) = _args();
        return name_;
    }

    /// @notice Returns the token symbol from immutable args
    function symbol() public view override returns (string memory) {
        (, string memory symbol_,) = _args();
        return symbol_;
    }

    /// @notice Returns the token decimals from immutable args
    function decimals() public view override returns (uint8) {
        (,, uint8 decimals_) = _args();
        return decimals_;
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Mints tokens to an address (owner only)
    /// @param to_ The address to mint tokens to
    /// @param amount_ The amount of tokens to mint
    function mint(address to_, uint256 amount_) external onlyOwner {
        _mint(to_, amount_);
    }

    /// @notice Burns tokens from an address (owner only)
    /// @param from_ The address to burn tokens from
    /// @param amount_ The amount of tokens to burn
    function burn(address from_, uint256 amount_) external onlyOwner {
        _burn(from_, amount_);
    }

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Decodes the immutable args appended to this clone
    function _args() private view returns (string memory, string memory, uint8) {
        return abi.decode(address(this).argsOnClone(), (string, string, uint8));
    }
}
