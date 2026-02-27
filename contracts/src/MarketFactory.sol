// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {LibClone} from "solady/utils/LibClone.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {Oracle} from "src/Oracle.sol";
import {Market} from "src/Market.sol";
import {ReceiverTemplate} from "src/keystone/ReceiverTemplate.sol";

/// @title MarketFactory
/// @author @aadeexyz
/// @notice Receives Chainlink CRE reports to create, seed, and settle attention markets
/// @dev Deployed as a clone with immutable args: abi.encode(address collateralToken, uint8 oracleDecimals, string name, string symbol, address marketImpl, address oracleImpl, address positionTokenImpl)
contract MarketFactory is IMarketFactory, ReceiverTemplate {
    using LibClone for address;
    using SafeTransferLib for address;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    address public oracle;

    address[] public markets;

    State public state;

    /*//////////////////////////////////////////////////////////////
                              INITIALIZER
    //////////////////////////////////////////////////////////////*/

    /// @notice Initializes the clone: creates Oracle clone and sets up receiver
    /// @param forwarderAddress_ The Chainlink forwarder address
    /// @param owner_ The owner address
    function initialize(address forwarderAddress_, address owner_) external {
        _initializeReceiverTemplate(forwarderAddress_, owner_);

        (address collateralToken_, uint8 oracleDecimals_, string memory name_,,,,) = _args();

        address oracleClone = _oracleImpl().clone(abi.encode(oracleDecimals_, name_));
        Oracle(oracleClone).initialize(forwarderAddress_, address(this));
        oracle = oracleClone;

        state = State.Creating;

        // Approve collateral token for the first market seeding
        uint256 balance = IERC20(collateralToken_).balanceOf(address(this));
        if (balance > 0) {
            collateralToken_.safeApprove(address(this), balance);
        }
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the collateral token address from immutable args
    function collateralToken() external view override returns (address) {
        (address collateralToken_,,,,,,) = _args();
        return collateralToken_;
    }

    /// @notice Returns the market factory name from immutable args
    function name() external view override returns (string memory) {
        (,, string memory name_,,,,) = _args();
        return name_;
    }

    /// @notice Returns the market factory symbol from immutable args
    function symbol() external view override returns (string memory) {
        (,,, string memory symbol_,,,) = _args();
        return symbol_;
    }

    /// @notice Returns the most recently created market address
    function latestMarket() external view override returns (address) {
        if (markets.length == 0) {
            return address(0);
        }
        return markets[markets.length - 1];
    }

    /// @notice Returns the collateral token decimals
    function decimals() external view override returns (uint8) {
        (address collateralToken_,,,,,,) = _args();
        return IERC20(collateralToken_).decimals();
    }

    /// @notice Returns the total number of markets created
    function totalMarkets() external view override returns (uint256) {
        return markets.length;
    }

    /*//////////////////////////////////////////////////////////////
                            INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Processes a CRE report: updates oracle, settles old market, creates and seeds new market
    /// @param report_ ABI-encoded (uint256 index, uint256 ema)
    function _processReport(bytes calldata report_) internal override {
        (uint256 index, uint256 ema) = abi.decode(report_, (uint256, uint256));

        Oracle(oracle).submitRound(index, ema);

        if (markets.length > 0) {
            address currentMarket = markets[markets.length - 1];
            IMarket(currentMarket).settle(IOracle(oracle).currentRoundId());
            IMarket(currentMarket).redeem(true);
            IMarket(currentMarket).redeem(false);
        }

        _seedMarket();

        if (state == State.Creating) {
            State oldState = state;
            state = State.Live;
            emit StateUpdated(oldState, State.Live);
        }
    }

    /// @notice Deploys a new Market clone
    function _createMarket() internal virtual returns (address) {
        (address collateralToken_,, string memory name_, string memory symbol_,,, address positionTokenImpl_) = _args();

        address marketClone = _marketImpl().clone(
            abi.encode(collateralToken_, oracle, name_, symbol_, positionTokenImpl_)
        );
        Market(marketClone).initialize(address(this));

        markets.push(marketClone);
        emit MarketCreated(marketClone);

        return marketClone;
    }

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Creates a new market and seeds it using the factory's collateral balance
    function _seedMarket() private {
        address market = _createMarket();

        IOracle.Round memory latestRound = IOracle(oracle).getLatestRound();
        uint256 index = latestRound.index;
        uint8 oracleDecimals_ = IOracle(oracle).decimals();

        (address collateralToken_,,,,,,) = _args();
        uint256 collateral = IERC20(collateralToken_).balanceOf(address(this));
        uint256 longCollateral = FixedPointMathLib.mulDiv(collateral, index, 10 ** oracleDecimals_);
        uint256 shortCollateral = collateral - longCollateral;

        collateralToken_.safeApprove(market, collateral);
        IMarket(market).seed(longCollateral, shortCollateral);

        emit MarketSeeded(market, longCollateral, shortCollateral);
    }

    /// @notice Returns the market implementation address from immutable args
    function _marketImpl() private view returns (address) {
        (,,,, address marketImpl_,,) = _args();
        return marketImpl_;
    }

    /// @notice Returns the oracle implementation address from immutable args
    function _oracleImpl() private view returns (address) {
        (,,,,, address oracleImpl_,) = _args();
        return oracleImpl_;
    }

    /// @notice Decodes the immutable args appended to this clone
    function _args()
        private
        view
        returns (address, uint8, string memory, string memory, address, address, address)
    {
        return abi.decode(
            address(this).argsOnClone(), (address, uint8, string, string, address, address, address)
        );
    }
}
