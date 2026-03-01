// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/// @title IRouter
/// @author @aadeexyz
/// @notice Interface for the universal router that handles all user-facing market operations
interface IRouter {
    /*//////////////////////////////////////////////////////////////
                                 ERRORS
    //////////////////////////////////////////////////////////////*/
    error MarketNotLive();
    error NoSettledMarkets();

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/
    event Minted(address indexed factory, address indexed market, address indexed account, bool isLong, uint256 collateralAmount);
    event Burned(address indexed factory, address indexed market, address indexed account, bool isLong, uint256 positionTokenAmount);
    event Redeemed(address indexed factory, address indexed market, address indexed account, bool isLong, uint256 collateralAmount);

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Mints position tokens on the latest market of a factory
    /// @param factory_ The market factory address
    /// @param isLong_ Whether to mint long (true) or short (false) position tokens
    /// @param collateralAmount_ Amount of collateral to deposit
    function mint(address factory_, bool isLong_, uint256 collateralAmount_) external;

    /// @notice Burns position tokens on the latest market of a factory
    /// @param factory_ The market factory address
    /// @param isLong_ Whether to burn long (true) or short (false) position tokens
    /// @param positionTokenAmount_ Amount of position tokens to burn
    function burn(address factory_, bool isLong_, uint256 positionTokenAmount_) external;

    /// @notice Redeems position tokens from a specific settled market
    /// @param factory_ The market factory address
    /// @param marketIndex_ The index of the settled market in the factory's markets array
    /// @param isLong_ Whether to redeem long (true) or short (false) position tokens
    function redeem(address factory_, uint256 marketIndex_, bool isLong_) external;

    /// @notice Redeems all position tokens from all settled markets in a factory
    /// @param factory_ The market factory address
    function redeemAll(address factory_) external;
}
