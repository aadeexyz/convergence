// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

interface IMarketFactory {
    event MarketCreated(address indexed market);

    function createMarket() external returns (address);

    function seedMarket() external;

    function collateralToken() external view returns (address);

    function latestMarket() external view returns (address);

    function markets(uint256 index) external view returns (address);

    function totalMarkets() external view returns (uint256);

    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function decimals() external view returns (uint8);
}
