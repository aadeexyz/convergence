// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/// @title IMarketFactory
/// @author @aadeexyz
/// @notice Interface for market factories that create and manage attention markets
interface IMarketFactory {
    /*//////////////////////////////////////////////////////////////
                                  ENUMS
    //////////////////////////////////////////////////////////////*/
    enum State {
        Creating,
        Live
    }

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/
    event MarketCreated(address indexed market);
    event MarketSeeded(address indexed market, uint256 longCollateral, uint256 shortCollateral);
    event StateUpdated(State oldState, State newState);

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    function collateralToken() external view returns (address);

    function oracle() external view returns (address);

    function latestMarket() external view returns (address);

    function markets(uint256 index) external view returns (address);

    function totalMarkets() external view returns (uint256);

    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function decimals() external view returns (uint8);

    function state() external view returns (State);
}
