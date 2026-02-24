// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/// @title IOracle
/// @author @aadeexyz
/// @notice Interface for the attention index oracle
interface IOracle {
    /*//////////////////////////////////////////////////////////////
                                 STRUCTS
    //////////////////////////////////////////////////////////////*/
    struct Round {
        uint256 id;
        uint256 timestamp;
        uint256 index;
    }

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/
    event AnswerSubmitted(uint256 indexed roundId, uint256 indexed timestamp, uint256 index, uint256 ema);

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    function currentRoundId() external view returns (uint256);

    function getRound(uint256 id) external view returns (Round memory);

    function rollingEMAWindow() external view returns (uint256[] memory);

    function getLatestRound() external view returns (Round memory);

    function decimals() external view returns (uint8);

    function keyword() external view returns (string memory);

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/
    function submitRound(uint256 index_, uint256 ema_) external;
}
