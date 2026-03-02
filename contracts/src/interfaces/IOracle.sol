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

    /// @notice Returns the current round ID
    function currentRoundId() external view returns (uint256);

    /// @notice Returns the round data for a given round ID
    /// @param id The round ID to query
    function getRound(uint256 id) external view returns (Round memory);

    /// @notice Returns the 30-element rolling EMA window
    function rollingEMAWindow() external view returns (uint256[] memory);

    /// @notice Returns the most recent round data
    function getLatestRound() external view returns (Round memory);

    /// @notice Returns the oracle decimals
    function decimals() external view returns (uint8);

    /// @notice Returns the keyword this oracle tracks
    function keyword() external view returns (string memory);

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Submits a new round via the owner (MarketFactory) during settlement
    /// @param index_ The attention index value (scaled to oracle decimals)
    /// @param ema_ The exponential moving average value
    function submitRound(uint256 index_, uint256 ema_) external;
}
