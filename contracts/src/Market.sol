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
    /*//////////////////////////////////////////////////////////////
                           TYPE DECLARATIONS
    //////////////////////////////////////////////////////////////*/
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

    /*//////////////////////////////////////////////////////////////
                              INITIALIZER
    //////////////////////////////////////////////////////////////*/

    /// @notice Initializes the clone: creates position token clones and sets owner
    /// @param owner_ The owner of this market (the MarketFactory contract)
    function initialize(address owner_) external {
        _initializeOwner(owner_);

        (address collateralToken_,, string memory name_, string memory symbol_, address positionTokenImpl_) = _args();
        uint8 collateralTokenDecimals = IERC20(collateralToken_).decimals();

        address longClone = LibClone.clone(
            positionTokenImpl_,
            abi.encode(string(abi.encodePacked("Long ", name_)), string(abi.encodePacked("L", symbol_)), collateralTokenDecimals)
        );
        PositionToken(longClone).initialize(address(this));
        longPositionToken = PositionToken(longClone);

        address shortClone = LibClone.clone(
            positionTokenImpl_,
            abi.encode(string(abi.encodePacked("Short ", name_)), string(abi.encodePacked("S", symbol_)), collateralTokenDecimals)
        );
        PositionToken(shortClone).initialize(address(this));
        shortPositionToken = PositionToken(shortClone);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Seeds the market with initial collateral split between long and short
    /// @param longCollateral_ Amount of collateral allocated to the long pool
    /// @param shortCollateral_ Amount of collateral allocated to the short pool
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
    }

    /// @notice Settles the market at a specific oracle round
    /// @param settlementRoundId_ The oracle round ID to settle at
    function settle(uint256 settlementRoundId_) external override onlyOwner {
        if (settled) {
            revert AlreadySettled();
        }

        settled = true;
        settlementRoundId = settlementRoundId_;

        emit MarketSettled(settlementRoundId_);
    }

    /// @notice Redeems position tokens for collateral after settlement
    /// @param isLong_ Whether to redeem long (true) or short (false) position tokens
    function redeem(bool isLong_) external override {
        if (!settled) {
            revert NotSettled();
        }

        IERC20 collateralToken_ = IERC20(_collateralToken());
        IOracle oracle_ = IOracle(_oracle());
        uint256 d = collateralToken_.decimals();
        uint8 od = oracle_.decimals();
        IOracle.Round memory round = oracle_.getRound(settlementRoundId);
        uint256 index = round.index;

        if (isLong_) {
            uint256 positionTokenAmount = longPositionToken.balanceOf(msg.sender);
            if (positionTokenAmount == 0) {
                revert InsufficientPositionTokens();
            }

            uint256 redeemPrice = FixedPointMathLib.mulDiv(index, 10 ** d, 10 ** od);
            uint256 collateralAmount = FixedPointMathLib.mulDiv(positionTokenAmount, redeemPrice, 10 ** d);

            longPositionToken.burn(msg.sender, positionTokenAmount);
            _collateralToken().safeTransfer(msg.sender, collateralAmount);

            emit PositionRedeemed(msg.sender, true, positionTokenAmount, collateralAmount);
        } else {
            uint256 positionTokenAmount = shortPositionToken.balanceOf(msg.sender);
            if (positionTokenAmount == 0) {
                revert InsufficientPositionTokens();
            }

            uint256 redeemPrice = 10 ** d - FixedPointMathLib.mulDiv(index, 10 ** d, 10 ** od);
            uint256 collateralAmount = FixedPointMathLib.mulDiv(positionTokenAmount, redeemPrice, 10 ** d);

            shortPositionToken.burn(msg.sender, positionTokenAmount);
            _collateralToken().safeTransfer(msg.sender, collateralAmount);

            emit PositionRedeemed(msg.sender, false, positionTokenAmount, collateralAmount);
        }
    }

    /// @notice Mints position tokens by depositing collateral
    /// @param isLong_ Whether to mint long (true) or short (false) position tokens
    /// @param collateralAmount_ Amount of collateral to deposit
    function mint(bool isLong_, uint256 collateralAmount_) external override {
        if (settled) {
            revert AlreadySettled();
        }
        if (collateralAmount_ == 0) {
            revert InsufficientCollateral();
        }

        uint256 positionTokenPrice = price(isLong_);
        uint256 positionTokenAmount =
            FixedPointMathLib.mulDiv(collateralAmount_, 10 ** IERC20(_collateralToken()).decimals(), positionTokenPrice);

        if (positionTokenAmount == 0) {
            revert InsufficientPositionTokens();
        }

        _collateralToken().safeTransferFrom(msg.sender, address(this), collateralAmount_);

        if (isLong_) {
            _amountInLongPosition += collateralAmount_;
            longPositionToken.mint(msg.sender, positionTokenAmount);
        } else {
            _amountInShortPosition += collateralAmount_;
            shortPositionToken.mint(msg.sender, positionTokenAmount);
        }

        emit PositionMinted(msg.sender, isLong_, collateralAmount_, positionTokenAmount);
    }

    /// @notice Burns position tokens and withdraws collateral
    /// @param isLong_ Whether to burn long (true) or short (false) position tokens
    /// @param positionTokenAmount_ Amount of position tokens to burn
    function burn(bool isLong_, uint256 positionTokenAmount_) external override {
        if (settled) {
            revert AlreadySettled();
        }
        if (positionTokenAmount_ == 0) {
            revert InsufficientPositionTokens();
        }

        uint256 positionTokenPrice = price(isLong_);
        uint256 collateralAmount =
            FixedPointMathLib.mulDiv(positionTokenAmount_, positionTokenPrice, 10 ** IERC20(_collateralToken()).decimals());

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

        emit PositionBurned(msg.sender, isLong_, positionTokenAmount_, collateralAmount);

        _collateralToken().safeTransfer(msg.sender, collateralAmount);
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the collateral token address from immutable args
    function collateralToken() external view override returns (IERC20) {
        return IERC20(_collateralToken());
    }

    /// @notice Returns the oracle address from immutable args
    function oracle() external view returns (IOracle) {
        return IOracle(_oracle());
    }

    /// @notice Returns the collateral token decimals
    function decimals() external view override returns (uint8) {
        return IERC20(_collateralToken()).decimals();
    }

    /// @notice Returns the current price of a position token
    /// @param isLong_ Whether to get the long (true) or short (false) price
    /// @return The price in collateral token units
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
        return abi.decode(LibClone.argsOnClone(address(this)), (address, address, string, string, address));
    }
}
