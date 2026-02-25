// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {Market} from "src/Market.sol";
import {IMarket} from "src/interfaces/IMarket.sol";
import {Oracle} from "src/Oracle.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {PositionToken} from "src/PositionToken.sol";
import {LibClone} from "solady/utils/LibClone.sol";
import {FixedPointMathLib} from "solady/utils/FixedPointMathLib.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {MockERC20} from "test/mocks/MockERC20.sol";

contract MarketTest is Test {
    MockERC20 public token;
    Oracle public oracle;
    Market public market;

    address public owner = address(this);
    address public user = address(0x1);
    address public user2 = address(0x2);

    uint8 constant COLLATERAL_DECIMALS = 6;
    uint8 constant ORACLE_DECIMALS = 8;

    function setUp() public {
        token = new MockERC20("USDC", "USDC", COLLATERAL_DECIMALS);

        Oracle oracleImpl = new Oracle();
        address oracleClone = LibClone.clone(address(oracleImpl), abi.encode(ORACLE_DECIMALS, "baby punch"));
        oracle = Oracle(oracleClone);
        oracle.initialize(owner);

        PositionToken positionTokenImpl = new PositionToken();
        Market marketImpl = new Market();
        address marketClone = LibClone.clone(
            address(marketImpl),
            abi.encode(address(token), address(oracle), "baby punch", "PUNCH", address(positionTokenImpl))
        );
        market = Market(marketClone);
        market.initialize(owner);

        token.mint(owner, 100_000e6);
        token.mint(user, 100_000e6);
        token.mint(user2, 100_000e6);
    }

    // ==================== Constructor ====================

    function test_constructor_setsCollateralToken() public view {
        assertEq(address(market.collateralToken()), address(token));
    }

    function test_constructor_setsOracle() public view {
        assertEq(address(market.oracle()), address(oracle));
    }

    function test_constructor_createsPositionTokens() public view {
        assertTrue(address(market.longPositionToken()) != address(0));
        assertTrue(address(market.shortPositionToken()) != address(0));
    }

    function test_constructor_positionTokenNames() public view {
        assertEq(market.longPositionToken().name(), "Long baby punch");
        assertEq(market.shortPositionToken().name(), "Short baby punch");
    }

    function test_constructor_positionTokenSymbols() public view {
        assertEq(market.longPositionToken().symbol(), "LPUNCH");
        assertEq(market.shortPositionToken().symbol(), "SPUNCH");
    }

    function test_constructor_positionTokenDecimals() public view {
        assertEq(market.longPositionToken().decimals(), COLLATERAL_DECIMALS);
        assertEq(market.shortPositionToken().decimals(), COLLATERAL_DECIMALS);
    }

    function test_constructor_setsOwner() public view {
        assertEq(market.owner(), owner);
    }

    function test_constructor_notSettled() public view {
        assertFalse(market.settled());
        assertEq(market.settlementRoundId(), 0);
    }

    function test_decimals() public view {
        assertEq(market.decimals(), COLLATERAL_DECIMALS);
    }

    // ==================== Price ====================

    function test_price_defaultsTo50Percent() public view {
        uint256 longPrice = market.price(true);
        uint256 shortPrice = market.price(false);
        assertEq(longPrice, 10 ** COLLATERAL_DECIMALS / 2);
        assertEq(shortPrice, 10 ** COLLATERAL_DECIMALS / 2);
    }

    function test_price_reflectsPoolRatio() public {
        token.approve(address(market), 1000e6);
        market.seed(700e6, 300e6);

        uint256 longPrice = market.price(true);
        uint256 shortPrice = market.price(false);

        assertGt(longPrice, shortPrice);
        assertApproxEqAbs(longPrice + shortPrice, 1e6, 2);
    }

    // ==================== Seed ====================

    function test_seed_transfersCollateral() public {
        uint256 balanceBefore = token.balanceOf(owner);
        token.approve(address(market), 1000e6);
        market.seed(600e6, 400e6);

        assertEq(token.balanceOf(owner), balanceBefore - 1000e6);
        assertEq(token.balanceOf(address(market)), 1000e6);
    }

    function test_seed_mintsEqualPositionTokens() public {
        token.approve(address(market), 1000e6);
        market.seed(600e6, 400e6);

        assertEq(market.longPositionToken().balanceOf(owner), 1000e6);
        assertEq(market.shortPositionToken().balanceOf(owner), 1000e6);
    }

    function test_seed_emitsMarketSeeded() public {
        token.approve(address(market), 1000e6);

        vm.expectEmit(true, false, false, true);
        emit IMarket.MarketSeeded(owner, 600e6, 400e6, 1000e6);

        market.seed(600e6, 400e6);
    }

    function test_seed_revertsIfAlreadySeeded() public {
        token.approve(address(market), 2000e6);
        market.seed(500e6, 500e6);

        vm.expectRevert(IMarket.AlreadySeeded.selector);
        market.seed(500e6, 500e6);
    }

    function test_seed_revertsIfZeroCollateral() public {
        vm.expectRevert(IMarket.InsufficientCollateral.selector);
        market.seed(0, 0);
    }

    function test_seed_revertsIfNotOwner() public {
        vm.prank(user);
        vm.expectRevert(Ownable.Unauthorized.selector);
        market.seed(500e6, 500e6);
    }

    // ==================== Mint ====================

    function _seedMarket() internal {
        token.approve(address(market), 1000e6);
        market.seed(500e6, 500e6);
    }

    function test_mint_long() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(true, 100e6);
        vm.stopPrank();

        assertGt(market.longPositionToken().balanceOf(user), 0);
    }

    function test_mint_short() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(false, 100e6);
        vm.stopPrank();

        assertGt(market.shortPositionToken().balanceOf(user), 0);
    }

    function test_mint_transfersCollateral() public {
        _seedMarket();

        uint256 balanceBefore = token.balanceOf(user);

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(true, 100e6);
        vm.stopPrank();

        assertEq(token.balanceOf(user), balanceBefore - 100e6);
    }

    function test_mint_emitsPositionMinted() public {
        _seedMarket();

        uint256 posPrice = market.price(true);
        uint256 expectedTokens = FixedPointMathLib.mulDiv(100e6, 1e6, posPrice);

        vm.startPrank(user);
        token.approve(address(market), 100e6);

        vm.expectEmit(true, true, false, true);
        emit IMarket.PositionMinted(user, true, 100e6, expectedTokens);

        market.mint(true, 100e6);
        vm.stopPrank();
    }

    function test_mint_revertsIfZeroCollateral() public {
        _seedMarket();

        vm.expectRevert(IMarket.InsufficientCollateral.selector);
        market.mint(true, 0);
    }

    function test_mint_revertsIfSettled() public {
        _seedMarket();
        oracle.submitRound(50000000, 40000000);
        market.settle(1);

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        vm.expectRevert(IMarket.AlreadySettled.selector);
        market.mint(true, 100e6);
        vm.stopPrank();
    }

    function test_mint_changesPriceForLong() public {
        _seedMarket();

        uint256 longPriceBefore = market.price(true);

        vm.startPrank(user);
        token.approve(address(market), 500e6);
        market.mint(true, 500e6);
        vm.stopPrank();

        uint256 longPriceAfter = market.price(true);
        assertGt(longPriceAfter, longPriceBefore);
    }

    // ==================== Burn ====================

    function test_burn_long() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(true, 100e6);

        uint256 posTokens = market.longPositionToken().balanceOf(user);
        uint256 collateralBefore = token.balanceOf(user);

        market.burn(true, posTokens);
        vm.stopPrank();

        assertEq(market.longPositionToken().balanceOf(user), 0);
        assertGt(token.balanceOf(user), collateralBefore);
    }

    function test_burn_short() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(false, 100e6);

        uint256 posTokens = market.shortPositionToken().balanceOf(user);
        uint256 collateralBefore = token.balanceOf(user);

        market.burn(false, posTokens);
        vm.stopPrank();

        assertEq(market.shortPositionToken().balanceOf(user), 0);
        assertGt(token.balanceOf(user), collateralBefore);
    }

    function test_burn_emitsPositionBurned() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(true, 100e6);

        uint256 posTokens = market.longPositionToken().balanceOf(user);
        uint256 posPrice = market.price(true);
        uint256 expectedCollateral = FixedPointMathLib.mulDiv(posTokens, posPrice, 1e6);

        vm.expectEmit(true, true, false, true);
        emit IMarket.PositionBurned(user, true, posTokens, expectedCollateral);

        market.burn(true, posTokens);
        vm.stopPrank();
    }

    function test_burn_revertsIfZeroAmount() public {
        _seedMarket();

        vm.expectRevert(IMarket.InsufficientPositionTokens.selector);
        market.burn(true, 0);
    }

    function test_burn_revertsIfSettled() public {
        _seedMarket();
        oracle.submitRound(50000000, 40000000);
        market.settle(1);

        vm.expectRevert(IMarket.AlreadySettled.selector);
        market.burn(true, 100e6);
    }

    // ==================== Settle ====================

    function test_settle_setsSettledFlag() public {
        _seedMarket();
        oracle.submitRound(60000000, 50000000);

        market.settle(1);

        assertTrue(market.settled());
        assertEq(market.settlementRoundId(), 1);
    }

    function test_settle_emitsMarketSettled() public {
        _seedMarket();
        oracle.submitRound(60000000, 50000000);

        vm.expectEmit(true, false, false, true);
        emit IMarket.MarketSettled(1);

        market.settle(1);
    }

    function test_settle_revertsIfAlreadySettled() public {
        _seedMarket();
        oracle.submitRound(60000000, 50000000);
        market.settle(1);

        vm.expectRevert(IMarket.AlreadySettled.selector);
        market.settle(1);
    }

    function test_settle_revertsIfNotOwner() public {
        _seedMarket();
        oracle.submitRound(60000000, 50000000);

        vm.prank(user);
        vm.expectRevert(Ownable.Unauthorized.selector);
        market.settle(1);
    }

    // ==================== Redeem ====================

    function test_redeem_long() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(true, 100e6);
        vm.stopPrank();

        oracle.submitRound(60000000, 50000000);
        market.settle(1);

        uint256 userLongTokens = market.longPositionToken().balanceOf(user);
        uint256 collateralBefore = token.balanceOf(user);

        uint256 redeemPrice = FixedPointMathLib.mulDiv(60000000, 1e6, 10 ** ORACLE_DECIMALS);
        uint256 expectedCollateral = FixedPointMathLib.mulDiv(userLongTokens, redeemPrice, 1e6);

        vm.prank(user);
        market.redeem(true);

        assertEq(market.longPositionToken().balanceOf(user), 0);
        assertEq(token.balanceOf(user), collateralBefore + expectedCollateral);
    }

    function test_redeem_short() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(false, 100e6);
        vm.stopPrank();

        oracle.submitRound(60000000, 50000000);
        market.settle(1);

        uint256 userShortTokens = market.shortPositionToken().balanceOf(user);
        uint256 collateralBefore = token.balanceOf(user);

        uint256 redeemPrice = 1e6 - FixedPointMathLib.mulDiv(60000000, 1e6, 10 ** ORACLE_DECIMALS);
        uint256 expectedCollateral = FixedPointMathLib.mulDiv(userShortTokens, redeemPrice, 1e6);

        vm.prank(user);
        market.redeem(false);

        assertEq(market.shortPositionToken().balanceOf(user), 0);
        assertEq(token.balanceOf(user), collateralBefore + expectedCollateral);
    }

    function test_redeem_emitsPositionRedeemed() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 100e6);
        market.mint(true, 100e6);
        vm.stopPrank();

        oracle.submitRound(70000000, 55000000);
        market.settle(1);

        uint256 userLongTokens = market.longPositionToken().balanceOf(user);
        uint256 redeemPrice = FixedPointMathLib.mulDiv(70000000, 1e6, 10 ** ORACLE_DECIMALS);
        uint256 expectedCollateral = FixedPointMathLib.mulDiv(userLongTokens, redeemPrice, 1e6);

        vm.prank(user);
        vm.expectEmit(true, true, false, true);
        emit IMarket.PositionRedeemed(user, true, userLongTokens, expectedCollateral);
        market.redeem(true);
    }

    function test_redeem_revertsIfNotSettled() public {
        _seedMarket();

        vm.prank(user);
        vm.expectRevert(IMarket.NotSettled.selector);
        market.redeem(true);
    }

    function test_redeem_revertsIfNoPositionTokens() public {
        _seedMarket();
        oracle.submitRound(60000000, 50000000);
        market.settle(1);

        vm.prank(user);
        vm.expectRevert(IMarket.InsufficientPositionTokens.selector);
        market.redeem(true);
    }

    function test_redeem_seederGetsBackFullCollateral() public {
        token.approve(address(market), 1000e6);
        market.seed(700e6, 300e6);

        oracle.submitRound(60000000, 50000000);
        market.settle(1);

        uint256 collateralBefore = token.balanceOf(owner);

        market.redeem(true);
        market.redeem(false);

        uint256 collateralAfter = token.balanceOf(owner);
        assertApproxEqAbs(collateralAfter - collateralBefore, 1000e6, 2);
    }

    function test_redeem_longAndShortPayoutsConsistent() public {
        _seedMarket();

        vm.startPrank(user);
        token.approve(address(market), 200e6);
        market.mint(true, 100e6);
        market.mint(false, 100e6);
        vm.stopPrank();

        oracle.submitRound(50000000, 40000000);
        market.settle(1);

        uint256 userLongTokens = market.longPositionToken().balanceOf(user);
        uint256 userShortTokens = market.shortPositionToken().balanceOf(user);

        uint256 longRedeemPrice = FixedPointMathLib.mulDiv(50000000, 1e6, 1e8);
        uint256 shortRedeemPrice = 1e6 - longRedeemPrice;

        uint256 longPayout = FixedPointMathLib.mulDiv(userLongTokens, longRedeemPrice, 1e6);
        uint256 shortPayout = FixedPointMathLib.mulDiv(userShortTokens, shortRedeemPrice, 1e6);

        vm.startPrank(user);
        market.redeem(true);
        market.redeem(false);
        vm.stopPrank();

        uint256 totalRedeemed = token.balanceOf(user) - (100_000e6 - 200e6);
        assertApproxEqAbs(totalRedeemed, longPayout + shortPayout, 2);
    }
}
