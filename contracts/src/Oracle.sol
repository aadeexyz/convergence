// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IOracle} from "src/interfaces/IOracle.sol";
import {Ownable} from "solady/auth/Ownable.sol";
import {LibClone} from "solady/utils/LibClone.sol";

/// @title Oracle
/// @author @aadeexyz
/// @notice Stores attention index rounds and a rolling EMA window
/// @dev Deployed as a clone with immutable args: abi.encode(uint8 decimals, string keyword)
contract Oracle is IOracle, Ownable {
    /*//////////////////////////////////////////////////////////////
                            STATE VARIABLES
    //////////////////////////////////////////////////////////////*/
    uint256 public currentRoundId;
    uint256[] private _rollingEMAWindow;

    mapping(uint256 id_ => Round) private _rounds;

    /*//////////////////////////////////////////////////////////////
                              INITIALIZER
    //////////////////////////////////////////////////////////////*/

    /// @notice Initializes the clone (can only be called once)
    /// @param owner_ The owner of this oracle (the MarketFactory contract)
    function initialize(address owner_) external {
        _initializeOwner(owner_);
        _rollingEMAWindow = new uint256[](30);
    }

    /*//////////////////////////////////////////////////////////////
                             VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the oracle decimals from immutable args
    function decimals() external view returns (uint8) {
        (uint8 decimals_,) = _args();
        return decimals_;
    }

    /// @notice Returns the keyword from immutable args
    function keyword() external view returns (string memory) {
        (, string memory keyword_) = _args();
        return keyword_;
    }

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

    /*//////////////////////////////////////////////////////////////
                            PRIVATE FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Decodes the immutable args appended to this clone
    function _args() private view returns (uint8, string memory) {
        return abi.decode(LibClone.argsOnClone(address(this)), (uint8, string));
    }
}
