// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {PositionToken} from "src/PositionToken.sol";
import {Ownable} from "solady/auth/Ownable.sol";

contract PositionTokenTest is Test {
    PositionToken public token;
    address public owner = address(this);
    address public user = address(0x1);
    address public nonOwner = address(0xBEEF);

    function setUp() public {
        token = new PositionToken("Long Bitcoin", "LBTC", 6);
    }

    function test_constructor_setsName() public view {
        assertEq(token.name(), "Long Bitcoin");
    }

    function test_constructor_setsSymbol() public view {
        assertEq(token.symbol(), "LBTC");
    }

    function test_constructor_setsDecimals() public view {
        assertEq(token.decimals(), 6);
    }

    function test_constructor_setsOwner() public view {
        assertEq(token.owner(), owner);
    }

    function test_mint_increasesBalance() public {
        token.mint(user, 1000e6);
        assertEq(token.balanceOf(user), 1000e6);
    }

    function test_mint_increasesTotalSupply() public {
        token.mint(user, 1000e6);
        assertEq(token.totalSupply(), 1000e6);
    }

    function test_mint_multipleRecipients() public {
        token.mint(user, 500e6);
        token.mint(nonOwner, 300e6);
        assertEq(token.balanceOf(user), 500e6);
        assertEq(token.balanceOf(nonOwner), 300e6);
        assertEq(token.totalSupply(), 800e6);
    }

    function test_burn_decreasesBalance() public {
        token.mint(user, 1000e6);
        token.burn(user, 400e6);
        assertEq(token.balanceOf(user), 600e6);
    }

    function test_burn_decreasesTotalSupply() public {
        token.mint(user, 1000e6);
        token.burn(user, 400e6);
        assertEq(token.totalSupply(), 600e6);
    }

    function test_mint_revertsForNonOwner() public {
        vm.prank(nonOwner);
        vm.expectRevert(Ownable.Unauthorized.selector);
        token.mint(user, 1000e6);
    }

    function test_burn_revertsForNonOwner() public {
        token.mint(user, 1000e6);

        vm.prank(nonOwner);
        vm.expectRevert(Ownable.Unauthorized.selector);
        token.burn(user, 500e6);
    }
}
