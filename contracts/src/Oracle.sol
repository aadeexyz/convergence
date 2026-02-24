// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IOracle} from "src/interfaces/IOracle.sol";
import {Ownable} from "solady/auth/Ownable.sol";

/// @title Oracle
/// @author @aadeexyz
/// @notice Stores attention index rounds and a rolling EMA window
contract Oracle is IOracle, Ownable {
    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    uint256 public currentRoundId;
    uint256[] private _rollingEMAWindow = new uint256[](30);

    uint8 public immutable decimals;
    string public keyword;

    mapping(uint256 id_ => Round) private _rounds;

    /*//////////////////////////////////////////////////////////////
                              CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/
    constructor(uint8 decimals_, string memory keyword_, address owner_) {
        decimals = decimals_;
        keyword = keyword_;
        _initializeOwner(owner_);
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the round data for a given round ID
    /// @param id The round ID to query
    function getRound(uint256 id) external view returns (Round memory) {
        return _rounds[id];
    }

    /// @notice Returns the most recent round data
    function getLatestRound() external view returns (Round memory) {
        return _rounds[currentRoundId];
    }

    /// @notice Returns the 30-element rolling EMA window
    function rollingEMAWindow() external view returns (uint256[] memory) {
        return _rollingEMAWindow;
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Submits a new round with an attention index and EMA value
    /// @param index_ The attention index value (scaled to oracle decimals)
    /// @param ema_ The exponential moving average value
    function submitRound(uint256 index_, uint256 ema_) external onlyOwner {
        _rollingEMAWindow[currentRoundId % 30] = ema_;

        currentRoundId++;
        _rounds[currentRoundId] = Round({id: currentRoundId, timestamp: block.timestamp, index: index_});

        emit AnswerSubmitted(currentRoundId, block.timestamp, index_, ema_);
    }
}
