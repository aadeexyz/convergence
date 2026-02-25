// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {Oracle} from "src/Oracle.sol";
import {IOracle} from "src/interfaces/IOracle.sol";
import {LibClone} from "solady/utils/LibClone.sol";
import {Ownable} from "solady/auth/Ownable.sol";

contract OracleTest is Test {
    Oracle public oracle;
    address public owner = address(this);
    address public nonOwner = address(0xBEEF);

    function setUp() public {
        Oracle impl = new Oracle();
        address clone = LibClone.clone(address(impl), abi.encode(uint8(8), "baby punch"));
        oracle = Oracle(clone);
        oracle.initialize(owner);
    }

    function test_constructor_setsDecimals() public view {
        assertEq(oracle.decimals(), 8);
    }

    function test_constructor_setsKeyword() public view {
        assertEq(oracle.keyword(), "baby punch");
    }

    function test_constructor_setsOwner() public view {
        assertEq(oracle.owner(), owner);
    }

    function test_constructor_initialRoundIdIsZero() public view {
        assertEq(oracle.currentRoundId(), 0);
    }

    function test_submitRound_incrementsRoundId() public {
        oracle.submitRound(50000000, 45000000);
        assertEq(oracle.currentRoundId(), 1);

        oracle.submitRound(60000000, 48000000);
        assertEq(oracle.currentRoundId(), 2);
    }

    function test_submitRound_storesRoundData() public {
        oracle.submitRound(70000000, 50000000);

        IOracle.Round memory round = oracle.getRound(1);
        assertEq(round.id, 1);
        assertEq(round.index, 70000000);
        assertEq(round.timestamp, block.timestamp);
    }

    function test_submitRound_updatesEMAWindow() public {
        oracle.submitRound(50000000, 42000000);

        uint256[] memory window = oracle.rollingEMAWindow();
        assertEq(window[0], 42000000);
    }

    function test_submitRound_emaWindowWrapsAt30() public {
        for (uint256 i = 0; i < 31; i++) {
            oracle.submitRound(50000000, uint256(i + 1) * 1000000);
        }

        uint256[] memory window = oracle.rollingEMAWindow();
        assertEq(window[0], 31000000);
        assertEq(window[1], 2000000);
    }

    function test_submitRound_emitsAnswerSubmitted() public {
        vm.expectEmit(true, true, false, true);
        emit IOracle.AnswerSubmitted(1, block.timestamp, 70000000, 50000000);

        oracle.submitRound(70000000, 50000000);
    }

    function test_getLatestRound_returnsLatest() public {
        oracle.submitRound(50000000, 40000000);
        oracle.submitRound(70000000, 55000000);

        IOracle.Round memory latest = oracle.getLatestRound();
        assertEq(latest.id, 2);
        assertEq(latest.index, 70000000);
    }

    function test_getLatestRound_returnsEmptyWhenNoRounds() public view {
        IOracle.Round memory latest = oracle.getLatestRound();
        assertEq(latest.id, 0);
        assertEq(latest.index, 0);
        assertEq(latest.timestamp, 0);
    }

    function test_getRound_returnsEmptyForNonExistent() public view {
        IOracle.Round memory round = oracle.getRound(999);
        assertEq(round.id, 0);
        assertEq(round.index, 0);
    }

    function test_rollingEMAWindow_returns30Elements() public view {
        uint256[] memory window = oracle.rollingEMAWindow();
        assertEq(window.length, 30);
    }

    function test_submitRound_revertsForNonOwner() public {
        vm.prank(nonOwner);
        vm.expectRevert(Ownable.Unauthorized.selector);
        oracle.submitRound(50000000, 40000000);
    }

    function test_multipleRounds_correctSequence() public {
        oracle.submitRound(30000000, 25000000);
        oracle.submitRound(50000000, 40000000);
        oracle.submitRound(80000000, 65000000);

        IOracle.Round memory r1 = oracle.getRound(1);
        IOracle.Round memory r2 = oracle.getRound(2);
        IOracle.Round memory r3 = oracle.getRound(3);

        assertEq(r1.index, 30000000);
        assertEq(r2.index, 50000000);
        assertEq(r3.index, 80000000);
        assertEq(oracle.currentRoundId(), 3);
    }
}
