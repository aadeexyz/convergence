// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {Oracle} from "src/Oracle.sol";
import {Market} from "src/Market.sol";
import {ReceiverTemplate} from "src/keystone/ReceiverTemplate.sol";

/// @title MarketFactory
/// @author @aadeexyz
/// @notice Receives Chainlink CRE reports to create, seed, and settle attention markets
contract MarketFactory is IMarketFactory, ReceiverTemplate {
    /*//////////////////////////////////////////////////////////////
                           TYPE DECLARATIONS
    //////////////////////////////////////////////////////////////*/
    using SafeTransferLib for address;

    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    address public immutable collateralToken;
    address public immutable oracle;

    address[] public markets;

    string public name;
    string public symbol;

    State public state;

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/
    constructor(
        address collateralToken_,
        address forwarderAddress_,
        uint8 oracleDecimals_,
        string memory name_,
        string memory symbol_
    ) ReceiverTemplate(forwarderAddress_, msg.sender) {
        collateralToken = collateralToken_;
        name = name_;
        symbol = symbol_;
        state = State.Creating;

        oracle = address(new Oracle(oracleDecimals_, name_, address(this)));
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the most recently created market address
    function latestMarket() external view override returns (address) {
        if (markets.length == 0) {
            return address(0);
        }
        return markets[markets.length - 1];
    }

    /// @notice Returns the collateral token decimals
    function decimals() external view override returns (uint8) {
        return IERC20(collateralToken).decimals();
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

    /// @notice Deploys a new Market contract
    function _createMarket() internal virtual returns (address) {
        address market = address(new Market(collateralToken, oracle, name, symbol));
        markets.push(market);

        emit MarketCreated(market);

        return market;
    }

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Creates a new market and seeds it using the factory's collateral balance
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

        emit MarketSeeded(market, longCollateral, shortCollateral);
    }
}
