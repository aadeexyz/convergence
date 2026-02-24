// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {PositionToken} from "src/PositionToken.sol";

// WAM - World Attention Market

contract Market is IMarket, Ownable {
    using SafeTransferLib for address;

    IERC20 public immutable collateralToken;

    PositionToken public immutable longPositionToken;
    PositionToken public immutable shortPositionToken;

    uint256 private _amountInLongPosition;
    uint256 private _amountInShortPosition;

    constructor(address collateralToken_, string memory name_, string memory symbol_) {
        collateralToken = IERC20(collateralToken_);

        uint8 collateralTokenDecimals = collateralToken.decimals();
        longPositionToken = new PositionToken(
            string(abi.encodePacked("Long ", name_)), string(abi.encodePacked("L", symbol_)), collateralTokenDecimals
        );
        shortPositionToken = new PositionToken(
            string(abi.encodePacked("Short ", name_)), string(abi.encodePacked("S", symbol_)), collateralTokenDecimals
        );

        _initializeOwner(msg.sender);
    }

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

    function mint(bool isLong_, uint256 collateralAmount_) external override {
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

    function burn(bool isLong_, uint256 positionTokenAmount_) external override {
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

    function decimals() external view override returns (uint8) {
        return collateralToken.decimals();
    }

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
