// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {LibString} from "solady/utils/LibString.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {IFactoryOfMarketFactories} from "src/interfaces/IFactoryOfMarketFactories.sol";
import {MarketFactory} from "src/MarketFactory.sol";

contract FactoryOfMarketFactories is IFactoryOfMarketFactories, Ownable {
    using SafeTransferLib for address;
    using LibString for string;

    address public immutable collateralToken;
    uint256 private _liquidityFee;
    uint256 private _protocolFee;

    address private _forwarderAddress;
    uint8 private _oracleDecimals;

    address[] private _marketFactories;
    mapping(string => address) private _marketFactoriesExists;

    constructor(
        address collateralToken_,
        uint256 liquidityFee_,
        uint256 protocolFee_,
        address forwarderAddress_,
        uint8 oracleDecimals_,
        address owner_
    ) {
        collateralToken = collateralToken_;
        _liquidityFee = liquidityFee_;
        _protocolFee = protocolFee_;
        _forwarderAddress = forwarderAddress_;
        _oracleDecimals = oracleDecimals_;

        _initializeOwner(owner_);
    }

    function setLiquidityFee(uint256 liquidityFee_) external override onlyOwner {
        uint256 oldFee = _liquidityFee;
        _liquidityFee = liquidityFee_;
        emit LiquidityFeeUpdated(oldFee, liquidityFee_);
    }

    function setProtocolFee(uint256 protocolFee_) external override onlyOwner {
        uint256 oldFee = _protocolFee;
        _protocolFee = protocolFee_;
        emit ProtocolFeeUpdated(oldFee, protocolFee_);
    }

    function createMarketFactory(string memory name_, string memory symbol_) external override returns (address) {
        if (bytes(name_).length == 0) {
            revert InvalidName();
        }
        if (bytes(symbol_).length == 0) {
            revert InvalidSymbol();
        }

        string memory name = name_.lower();
        string memory symbol = symbol_.upper();

        if (_marketFactoriesExists[name] != address(0)) {
            revert MarketFactoryAlreadyExists(_marketFactoriesExists[name]);
        }

        address marketFactory = address(new MarketFactory(collateralToken, _forwarderAddress, _oracleDecimals, name, symbol));

        _marketFactories.push(marketFactory);
        _marketFactoriesExists[name] = marketFactory;

        collateralToken.safeTransferFrom(msg.sender, marketFactory, _liquidityFee);
        collateralToken.safeTransferFrom(msg.sender, owner(), _protocolFee);

        emit MarketFactoryCreated(marketFactory, name, symbol);

        return marketFactory;
    }

    function creationFee() external view override returns (uint256) {
        return _liquidityFee + _protocolFee;
    }

    function totalMarketFactories() external view override returns (uint256) {
        return _marketFactories.length;
    }

    function marketFactories() external view override returns (address[] memory) {
        return _marketFactories;
    }

    function marketFactory(uint256 index_) external view override returns (address) {
        return _marketFactories[index_];
    }

}
