// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IMarket} from "src/interfaces/IMarket.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IOracle} from "src/interfaces/IOracle.sol";

/// @title ILens
/// @author @aadeexyz
/// @notice Interface for read-only views into factories, markets, oracles, and redeemable positions
interface ILens {
    /*//////////////////////////////////////////////////////////////
                                 STRUCTS
    //////////////////////////////////////////////////////////////*/

    struct FactoryView {
        string name;
        string symbol;
        address oracle;
        address collateralToken;
        uint8 decimals;
        IMarketFactory.State state;
        uint256 totalMarkets;
        address latestMarket;
    }

    struct MarketView {
        address market;
        address longPositionToken;
        address shortPositionToken;
        uint256 longPrice;
        uint256 shortPrice;
        uint256 totalLiquidity;
        uint8 decimals;
        bool settled;
        uint256 settlementRoundId;
    }

    struct OracleView {
        uint256 currentRoundId;
        IOracle.Round latestRound;
        uint8 decimals;
        string keyword;
        uint256[] rollingEMAWindow;
    }

    struct RedeemablePosition {
        address market;
        uint256 marketIndex;
        bool isLong;
        uint256 positionTokenBalance;
        uint256 redeemableCollateral;
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns a full view of a market factory
    /// @param factory_ The market factory address
    function getFactory(address factory_) external view returns (FactoryView memory);

    /// @notice Returns a full view of a market
    /// @param market_ The market address
    function getMarket(address market_) external view returns (MarketView memory);

    /// @notice Returns a full view of the oracle for a factory
    /// @param factory_ The market factory address
    function getOracle(address factory_) external view returns (OracleView memory);

    /// @notice Returns the latest oracle price for a factory
    /// @param factory_ The market factory address
    /// @return index The attention index value
    /// @return timestamp The timestamp of the latest round
    function getOraclePrice(address factory_) external view returns (uint256 index, uint256 timestamp);

    /// @notice Returns the current implied market price for a position
    /// @param factory_ The market factory address
    /// @param isLong_ Whether to get the long (true) or short (false) price
    function getMarketPrice(address factory_, bool isLong_) external view returns (uint256);

    /// @notice Returns a full view of the latest market in a factory
    /// @param factory_ The market factory address
    function getLatestMarket(address factory_) external view returns (MarketView memory);

    /// @notice Returns full views of all markets in a factory
    /// @param factory_ The market factory address
    function getAllMarkets(address factory_) external view returns (MarketView[] memory);

    /// @notice Returns all redeemable positions for an account across all settled markets in a factory
    /// @param factory_ The market factory address
    /// @param account_ The account to check
    function getRedeemable(address factory_, address account_) external view returns (RedeemablePosition[] memory);

    /// @notice Returns the redeemable collateral amount for a specific position
    /// @param market_ The market address
    /// @param account_ The account to check
    /// @param isLong_ Whether to check long (true) or short (false) position
    function getRedeemableForMarket(address market_, address account_, bool isLong_) external view returns (uint256);

    /// @notice Returns all oracle rounds for a factory
    /// @param factory_ The market factory address
    function getAllRounds(address factory_) external view returns (IOracle.Round[] memory);

    /// @notice Returns price snapshots from the latest market in a factory
    /// @param factory_ The market factory address
    /// @param count_ Max number of most recent snapshots to return (0 = all)
    function getMarkPriceHistory(address factory_, uint256 count_) external view returns (IMarket.PriceSnapshot[] memory);
}
