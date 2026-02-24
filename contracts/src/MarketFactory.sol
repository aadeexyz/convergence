// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {Oracle} from "src/Oracle.sol";
import {Market} from "src/Market.sol";

contract MarketFactory is IMarketFactory, Ownable {
    using SafeTransferLib for address;

    address public immutable collateralToken;
    address public immutable oracle;

    address[] public markets;

    string public name;
    string public symbol;

    State public state;

    constructor(
        address collateralToken_,
        address forwarderAddress_,
        uint8 oracleDecimals_,
        string memory name_,
        string memory symbol_
    ) {
        collateralToken = collateralToken_;
        name = name_;
        symbol = symbol_;
        state = State.Creating;

        oracle = address(new Oracle(oracleDecimals_, name_, forwarderAddress_, msg.sender));

        _initializeOwner(msg.sender);
    }

    function latestMarket() external view override returns (address) {
        return markets[markets.length - 1];
    }

    function decimals() external view override returns (uint8) {
        return IERC20(collateralToken).decimals();
    }

    function totalMarkets() external view override returns (uint256) {
        return markets.length;
    }

    function _createMarket() internal virtual returns (address) {
        address market = address(new Market(collateralToken, name, symbol));
        markets.push(market);

        emit MarketCreated(market);

        return market;
    }

    function _seedMarket() private {
        address market = _createMarket();

        IOracle.Round memory latestRound = IOracle(oracle).getLatestRound();
        uint256 index = latestRound.index;
        uint8 oracleDecimals = IOracle(oracle).decimals();

        uint256 collateral = IERC20(collateralToken).balanceOf(address(this));
        uint256 longCollateral = FixedPointMathLib.mulDiv(collateral, index, 10 ** oracleDecimals);
        uint256 shortCollateral = collateral - longCollateral;

        collateralToken.safeApprove(market, collateral);
        IMarket(market).seed(longCollateral, shortCollateral);
    }
}
