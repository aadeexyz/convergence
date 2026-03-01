// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {IRouter} from "src/interfaces/IRouter.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {PositionToken} from "src/PositionToken.sol";

/// @title Router
/// @author @aadeexyz
/// @notice Universal router that handles all user-facing market operations across factories
/// @dev Users approve collateral to this contract once; position tokens are approved per-token as needed
contract Router is IRouter {
    using SafeTransferLib for address;

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IRouter
    function mint(address factory_, bool isLong_, uint256 collateralAmount_) external override {
        IMarketFactory factory = IMarketFactory(factory_);
        if (factory.state() != IMarketFactory.State.Live) {
            revert MarketNotLive();
        }

        address market = factory.latestMarket();
        address collateralToken = factory.collateralToken();

        // Transfer collateral from user to this contract
        collateralToken.safeTransferFrom(msg.sender, address(this), collateralAmount_);

        // Approve market and mint directly to user
        collateralToken.safeApprove(market, collateralAmount_);
        IMarket(market).mint(isLong_, collateralAmount_, msg.sender);

        emit Minted(factory_, market, msg.sender, isLong_, collateralAmount_);
    }

    /// @inheritdoc IRouter
    function burn(address factory_, bool isLong_, uint256 positionTokenAmount_) external override {
        IMarketFactory factory = IMarketFactory(factory_);
        if (factory.state() != IMarketFactory.State.Live) {
            revert MarketNotLive();
        }

        address market = factory.latestMarket();

        // Transfer position tokens from user to this contract
        PositionToken positionToken = isLong_
            ? IMarket(market).longPositionToken()
            : IMarket(market).shortPositionToken();
        address(positionToken).safeTransferFrom(msg.sender, address(this), positionTokenAmount_);

        // Burn position tokens, collateral goes directly to user
        IMarket(market).burn(isLong_, positionTokenAmount_, msg.sender);

        emit Burned(factory_, market, msg.sender, isLong_, positionTokenAmount_);
    }

    /// @inheritdoc IRouter
    function redeem(address factory_, uint256 marketIndex_, bool isLong_) external override {
        IMarketFactory factory = IMarketFactory(factory_);
        address market = factory.markets(marketIndex_);

        _redeem(market, isLong_);

        emit Redeemed(factory_, market, msg.sender, isLong_, 0);
    }

    /// @inheritdoc IRouter
    function redeemAll(address factory_) external override {
        IMarketFactory factory = IMarketFactory(factory_);
        uint256 totalMarkets = factory.totalMarkets();

        if (totalMarkets == 0) {
            revert NoSettledMarkets();
        }

        for (uint256 i = 0; i < totalMarkets; i++) {
            address market = factory.markets(i);

            if (!IMarket(market).settled()) {
                continue;
            }

            _redeem(market, true);
            _redeem(market, false);
        }
    }

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Redeems position tokens from a settled market, collateral goes directly to user
    function _redeem(address market_, bool isLong_) private {
        if (!IMarket(market_).settled()) {
            return;
        }

        PositionToken positionToken = isLong_
            ? IMarket(market_).longPositionToken()
            : IMarket(market_).shortPositionToken();

        uint256 userBalance = positionToken.balanceOf(msg.sender);
        if (userBalance == 0) {
            return;
        }

        // Transfer position tokens from user to this contract
        address(positionToken).safeTransferFrom(msg.sender, address(this), userBalance);

        // Redeem — collateral goes directly to user via recipient
        IMarket(market_).redeem(isLong_, msg.sender);
    }
}
