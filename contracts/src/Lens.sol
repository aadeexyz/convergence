// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {ILens} from "src/interfaces/ILens.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {PositionToken} from "src/PositionToken.sol";

/// @title Lens
/// @author @aadeexyz
/// @notice Read-only views into factories, markets, oracles, and redeemable positions
contract Lens is ILens {
    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc ILens
    function getFactory(address factory_) external view override returns (FactoryView memory) {
        IMarketFactory f = IMarketFactory(factory_);
        return FactoryView({
            name: f.name(),
            symbol: f.symbol(),
            oracle: f.oracle(),
            collateralToken: f.collateralToken(),
            decimals: f.decimals(),
            state: f.state(),
            totalMarkets: f.totalMarkets(),
            latestMarket: f.latestMarket()
        });
    }

    /// @inheritdoc ILens
    function getMarket(address market_) external view override returns (MarketView memory) {
        return _buildMarketView(market_);
    }

    /// @inheritdoc ILens
    function getOracle(address factory_) external view override returns (OracleView memory) {
        IOracle o = IOracle(IMarketFactory(factory_).oracle());
        uint256 roundId = o.currentRoundId();
        IOracle.Round memory latestRound;
        if (roundId > 0) {
            latestRound = o.getLatestRound();
        }
        return OracleView({
            currentRoundId: roundId,
            latestRound: latestRound,
            decimals: o.decimals(),
            keyword: o.keyword(),
            rollingEMAWindow: o.rollingEMAWindow()
        });
    }

    /// @inheritdoc ILens
    function getOraclePrice(address factory_) external view override returns (uint256 index, uint256 timestamp) {
        IMarketFactory f = IMarketFactory(factory_);
        IOracle o = IOracle(f.oracle());
        uint256 roundId = o.currentRoundId();
        if (roundId == 0) {
            return (0, 0);
        }
        IOracle.Round memory round = o.getLatestRound();
        return (round.index, round.timestamp);
    }

    /// @inheritdoc ILens
    function getMarketPrice(address factory_, bool isLong_) external view override returns (uint256) {
        IMarketFactory f = IMarketFactory(factory_);
        address market = f.latestMarket();
        if (market == address(0)) {
            return 0;
        }
        return IMarket(market).price(isLong_);
    }

    /// @inheritdoc ILens
    function getLatestMarket(address factory_) external view override returns (MarketView memory) {
        IMarketFactory f = IMarketFactory(factory_);
        address market = f.latestMarket();
        if (market == address(0)) {
            return MarketView({
                market: address(0),
                longPositionToken: address(0),
                shortPositionToken: address(0),
                longPrice: 0,
                shortPrice: 0,
                totalLiquidity: 0,
                decimals: 0,
                settled: false,
                settlementRoundId: 0
            });
        }
        return _buildMarketView(market);
    }

    /// @inheritdoc ILens
    function getAllMarkets(address factory_) external view override returns (MarketView[] memory) {
        IMarketFactory f = IMarketFactory(factory_);
        uint256 total = f.totalMarkets();
        MarketView[] memory views = new MarketView[](total);
        for (uint256 i = 0; i < total; i++) {
            views[i] = _buildMarketView(f.markets(i));
        }
        return views;
    }

    /// @inheritdoc ILens
    function getRedeemable(address factory_, address account_)
        external
        view
        override
        returns (RedeemablePosition[] memory)
    {
        IMarketFactory f = IMarketFactory(factory_);
        uint256 totalMarkets = f.totalMarkets();

        // First pass: count redeemable positions
        uint256 count = 0;
        for (uint256 i = 0; i < totalMarkets; i++) {
            address market = f.markets(i);
            if (!IMarket(market).settled()) {
                continue;
            }
            if (IMarket(market).longPositionToken().balanceOf(account_) > 0) {
                count++;
            }
            if (IMarket(market).shortPositionToken().balanceOf(account_) > 0) {
                count++;
            }
        }

        // Second pass: populate
        RedeemablePosition[] memory positions = new RedeemablePosition[](count);
        uint256 idx = 0;
        for (uint256 i = 0; i < totalMarkets; i++) {
            address market = f.markets(i);
            if (!IMarket(market).settled()) {
                continue;
            }

            uint256 longBalance = IMarket(market).longPositionToken().balanceOf(account_);
            if (longBalance > 0) {
                positions[idx] = RedeemablePosition({
                    market: market,
                    marketIndex: i,
                    isLong: true,
                    positionTokenBalance: longBalance,
                    redeemableCollateral: _calculateRedeemable(market, longBalance, true)
                });
                idx++;
            }

            uint256 shortBalance = IMarket(market).shortPositionToken().balanceOf(account_);
            if (shortBalance > 0) {
                positions[idx] = RedeemablePosition({
                    market: market,
                    marketIndex: i,
                    isLong: false,
                    positionTokenBalance: shortBalance,
                    redeemableCollateral: _calculateRedeemable(market, shortBalance, false)
                });
                idx++;
            }
        }

        return positions;
    }

    /// @inheritdoc ILens
    function getRedeemableForMarket(address market_, address account_, bool isLong_)
        external
        view
        override
        returns (uint256)
    {
        if (!IMarket(market_).settled()) {
            return 0;
        }

        PositionToken positionToken =
            isLong_ ? IMarket(market_).longPositionToken() : IMarket(market_).shortPositionToken();

        uint256 balance = positionToken.balanceOf(account_);
        if (balance == 0) {
            return 0;
        }

        return _calculateRedeemable(market_, balance, isLong_);
    }

    /// @inheritdoc ILens
    function getAllRounds(address factory_) external view override returns (IOracle.Round[] memory) {
        IOracle o = IOracle(IMarketFactory(factory_).oracle());
        uint256 total = o.currentRoundId();
        IOracle.Round[] memory rounds = new IOracle.Round[](total);
        for (uint256 i = 1; i <= total; i++) {
            rounds[i - 1] = o.getRound(i);
        }
        return rounds;
    }

    /// @inheritdoc ILens
    function getMarkPriceHistory(address factory_, uint256 count_) external view override returns (IMarket.PriceSnapshot[] memory) {
        IMarketFactory f = IMarketFactory(factory_);
        uint256 total = f.totalMarkets();
        if (total == 0) return new IMarket.PriceSnapshot[](0);

        address latestMarket = f.markets(total - 1);
        IMarket.PriceSnapshot[] memory all = IMarket(latestMarket).priceSnapshots();

        if (count_ == 0 || count_ >= all.length) return all;

        IMarket.PriceSnapshot[] memory result = new IMarket.PriceSnapshot[](count_);
        uint256 start = all.length - count_;
        for (uint256 i = 0; i < count_; i++) {
            result[i] = all[start + i];
        }
        return result;
    }

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Builds a MarketView struct for a given market address
    function _buildMarketView(address market_) private view returns (MarketView memory) {
        IMarket m = IMarket(market_);
        uint256 liquidity = IERC20(address(m.collateralToken())).balanceOf(market_);
        return MarketView({
            market: market_,
            longPositionToken: address(m.longPositionToken()),
            shortPositionToken: address(m.shortPositionToken()),
            longPrice: m.price(true),
            shortPrice: m.price(false),
            totalLiquidity: liquidity,
            decimals: m.decimals(),
            settled: m.settled(),
            settlementRoundId: m.settlementRoundId()
        });
    }

    /// @notice Calculates the collateral amount redeemable for a given position token balance
    function _calculateRedeemable(address market_, uint256 positionTokenBalance_, bool isLong_)
        private
        view
        returns (uint256)
    {
        IMarket m = IMarket(market_);
        IERC20 collateralToken = m.collateralToken();
        IOracle oracle = m.oracle();

        uint256 d = collateralToken.decimals();
        uint8 od = oracle.decimals();
        IOracle.Round memory round = oracle.getRound(m.settlementRoundId());
        uint256 index = round.index;

        uint256 redeemPrice;
        if (isLong_) {
            redeemPrice = FixedPointMathLib.mulDiv(index, 10 ** d, 10 ** od);
        } else {
            redeemPrice = 10 ** d - FixedPointMathLib.mulDiv(index, 10 ** d, 10 ** od);
        }

        return FixedPointMathLib.mulDiv(positionTokenBalance_, redeemPrice, 10 ** d);
    }
}
