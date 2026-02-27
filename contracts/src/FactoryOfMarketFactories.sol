// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {LibString} from "solady/utils/LibString.sol";
import {LibClone} from "solady/utils/LibClone.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {IFactoryOfMarketFactories} from "src/interfaces/IFactoryOfMarketFactories.sol";
import {MarketFactory} from "src/MarketFactory.sol";
import {Market} from "src/Market.sol";
import {Oracle} from "src/Oracle.sol";
import {PositionToken} from "src/PositionToken.sol";

/// @title FactoryOfMarketFactories
/// @author @aadeexyz
/// @notice Factory contract that deploys implementation contracts once and clones MarketFactory instances
contract FactoryOfMarketFactories is IFactoryOfMarketFactories, Ownable {
    using LibClone for address;
    using SafeTransferLib for address;
    using LibString for string;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    address public immutable collateralToken;
    uint256 public liquidityFee;
    uint256 public protocolFee;

    address public forwarderAddress;
    uint8 public oracleDecimals;

    address public immutable marketFactoryImpl;
    address public immutable marketImpl;
    address public immutable oracleImpl;
    address public immutable positionTokenImpl;

    address[] private _marketFactories;
    mapping(string => address) private _marketFactoriesExists;

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/
    constructor(
        address collateralToken_,
        uint256 liquidityFee_,
        uint256 protocolFee_,
        address forwarderAddress_,
        uint8 oracleDecimals_,
        address owner_
    ) {
        collateralToken = collateralToken_;
        liquidityFee = liquidityFee_;
        protocolFee = protocolFee_;
        forwarderAddress = forwarderAddress_;
        oracleDecimals = oracleDecimals_;

        marketFactoryImpl = address(new MarketFactory());
        marketImpl = address(new Market());
        oracleImpl = address(new Oracle());
        positionTokenImpl = address(new PositionToken());

        _initializeOwner(owner_);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Updates the liquidity fee charged on market factory creation
    /// @param liquidityFee_ The new liquidity fee amount
    function setLiquidityFee(uint256 liquidityFee_) external override onlyOwner {
        uint256 oldFee = liquidityFee;
        liquidityFee = liquidityFee_;
        emit LiquidityFeeUpdated(oldFee, liquidityFee_);
    }

    /// @notice Updates the protocol fee charged on market factory creation
    /// @param protocolFee_ The new protocol fee amount
    function setProtocolFee(uint256 protocolFee_) external override onlyOwner {
        uint256 oldFee = protocolFee;
        protocolFee = protocolFee_;
        emit ProtocolFeeUpdated(oldFee, protocolFee_);
    }

    /// @notice Creates a new MarketFactory clone with the given name and symbol
    /// @param name_ The name for the market factory (will be lowercased)
    /// @param symbol_ The symbol for the market factory (will be uppercased)
    /// @return The address of the newly created MarketFactory clone
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

        address mf = marketFactoryImpl.clone(
            abi.encode(collateralToken, oracleDecimals, name, symbol, marketImpl, oracleImpl, positionTokenImpl)
        );

        _marketFactories.push(mf);
        _marketFactoriesExists[name] = mf;

        collateralToken.safeTransferFrom(msg.sender, mf, liquidityFee);
        collateralToken.safeTransferFrom(msg.sender, owner(), protocolFee);

        MarketFactory(mf).initialize(forwarderAddress, address(this));

        emit MarketFactoryCreated(mf, name, symbol);

        return mf;
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the total creation fee (liquidity + protocol)
    function creationFee() external view override returns (uint256) {
        return liquidityFee + protocolFee;
    }

    /// @notice Returns the total number of market factories created
    function totalMarketFactories() external view override returns (uint256) {
        return _marketFactories.length;
    }

    /// @notice Returns all market factory addresses
    function marketFactories() external view override returns (address[] memory) {
        return _marketFactories;
    }

    /// @notice Returns the market factory address at a given index
    /// @param index_ The index in the market factories array
    function marketFactory(uint256 index_) external view override returns (address) {
        return _marketFactories[index_];
    }
}
