// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {PositionToken} from "src/PositionToken.sol";

/// @title IMarket
/// @author @aadeexyz
/// @notice Interface for binary prediction markets with long/short positions
interface IMarket {
    /*//////////////////////////////////////////////////////////////
                                 STRUCTS
    //////////////////////////////////////////////////////////////*/
    struct PriceSnapshot {
        uint256 timestamp;
        uint256 longPrice;
    }

    /*//////////////////////////////////////////////////////////////
                                 ERRORS
    //////////////////////////////////////////////////////////////*/
    error InsufficientCollateral();
    error InsufficientPositionTokens();
    error AlreadySeeded();
    error AlreadySettled();
    error NotSettled();

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/
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

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the collateral token address
    function collateralToken() external view returns (IERC20);

    /// @notice Returns the oracle address
    function oracle() external view returns (IOracle);

    /// @notice Returns the long position token
    function longPositionToken() external view returns (PositionToken);

    /// @notice Returns the short position token
    function shortPositionToken() external view returns (PositionToken);

    /// @notice Returns the current price of a position token
    /// @param isLong_ Whether to get the long (true) or short (false) price
    /// @return The price in collateral token units
    function price(bool isLong_) external view returns (uint256);

    /// @notice Returns the collateral token decimals
    function decimals() external view returns (uint8);

    /// @notice Whether the market has been settled
    function settled() external view returns (bool);

    /// @notice The oracle round ID used for settlement
    function settlementRoundId() external view returns (uint256);

    /// @notice Returns the remaining long pool collateral after settlement
    function settlementLongPool() external view returns (uint256);

    /// @notice Returns the remaining short pool collateral after settlement
    function settlementShortPool() external view returns (uint256);

    /// @notice Returns all price snapshots recorded after each trade
    function priceSnapshots() external view returns (PriceSnapshot[] memory);

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Seeds the market with initial collateral split between long and short
    /// @param longCollateral_ Amount of collateral allocated to the long pool
    /// @param shortCollateral_ Amount of collateral allocated to the short pool
    function seed(uint256 longCollateral_, uint256 shortCollateral_) external;

    /// @notice Mints position tokens by depositing collateral
    /// @param isLong_ Whether to mint long (true) or short (false) position tokens
    /// @param collateralAmount_ Amount of collateral to deposit
    /// @param recipient_ The address to receive the position tokens
    function mint(bool isLong_, uint256 collateralAmount_, address recipient_) external;

    /// @notice Burns position tokens and withdraws collateral
    /// @param isLong_ Whether to burn long (true) or short (false) position tokens
    /// @param positionTokenAmount_ Amount of position tokens to burn
    /// @param recipient_ The address to receive the collateral
    function burn(bool isLong_, uint256 positionTokenAmount_, address recipient_) external;

    /// @notice Settles the market at a specific oracle round
    /// @param settlementRoundId_ The oracle round ID to settle at
    function settle(uint256 settlementRoundId_) external;

    /// @notice Redeems position tokens for collateral after settlement
    /// @dev Uses pro-rata pool-based redemption with pools snapshotted at settlement.
    ///      Each holder gets their share of their side's pool proportional to
    ///      their token balance vs remaining supply. Pools are decremented on each
    ///      redemption so partial redeems never affect the other side.
    /// @param isLong_ Whether to redeem long (true) or short (false) position tokens
    /// @param recipient_ The address to receive the collateral
    function redeem(bool isLong_, address recipient_) external;
}
