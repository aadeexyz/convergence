// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {MarketFactory} from "src/MarketFactory.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {Oracle} from "src/Oracle.sol";
import {Market} from "src/Market.sol";
import {PositionToken} from "src/PositionToken.sol";
import {LibClone} from "solady/utils/LibClone.sol";
import {MockERC20} from "test/mocks/MockERC20.sol";

contract MarketFactoryTest is Test {
    MockERC20 public token;
    MarketFactory public factory;

    address public forwarder = address(0xF0F0);
    address public deployer = address(this);

    uint8 constant ORACLE_DECIMALS = 8;
    uint256 constant SEED_AMOUNT = 1000e6;

    function setUp() public {
        token = new MockERC20("USDC", "USDC", 6);

        MarketFactory mfImpl = new MarketFactory();
        Market marketImpl = new Market();
        Oracle oracleImpl = new Oracle();
        PositionToken ptImpl = new PositionToken();

        address clone = LibClone.clone(
            address(mfImpl),
            abi.encode(
                address(token),
                ORACLE_DECIMALS,
                "baby punch",
                "PUNCH",
                address(marketImpl),
                address(oracleImpl),
                address(ptImpl)
            )
        );
        factory = MarketFactory(clone);
        factory.initialize(forwarder, deployer);
    }

    // ==================== Constructor ====================

    function test_constructor_setsCollateralToken() public view {
        assertEq(factory.collateralToken(), address(token));
    }

    function test_constructor_setsName() public view {
        assertEq(factory.name(), "baby punch");
    }

    function test_constructor_setsSymbol() public view {
        assertEq(factory.symbol(), "PUNCH");
    }

    function test_constructor_startsInCreatingState() public view {
        assertEq(uint256(factory.state()), uint256(IMarketFactory.State.Creating));
    }

    function test_constructor_createsOracle() public view {
        assertTrue(factory.oracle() != address(0));
    }

    function test_constructor_oracleOwnedByFactory() public view {
        Oracle oracle = Oracle(factory.oracle());
        assertEq(oracle.owner(), address(factory));
    }

    function test_constructor_oracleHasCorrectDecimals() public view {
        Oracle oracle = Oracle(factory.oracle());
        assertEq(oracle.decimals(), ORACLE_DECIMALS);
    }

    function test_constructor_oracleHasCorrectKeyword() public view {
        Oracle oracle = Oracle(factory.oracle());
        assertEq(oracle.keyword(), "baby punch");
    }

    function test_constructor_setsOwner() public view {
        assertEq(factory.owner(), deployer);
    }

    // ==================== View Functions ====================

    function test_latestMarket_returnsZeroWhenEmpty() public view {
        assertEq(factory.latestMarket(), address(0));
    }

    function test_totalMarkets_returnsZero() public view {
        assertEq(factory.totalMarkets(), 0);
    }

    function test_decimals_returnsCollateralDecimals() public view {
        assertEq(factory.decimals(), 6);
    }

    // ==================== onReport - First Round ====================

    function _fundAndReport(uint256 index, uint256 ema) internal {
        token.mint(address(factory), SEED_AMOUNT);
        vm.prank(forwarder);
        factory.onReport("", abi.encode(index, ema));
    }

    function test_onReport_firstRound_createsMarket() public {
        _fundAndReport(70000000, 50000000);
        assertEq(factory.totalMarkets(), 1);
        assertTrue(factory.latestMarket() != address(0));
    }

    function test_onReport_firstRound_submitsOracleRound() public {
        _fundAndReport(70000000, 50000000);

        Oracle oracle = Oracle(factory.oracle());
        assertEq(oracle.currentRoundId(), 1);

        IOracle.Round memory round = oracle.getRound(1);
        assertEq(round.index, 70000000);
    }

    function test_onReport_firstRound_seedsMarket() public {
        _fundAndReport(70000000, 50000000);

        address marketAddr = factory.latestMarket();
        Market market = Market(marketAddr);

        assertEq(market.longPositionToken().balanceOf(address(factory)), SEED_AMOUNT);
        assertEq(market.shortPositionToken().balanceOf(address(factory)), SEED_AMOUNT);
        assertEq(token.balanceOf(marketAddr), SEED_AMOUNT);

        uint256 longPrice = market.price(true);
        uint256 shortPrice = market.price(false);
        assertApproxEqAbs(longPrice, 700000, 1);
        assertApproxEqAbs(shortPrice, 300000, 1);
    }

    function test_onReport_firstRound_transitionsToLive() public {
        _fundAndReport(70000000, 50000000);
        assertEq(uint256(factory.state()), uint256(IMarketFactory.State.Live));
    }

    function test_onReport_firstRound_emitsStateUpdated() public {
        token.mint(address(factory), SEED_AMOUNT);

        vm.expectEmit(false, false, false, true);
        emit IMarketFactory.StateUpdated(IMarketFactory.State.Creating, IMarketFactory.State.Live);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(70000000), uint256(50000000)));
    }

    function test_onReport_firstRound_emitsMarketCreated() public {
        token.mint(address(factory), SEED_AMOUNT);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(70000000), uint256(50000000)));

        assertEq(factory.totalMarkets(), 1);
    }

    function test_onReport_factoryCollateralIsZeroAfterSeed() public {
        _fundAndReport(70000000, 50000000);
        assertEq(token.balanceOf(address(factory)), 0);
    }

    // ==================== onReport - Second Round (Settle + Seed) ====================

    function test_onReport_secondRound_settlesFirstMarket() public {
        _fundAndReport(70000000, 50000000);

        address firstMarket = factory.latestMarket();

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(60000000), uint256(48000000)));

        assertTrue(Market(firstMarket).settled());
        assertEq(Market(firstMarket).settlementRoundId(), 2);
    }

    function test_onReport_secondRound_createsNewMarket() public {
        _fundAndReport(70000000, 50000000);

        address firstMarket = factory.latestMarket();

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(60000000), uint256(48000000)));

        assertEq(factory.totalMarkets(), 2);
        assertTrue(factory.latestMarket() != firstMarket);
    }

    function test_onReport_secondRound_factoryRedeemsFromFirst() public {
        _fundAndReport(70000000, 50000000);

        address firstMarket = factory.latestMarket();
        Market m = Market(firstMarket);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(60000000), uint256(48000000)));

        assertEq(m.longPositionToken().balanceOf(address(factory)), 0);
        assertEq(m.shortPositionToken().balanceOf(address(factory)), 0);
    }

    function test_onReport_secondRound_seedsNewMarket() public {
        _fundAndReport(70000000, 50000000);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(60000000), uint256(48000000)));

        address newMarket = factory.latestMarket();
        assertGt(token.balanceOf(newMarket), 0);
    }

    function test_onReport_secondRound_stateRemainsLive() public {
        _fundAndReport(70000000, 50000000);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(60000000), uint256(48000000)));

        assertEq(uint256(factory.state()), uint256(IMarketFactory.State.Live));
    }

    function test_onReport_threeRounds_fullCycle() public {
        _fundAndReport(70000000, 50000000);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(60000000), uint256(48000000)));

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(40000000), uint256(42000000)));

        assertEq(factory.totalMarkets(), 3);

        Oracle oracle = Oracle(factory.oracle());
        assertEq(oracle.currentRoundId(), 3);

        assertTrue(Market(factory.markets(0)).settled());
        assertTrue(Market(factory.markets(1)).settled());
        assertFalse(Market(factory.markets(2)).settled());
    }

    function test_onReport_collateralPreservedAcrossRounds() public {
        _fundAndReport(70000000, 50000000);

        address market1 = factory.latestMarket();
        uint256 collateralInMarket1 = token.balanceOf(market1);

        vm.prank(forwarder);
        factory.onReport("", abi.encode(uint256(30000000), uint256(25000000)));

        address market2 = factory.latestMarket();
        uint256 collateralInMarket2 = token.balanceOf(market2);

        assertApproxEqAbs(collateralInMarket2, collateralInMarket1, 2);
    }

    // ==================== Reverts ====================

    function test_onReport_revertsIfNotForwarder() public {
        token.mint(address(factory), SEED_AMOUNT);

        vm.prank(address(0xDEAD));
        vm.expectRevert();
        factory.onReport("", abi.encode(uint256(70000000), uint256(50000000)));
    }
}
