// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {PositionToken} from "src/PositionToken.sol";

/// @title Market
/// @author @aadeexyz
/// @notice A binary prediction market backed by collateral tokens with long/short positions
contract Market is IMarket, Ownable {
    /*//////////////////////////////////////////////////////////////
                           TYPE DECLARATIONS
    //////////////////////////////////////////////////////////////*/
    using SafeTransferLib for address;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    IERC20 public immutable collateralToken;
    IOracle public immutable oracle;

    PositionToken public immutable longPositionToken;
    PositionToken public immutable shortPositionToken;

    uint256 private _amountInLongPosition;
    uint256 private _amountInShortPosition;

    bool public settled;
    uint256 public settlementRoundId;

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/
    constructor(address collateralToken_, address oracle_, string memory name_, string memory symbol_) {
        collateralToken = IERC20(collateralToken_);
        oracle = IOracle(oracle_);

        uint8 collateralTokenDecimals = collateralToken.decimals();
        longPositionToken = new PositionToken(
            string(abi.encodePacked("Long ", name_)), string(abi.encodePacked("L", symbol_)), collateralTokenDecimals
        );
        shortPositionToken = new PositionToken(
            string(abi.encodePacked("Short ", name_)), string(abi.encodePacked("S", symbol_)), collateralTokenDecimals
        );

        _initializeOwner(msg.sender);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Seeds the market with initial collateral split between long and short
    /// @param longCollateral_ Amount of collateral allocated to the long pool
    /// @param shortCollateral_ Amount of collateral allocated to the short pool
    function seed(uint256 longCollateral_, uint256 shortCollateral_) external override onlyOwner {
        if (_amountInLongPosition != 0 || _amountInShortPosition != 0) {
            revert AlreadySeeded();
        }

        uint256 totalCollateral = longCollateral_ + shortCollateral_;
        if (totalCollateral == 0) {
            revert InsufficientCollateral();
        }

        address(collateralToken).safeTransferFrom(msg.sender, address(this), totalCollateral);

        _amountInLongPosition = longCollateral_;
        _amountInShortPosition = shortCollateral_;

        longPositionToken.mint(msg.sender, totalCollateral);
        shortPositionToken.mint(msg.sender, totalCollateral);

        emit MarketSeeded(msg.sender, longCollateral_, shortCollateral_, totalCollateral);
    }

    /// @notice Settles the market at a specific oracle round
    /// @param settlementRoundId_ The oracle round ID to settle at
    function settle(uint256 settlementRoundId_) external override onlyOwner {
        if (settled) {
            revert AlreadySettled();
        }

        settled = true;
        settlementRoundId = settlementRoundId_;

        emit MarketSettled(settlementRoundId_);
    }

    /// @notice Redeems position tokens for collateral after settlement
    /// @param isLong_ Whether to redeem long (true) or short (false) position tokens
    function redeem(bool isLong_) external override {
        if (!settled) {
            revert NotSettled();
        }

        uint256 d = collateralToken.decimals();
        uint8 od = oracle.decimals();
        IOracle.Round memory round = oracle.getRound(settlementRoundId);
        uint256 index = round.index;

        if (isLong_) {
            uint256 positionTokenAmount = longPositionToken.balanceOf(msg.sender);
            if (positionTokenAmount == 0) {
                revert InsufficientPositionTokens();
            }

            uint256 redeemPrice = FixedPointMathLib.mulDiv(index, 10 ** d, 10 ** od);
            uint256 collateralAmount = FixedPointMathLib.mulDiv(positionTokenAmount, redeemPrice, 10 ** d);

            longPositionToken.burn(msg.sender, positionTokenAmount);
            address(collateralToken).safeTransfer(msg.sender, collateralAmount);

            emit PositionRedeemed(msg.sender, true, positionTokenAmount, collateralAmount);
        } else {
            uint256 positionTokenAmount = shortPositionToken.balanceOf(msg.sender);
            if (positionTokenAmount == 0) {
                revert InsufficientPositionTokens();
            }

            uint256 redeemPrice = 10 ** d - FixedPointMathLib.mulDiv(index, 10 ** d, 10 ** od);
            uint256 collateralAmount = FixedPointMathLib.mulDiv(positionTokenAmount, redeemPrice, 10 ** d);

            shortPositionToken.burn(msg.sender, positionTokenAmount);
            address(collateralToken).safeTransfer(msg.sender, collateralAmount);

            emit PositionRedeemed(msg.sender, false, positionTokenAmount, collateralAmount);
        }
    }

    /// @notice Mints position tokens by depositing collateral
    /// @param isLong_ Whether to mint long (true) or short (false) position tokens
    /// @param collateralAmount_ Amount of collateral to deposit
    function mint(bool isLong_, uint256 collateralAmount_) external override {
        if (settled) {
            revert AlreadySettled();
        }
        if (collateralAmount_ == 0) {
            revert InsufficientCollateral();
        }

        uint256 positionTokenPrice = price(isLong_);
        uint256 positionTokenAmount =
            FixedPointMathLib.mulDiv(collateralAmount_, 10 ** collateralToken.decimals(), positionTokenPrice);

        if (positionTokenAmount == 0) {
            revert InsufficientPositionTokens();
        }

        address(collateralToken).safeTransferFrom(msg.sender, address(this), collateralAmount_);

        if (isLong_) {
            _amountInLongPosition += collateralAmount_;
            longPositionToken.mint(msg.sender, positionTokenAmount);
        } else {
            _amountInShortPosition += collateralAmount_;
            shortPositionToken.mint(msg.sender, positionTokenAmount);
        }

        emit PositionMinted(msg.sender, isLong_, collateralAmount_, positionTokenAmount);
    }

    /// @notice Burns position tokens and withdraws collateral
    /// @param isLong_ Whether to burn long (true) or short (false) position tokens
    /// @param positionTokenAmount_ Amount of position tokens to burn
    function burn(bool isLong_, uint256 positionTokenAmount_) external override {
        if (settled) {
            revert AlreadySettled();
        }
        if (positionTokenAmount_ == 0) {
            revert InsufficientPositionTokens();
        }

        uint256 positionTokenPrice = price(isLong_);
        uint256 collateralAmount =
            FixedPointMathLib.mulDiv(positionTokenAmount_, positionTokenPrice, 10 ** collateralToken.decimals());

        if (collateralAmount == 0) {
            revert InsufficientCollateral();
        }

        if (isLong_) {
            _amountInLongPosition -= collateralAmount;
            longPositionToken.burn(msg.sender, positionTokenAmount_);
        } else {
            _amountInShortPosition -= collateralAmount;
            shortPositionToken.burn(msg.sender, positionTokenAmount_);
        }

        emit PositionBurned(msg.sender, isLong_, positionTokenAmount_, collateralAmount);

        address(collateralToken).safeTransfer(msg.sender, collateralAmount);
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the collateral token decimals
    function decimals() external view override returns (uint8) {
        return collateralToken.decimals();
    }

    /// @notice Returns the current price of a position token
    /// @param isLong_ Whether to get the long (true) or short (false) price
    /// @return The price in collateral token units
    function price(bool isLong_) public view override returns (uint256) {
        uint256 totalAmount = _amountInLongPosition + _amountInShortPosition;
        if (totalAmount == 0) {
            return (10 ** collateralToken.decimals()) / 2;
        }

        if (isLong_) {
            return FixedPointMathLib.mulDivUp(_amountInLongPosition, 10 ** collateralToken.decimals(), totalAmount);
        } else {
            return FixedPointMathLib.mulDivUp(_amountInShortPosition, 10 ** collateralToken.decimals(), totalAmount);
        }
    }
}
