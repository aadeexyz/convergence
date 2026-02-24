// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {PositionToken} from "src/PositionToken.sol";

interface IMarket {
    error InsufficientCollateral();
    error InsufficientPositionTokens();
    error AlreadySeeded();
    error AlreadySettled();
    error NotSettled();

    event PositionMinted(
        address indexed account, bool indexed isLong, uint256 collateralAmount, uint256 positionTokenAmount
    );
    event PositionBurned(
        address indexed account, bool indexed isLong, uint256 positionTokenAmount, uint256 collateralAmount
    );
    event MarketSeeded(
        address indexed seeder, uint256 longCollateral, uint256 shortCollateral, uint256 positionTokenAmount
    );
    event MarketSettled(uint256 indexed settlementRoundId);
    event PositionRedeemed(
        address indexed account, bool indexed isLong, uint256 positionTokenAmount, uint256 collateralAmount
    );

    function collateralToken() external view returns (IERC20);

    function longPositionToken() external view returns (PositionToken);

    function shortPositionToken() external view returns (PositionToken);

    function price(bool isLong_) external view returns (uint256);

    function decimals() external view returns (uint8);

    function seed(uint256 longCollateral_, uint256 shortCollateral_) external;

    function mint(bool isLong_, uint256 collateralAmount_) external;

    function burn(bool isLong_, uint256 positionTokenAmount_) external;

    function settle(uint256 settlementRoundId_) external;

    function redeem(bool isLong_) external;

    function settled() external view returns (bool);

    function settlementRoundId() external view returns (uint256);
}
