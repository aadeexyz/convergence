// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {Market} from "src/Market.sol";

contract MarketFactory is IMarketFactory, Ownable {
    address public immutable collateralToken;

    address[] public markets;

    string public name;
    string public symbol;

    constructor(address collateralToken_, string memory name_, string memory symbol_) {
        collateralToken = collateralToken_;
        name = name_;
        symbol = symbol_;

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

    function createMarket() external override onlyOwner returns (address) {
        return _createMarket();
    }

    function seedMarket() external override onlyOwner {}

    function _createMarket() internal virtual returns (address) {
        address market = address(new Market(collateralToken, name, symbol));
        markets.push(market);

        emit MarketCreated(market);

        return market;
    }
}
