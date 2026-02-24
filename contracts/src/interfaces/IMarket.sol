// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {PositionToken} from "src/PositionToken.sol";

interface IMarket {
    error InsufficientCollateral();

    error InsufficientPositionTokens();

    event PositionMinted(address indexed account, bool indexed isLong, uint256 collateralAmount, uint256 positionTokenAmount);
    event PositionBurned(address indexed account, bool indexed isLong, uint256 positionTokenAmount, uint256 collateralAmount);

    function collateralToken() external view returns (IERC20);

    function longPositionToken() external view returns (PositionToken);

    function shortPositionToken() external view returns (PositionToken);

    function price(bool isLong_) external view returns (uint256);

    function decimals() external view returns (uint8);

    function mint(bool isLong_, uint256 collateralAmount_) external;

    function burn(bool isLong_, uint256 positionTokenAmount_) external;
}
