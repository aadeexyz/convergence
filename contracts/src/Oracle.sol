// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IOracle} from "src/interfaces/IOracle.sol";
import {Ownable} from "solady/auth/Ownable.sol";

contract Oracle is IOracle, Ownable {
    uint256 public currentRoundId;
    uint256[] private _rollingEMAWindow = new uint256[](30);

    uint8 public immutable decimals;
    string public keyword;

    mapping(uint256 id_ => Round) private _rounds;

    constructor(uint8 decimals_, string memory keyword_, address owner_) {
        decimals = decimals_;
        keyword = keyword_;
        _initializeOwner(owner_);
    }

    function getRound(uint256 id) external view returns (Round memory) {
        return _rounds[id];
    }

    function getLatestRound() external view returns (Round memory) {
        return _rounds[currentRoundId];
    }

    function rollingEMAWindow() external view returns (uint256[] memory) {
        return _rollingEMAWindow;
    }

    function submitRound(uint256 index_, uint256 ema_) external onlyOwner {
        _rollingEMAWindow[currentRoundId % 30] = ema_;

        currentRoundId++;
        _rounds[currentRoundId] = Round({id: currentRoundId, timestamp: block.timestamp, index: index_});

        emit AnswerSubmitted(currentRoundId, block.timestamp, index_, ema_);
    }
}
