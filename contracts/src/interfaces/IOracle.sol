// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

interface IOracle {
    struct Round {
        uint256 id;
        uint256 timestamp;
        uint256 index;
    }

    event AnswerSubmitted(uint256 indexed roundId, uint256 indexed timestamp, uint256 index, uint256 ema);

    function currentRoundId() external view returns (uint256);

    function getRound(uint256 id) external view returns (Round memory);

    function rollingEMAWindow() external view returns (uint256[] memory);

    function getLatestRound() external view returns (Round memory);

    function decimals() external view returns (uint8);

    function keyword() external view returns (string memory);

    function submitRound(uint256 index_, uint256 ema_) external;
}
