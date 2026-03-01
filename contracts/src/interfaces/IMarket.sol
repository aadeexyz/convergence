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
    function collateralToken() external view returns (IERC20);

    function oracle() external view returns (IOracle);

    function longPositionToken() external view returns (PositionToken);

    function shortPositionToken() external view returns (PositionToken);

    function price(bool isLong_) external view returns (uint256);

    function decimals() external view returns (uint8);

    function settled() external view returns (bool);

    function settlementRoundId() external view returns (uint256);

    function priceSnapshots() external view returns (PriceSnapshot[] memory);

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    function seed(uint256 longCollateral_, uint256 shortCollateral_) external;

    function mint(bool isLong_, uint256 collateralAmount_, address recipient_) external;

    function burn(bool isLong_, uint256 positionTokenAmount_, address recipient_) external;

    function settle(uint256 settlementRoundId_) external;

    function redeem(bool isLong_, address recipient_) external;
}
