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

    /// @notice Returns the collateral token address
    function collateralToken() external view returns (address);

    /// @notice Returns the oracle address
    function oracle() external view returns (address);

    /// @notice Returns the most recently created market address
    function latestMarket() external view returns (address);

    /// @notice Returns the market address at the given index
    /// @param index The index in the markets array
    function markets(uint256 index) external view returns (address);

    /// @notice Returns the total number of markets created
    function totalMarkets() external view returns (uint256);

    /// @notice Returns the market factory name
    function name() external view returns (string memory);

    /// @notice Returns the market factory symbol
    function symbol() external view returns (string memory);

    /// @notice Returns the collateral token decimals
    function decimals() external view returns (uint8);

    /// @notice Returns the current factory state
    function state() external view returns (State);
}
