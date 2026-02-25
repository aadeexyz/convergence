// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {FactoryOfMarketFactories} from "src/FactoryOfMarketFactories.sol";
import {IFactoryOfMarketFactories} from "src/interfaces/IFactoryOfMarketFactories.sol";
import {MarketFactory} from "src/MarketFactory.sol";
import {IMarketFactory} from "src/interfaces/IMarketFactory.sol";
import {Oracle} from "src/Oracle.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {MockERC20} from "test/mocks/MockERC20.sol";

contract FactoryOfMarketFactoriesTest is Test {
    MockERC20 public token;
    FactoryOfMarketFactories public fof;

    address public owner = address(this);
    address public user = address(0x1);
    address public forwarder = address(0xF0F0);

    uint256 constant LIQUIDITY_FEE = 1000e6;
    uint256 constant PROTOCOL_FEE = 100e6;
    uint8 constant ORACLE_DECIMALS = 8;

    function setUp() public {
        token = new MockERC20("USDC", "USDC", 6);
        fof = new FactoryOfMarketFactories(
            address(token), LIQUIDITY_FEE, PROTOCOL_FEE, forwarder, ORACLE_DECIMALS, owner
        );

        token.mint(user, 100_000e6);
    }

    // ==================== Constructor ====================

    function test_constructor_setsCollateralToken() public view {
        assertEq(fof.collateralToken(), address(token));
    }

    function test_constructor_setsOwner() public view {
        assertEq(fof.owner(), owner);
    }

    function test_constructor_setsCreationFee() public view {
        assertEq(fof.creationFee(), LIQUIDITY_FEE + PROTOCOL_FEE);
    }

    function test_constructor_zeroMarketFactories() public view {
        assertEq(fof.totalMarketFactories(), 0);
    }

    // ==================== setLiquidityFee ====================

    function test_setLiquidityFee_updatesFee() public {
        fof.setLiquidityFee(2000e6);
        assertEq(fof.creationFee(), 2000e6 + PROTOCOL_FEE);
    }

    function test_setLiquidityFee_emitsEvent() public {
        vm.expectEmit(false, false, false, true);
        emit IFactoryOfMarketFactories.LiquidityFeeUpdated(LIQUIDITY_FEE, 2000e6);
        fof.setLiquidityFee(2000e6);
    }

    function test_setLiquidityFee_revertsForNonOwner() public {
        vm.prank(user);
        vm.expectRevert(Ownable.Unauthorized.selector);
        fof.setLiquidityFee(2000e6);
    }

    // ==================== setProtocolFee ====================

    function test_setProtocolFee_updatesFee() public {
        fof.setProtocolFee(200e6);
        assertEq(fof.creationFee(), LIQUIDITY_FEE + 200e6);
    }

    function test_setProtocolFee_emitsEvent() public {
        vm.expectEmit(false, false, false, true);
        emit IFactoryOfMarketFactories.ProtocolFeeUpdated(PROTOCOL_FEE, 200e6);
        fof.setProtocolFee(200e6);
    }

    function test_setProtocolFee_revertsForNonOwner() public {
        vm.prank(user);
        vm.expectRevert(Ownable.Unauthorized.selector);
        fof.setProtocolFee(200e6);
    }

    // ==================== createMarketFactory ====================

    function _createFactory(string memory name, string memory symbol) internal returns (address) {
        vm.startPrank(user);
        token.approve(address(fof), LIQUIDITY_FEE + PROTOCOL_FEE);
        address factoryAddr = fof.createMarketFactory(name, symbol);
        vm.stopPrank();
        return factoryAddr;
    }

    function test_createMarketFactory_createsFactory() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        assertTrue(factoryAddr != address(0));
    }

    function test_createMarketFactory_incrementsCount() public {
        _createFactory("Baby Punch", "PUNCH");
        assertEq(fof.totalMarketFactories(), 1);
    }

    function test_createMarketFactory_lowercasesName() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        MarketFactory mf = MarketFactory(factoryAddr);
        assertEq(mf.name(), "baby punch");
    }

    function test_createMarketFactory_uppercasesSymbol() public {
        address factoryAddr = _createFactory("Baby Punch", "punch");
        MarketFactory mf = MarketFactory(factoryAddr);
        assertEq(mf.symbol(), "PUNCH");
    }

    function test_createMarketFactory_transfersLiquidityFee() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        assertEq(token.balanceOf(factoryAddr), LIQUIDITY_FEE);
    }

    function test_createMarketFactory_transfersProtocolFee() public {
        uint256 ownerBalanceBefore = token.balanceOf(owner);
        _createFactory("Baby Punch", "PUNCH");
        assertEq(token.balanceOf(owner), ownerBalanceBefore + PROTOCOL_FEE);
    }

    function test_createMarketFactory_emitsEvent() public {
        vm.startPrank(user);
        token.approve(address(fof), LIQUIDITY_FEE + PROTOCOL_FEE);

        vm.expectEmit(false, false, false, true);
        emit IFactoryOfMarketFactories.MarketFactoryCreated(address(0), "baby punch", "PUNCH");

        fof.createMarketFactory("Baby Punch", "PUNCH");
        vm.stopPrank();
    }

    function test_createMarketFactory_factoryStateIsCreating() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        MarketFactory mf = MarketFactory(factoryAddr);
        assertEq(uint256(mf.state()), uint256(IMarketFactory.State.Creating));
    }

    function test_createMarketFactory_factoryHasOracle() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        MarketFactory mf = MarketFactory(factoryAddr);
        assertTrue(mf.oracle() != address(0));
    }

    function test_createMarketFactory_oracleHasCorrectKeyword() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        MarketFactory mf = MarketFactory(factoryAddr);
        Oracle oracle = Oracle(mf.oracle());
        assertEq(oracle.keyword(), "baby punch");
    }

    function test_createMarketFactory_factoryOwnerIsFOF() public {
        address factoryAddr = _createFactory("Baby Punch", "PUNCH");
        MarketFactory mf = MarketFactory(factoryAddr);
        assertEq(mf.owner(), address(fof));
    }

    function test_createMarketFactory_marketFactoriesArray() public {
        address f1 = _createFactory("Baby Punch", "PUNCH");

        token.mint(user, 100_000e6);
        address f2 = _createFactory("Sock Rocket", "SOCK");

        address[] memory factories = fof.marketFactories();
        assertEq(factories.length, 2);
        assertEq(factories[0], f1);
        assertEq(factories[1], f2);
    }

    function test_createMarketFactory_marketFactoryByIndex() public {
        address f1 = _createFactory("Baby Punch", "PUNCH");
        assertEq(fof.marketFactory(0), f1);
    }

    // ==================== createMarketFactory Reverts ====================

    function test_createMarketFactory_revertsIfEmptyName() public {
        vm.prank(user);
        vm.expectRevert(IFactoryOfMarketFactories.InvalidName.selector);
        fof.createMarketFactory("", "PUNCH");
    }

    function test_createMarketFactory_revertsIfEmptySymbol() public {
        vm.prank(user);
        vm.expectRevert(IFactoryOfMarketFactories.InvalidSymbol.selector);
        fof.createMarketFactory("Baby Punch", "");
    }

    function test_createMarketFactory_revertsIfAlreadyExists() public {
        _createFactory("Baby Punch", "PUNCH");

        vm.startPrank(user);
        token.approve(address(fof), LIQUIDITY_FEE + PROTOCOL_FEE);
        vm.expectRevert();
        fof.createMarketFactory("Baby Punch", "PUNCH");
        vm.stopPrank();
    }

    function test_createMarketFactory_revertsIfAlreadyExistsCaseInsensitive() public {
        _createFactory("Baby Punch", "PUNCH");

        vm.startPrank(user);
        token.approve(address(fof), LIQUIDITY_FEE + PROTOCOL_FEE);
        vm.expectRevert();
        fof.createMarketFactory("BABY PUNCH", "PUNCH");
        vm.stopPrank();
    }
}
