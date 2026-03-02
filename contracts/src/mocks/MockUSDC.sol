// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {ERC20} from "solady/tokens/ERC20.sol";
import {Ownable} from "solady/auth/Ownable.sol";

/// @title MockUSDC
/// @author @aadeexyz
/// @notice Mock USDC token with rate-limited public faucet and owner-only arbitrary mint
contract MockUSDC is ERC20, Ownable {
    uint256 private _mintAmount;
    mapping(address => uint256) private _lastMintedTimestamp;

    error CanMintOnlyOncePerDay(uint256 lastMinted, uint256 timeElapsed);

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(uint256 mintAmount_, address owner_) {
        _mintAmount = mintAmount_;
        _initializeOwner(owner_);
    }

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

    /// @notice Returns the faucet mint amount
    function mintAmount() external view returns (uint256) {
        return _mintAmount;
    }

    /// @notice Returns the last time an address used the faucet
    function lastMintedTimestamp(address account_) external view returns (uint256) {
        return _lastMintedTimestamp[account_];
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Public faucet: anyone can mint once per day
    function mint() external {
        uint256 timeElapsed = block.timestamp - _lastMintedTimestamp[msg.sender];
        if (timeElapsed < 1 days) {
            revert CanMintOnlyOncePerDay(_lastMintedTimestamp[msg.sender], timeElapsed);
        }

        _lastMintedTimestamp[msg.sender] = block.timestamp;
        _mint(msg.sender, _mintAmount);
    }

    /// @notice Owner-only mint for seeding factories
    /// @param to_ The address to mint tokens to
    /// @param amount_ The amount of tokens to mint
    function mint(address to_, uint256 amount_) external onlyOwner {
        _mint(to_, amount_);
    }
}
