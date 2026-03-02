// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {LibClone} from "solady/utils/LibClone.sol";
import {PositionToken} from "src/PositionToken.sol";

/// @title Market
/// @author @aadeexyz
/// @notice A binary prediction market backed by collateral tokens with long/short positions
/// @dev Deployed as a clone with immutable args: abi.encode(address collateralToken, address oracle, string name, string symbol, address positionTokenImpl)
contract Market is IMarket, Ownable {
    using LibClone for address;
    using SafeTransferLib for address;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    PositionToken public longPositionToken;
    PositionToken public shortPositionToken;

    uint256 private _amountInLongPosition;
    uint256 private _amountInShortPosition;

    bool public settled;
    uint256 public settlementRoundId;

    uint256 private _settlementLongPool;
    uint256 private _settlementShortPool;

    IMarket.PriceSnapshot[] private _priceSnapshots;

    /*//////////////////////////////////////////////////////////////
                              INITIALIZER
    //////////////////////////////////////////////////////////////*/

    /// @notice Initializes the clone: creates position token clones and sets owner
    /// @param owner_ The owner of this market (the MarketFactory contract)
    function initialize(address owner_) external {
        _initializeOwner(owner_);

        (address collateralToken_,, string memory name_, string memory symbol_, address positionTokenImpl_) = _args();
        uint8 collateralTokenDecimals = IERC20(collateralToken_).decimals();

        address longClone = positionTokenImpl_.clone(
            abi.encode(
                string(abi.encodePacked("Long ", name_)),
                string(abi.encodePacked("L", symbol_)),
                collateralTokenDecimals
            )
        );
        PositionToken(longClone).initialize(address(this));
        longPositionToken = PositionToken(longClone);

        address shortClone = positionTokenImpl_.clone(
            abi.encode(
                string(abi.encodePacked("Short ", name_)),
                string(abi.encodePacked("S", symbol_)),
                collateralTokenDecimals
            )
        );
        PositionToken(shortClone).initialize(address(this));
        shortPositionToken = PositionToken(shortClone);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IMarket
    function seed(uint256 longCollateral_, uint256 shortCollateral_) external override onlyOwner {
        if (_amountInLongPosition != 0 || _amountInShortPosition != 0) {
            revert AlreadySeeded();
        }

        uint256 totalCollateral = longCollateral_ + shortCollateral_;
        if (totalCollateral == 0) {
            revert InsufficientCollateral();
        }

        _collateralToken().safeTransferFrom(msg.sender, address(this), totalCollateral);

        _amountInLongPosition = longCollateral_;
        _amountInShortPosition = shortCollateral_;

        longPositionToken.mint(msg.sender, totalCollateral);
        shortPositionToken.mint(msg.sender, totalCollateral);

        emit MarketSeeded(msg.sender, longCollateral_, shortCollateral_, totalCollateral);

        _priceSnapshots.push(IMarket.PriceSnapshot({timestamp: block.timestamp, longPrice: price(true)}));
    }

    /// @inheritdoc IMarket
    function settle(uint256 settlementRoundId_) external override onlyOwner {
        if (settled) {
            revert AlreadySettled();
        }

        settled = true;
        settlementRoundId = settlementRoundId_;

        // Snapshot pool allocations using oracle index so partial redemptions don't shift the other side
        uint256 totalCollateral = IERC20(_collateralToken()).balanceOf(address(this));
        uint8 decimals_ = IOracle(_oracle()).decimals();
        uint256 index = IOracle(_oracle()).getRound(settlementRoundId_).index;
        _settlementLongPool = FixedPointMathLib.mulDiv(totalCollateral, index, 10 ** decimals_);
        _settlementShortPool = totalCollateral - _settlementLongPool;

        emit MarketSettled(settlementRoundId_);
    }

    /// @inheritdoc IMarket
    function redeem(bool isLong_, address recipient_) external override {
        if (!settled) {
            revert NotSettled();
        }

        if (isLong_) {
            uint256 positionTokenAmount = longPositionToken.balanceOf(msg.sender);
            if (positionTokenAmount == 0) {
                revert InsufficientPositionTokens();
            }

            uint256 totalSupply = longPositionToken.totalSupply();
            uint256 collateralAmount = FixedPointMathLib.mulDiv(_settlementLongPool, positionTokenAmount, totalSupply);

            _settlementLongPool -= collateralAmount;
            longPositionToken.burn(msg.sender, positionTokenAmount);
            _collateralToken().safeTransfer(recipient_, collateralAmount);

            emit PositionRedeemed(recipient_, true, positionTokenAmount, collateralAmount);
        } else {
            uint256 positionTokenAmount = shortPositionToken.balanceOf(msg.sender);
            if (positionTokenAmount == 0) {
                revert InsufficientPositionTokens();
            }

            uint256 totalSupply = shortPositionToken.totalSupply();
            uint256 collateralAmount = FixedPointMathLib.mulDiv(_settlementShortPool, positionTokenAmount, totalSupply);

            _settlementShortPool -= collateralAmount;
            shortPositionToken.burn(msg.sender, positionTokenAmount);
            _collateralToken().safeTransfer(recipient_, collateralAmount);

            emit PositionRedeemed(recipient_, false, positionTokenAmount, collateralAmount);
        }
    }

    /// @inheritdoc IMarket
    function mint(bool isLong_, uint256 collateralAmount_, address recipient_) external override {
        if (settled) {
            revert AlreadySettled();
        }
        if (collateralAmount_ == 0) {
            revert InsufficientCollateral();
        }

        uint256 positionTokenPrice = price(isLong_);
        uint256 positionTokenAmount = FixedPointMathLib.mulDiv(
            collateralAmount_, 10 ** IERC20(_collateralToken()).decimals(), positionTokenPrice
        );

        if (positionTokenAmount == 0) {
            revert InsufficientPositionTokens();
        }

        _collateralToken().safeTransferFrom(msg.sender, address(this), collateralAmount_);

        if (isLong_) {
            _amountInLongPosition += collateralAmount_;
            longPositionToken.mint(recipient_, positionTokenAmount);
        } else {
            _amountInShortPosition += collateralAmount_;
            shortPositionToken.mint(recipient_, positionTokenAmount);
        }

        emit PositionMinted(recipient_, isLong_, collateralAmount_, positionTokenAmount);

        _priceSnapshots.push(IMarket.PriceSnapshot({timestamp: block.timestamp, longPrice: price(true)}));
    }

    /// @inheritdoc IMarket
    function burn(bool isLong_, uint256 positionTokenAmount_, address recipient_) external override {
        if (settled) {
            revert AlreadySettled();
        }
        if (positionTokenAmount_ == 0) {
            revert InsufficientPositionTokens();
        }

        uint256 positionTokenPrice = price(isLong_);
        uint256 collateralAmount = FixedPointMathLib.mulDiv(
            positionTokenAmount_, positionTokenPrice, 10 ** IERC20(_collateralToken()).decimals()
        );

        if (collateralAmount == 0) {
            revert InsufficientCollateral();
        }

        if (isLong_) {
            _amountInLongPosition -= collateralAmount;
            longPositionToken.burn(msg.sender, positionTokenAmount_);
        } else {
            _amountInShortPosition -= collateralAmount;
            shortPositionToken.burn(msg.sender, positionTokenAmount_);
        }

        emit PositionBurned(recipient_, isLong_, positionTokenAmount_, collateralAmount);

        _collateralToken().safeTransfer(recipient_, collateralAmount);

        _priceSnapshots.push(IMarket.PriceSnapshot({timestamp: block.timestamp, longPrice: price(true)}));
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IMarket
    function collateralToken() external view override returns (IERC20) {
        return IERC20(_collateralToken());
    }

    /// @inheritdoc IMarket
    function oracle() external view returns (IOracle) {
        return IOracle(_oracle());
    }

    /// @inheritdoc IMarket
    function decimals() external view override returns (uint8) {
        return IERC20(_collateralToken()).decimals();
    }

    /// @inheritdoc IMarket
    function price(bool isLong_) public view override returns (uint256) {
        uint256 d = IERC20(_collateralToken()).decimals();
        uint256 totalAmount = _amountInLongPosition + _amountInShortPosition;
        if (totalAmount == 0) {
            return (10 ** d) / 2;
        }

        if (isLong_) {
            return FixedPointMathLib.mulDivUp(_amountInLongPosition, 10 ** d, totalAmount);
        } else {
            return FixedPointMathLib.mulDivUp(_amountInShortPosition, 10 ** d, totalAmount);
        }
    }

    /// @inheritdoc IMarket
    function settlementLongPool() external view override returns (uint256) {
        return _settlementLongPool;
    }

    /// @inheritdoc IMarket
    function settlementShortPool() external view override returns (uint256) {
        return _settlementShortPool;
    }

    /// @inheritdoc IMarket
    function priceSnapshots() external view override returns (IMarket.PriceSnapshot[] memory) {
        return _priceSnapshots;
    }

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the collateral token address from immutable args
    function _collateralToken() private view returns (address) {
        (address collateralToken_,,,,) = _args();
        return collateralToken_;
    }

    /// @notice Returns the oracle address from immutable args
    function _oracle() private view returns (address) {
        (, address oracle_,,,) = _args();
        return oracle_;
    }

    /// @notice Decodes the immutable args appended to this clone
    function _args() private view returns (address, address, string memory, string memory, address) {
        return abi.decode(address(this).argsOnClone(), (address, address, string, string, address));
    }
}
