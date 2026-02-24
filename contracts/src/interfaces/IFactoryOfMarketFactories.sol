// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

interface IFactoryOfMarketFactories {
    error InvalidName();
    error InvalidSymbol();
    error MarketFactoryAlreadyExists(address marketFactory);

    event MarketFactoryCreated(address indexed marketFactory, string name, string symbol);
    event LiquidityFeeUpdated(uint256 oldFee, uint256 newFee);
    event ProtocolFeeUpdated(uint256 oldFee, uint256 newFee);

    function createMarketFactory(string memory name_, string memory symbol_) external returns (address);

    function setLiquidityFee(uint256 liquidityFee_) external;

    function setProtocolFee(uint256 protocolFee_) external;

    function creationFee() external view returns (uint256);

    function collateralToken() external view returns (address);

    function totalMarketFactories() external view returns (uint256);

    function marketFactories() external view returns (address[] memory);

    function marketFactory(uint256 index_) external view returns (address);
}
