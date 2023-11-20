// Copyright 2021 - See NOTICE file for copyright holders.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.8.0;
pragma experimental ABIEncoderV2;

import "./perun-eth-contracts/contracts/App.sol";

/**
 * @notice FLApp is a channel app for playing tic tac toe.
 * The data is encoded as follows:
 * - data[0]: The index of the next actor.
 * - data[i], i in [1,10]: The value of field i. 0 means no tick, 1 means tick by player 1, 2 means tick by player 2.
 */
contract FLApp is App {
    uint8 constant actorDataIndex = 0;
    uint8 constant actorDataLength = 1;
    uint8 constant gridDataIndex = actorDataIndex + actorDataLength;
    uint8 constant gridDataLength = 9;
    uint8 constant appDataLength = gridDataIndex + gridDataLength; // Actor index + grid.
    uint8 constant numParts = 2;
    uint8 constant notSet = 0;
    uint8 constant firstPlayer = 1;
    uint8 constant secondPlayer = 2;
    uint8 constant threshold = 60;

    /**
     * @notice ValidTransition checks if there was a valid transition between two states.
     * @param params The parameters of the channel.
     * @param from The current state.
     * @param to The potential next state.
     * @param signerIdx Index of the participant who signed this transition.
     */
    function validTransition(
        Channel.Params calldata params,
        Channel.State calldata from,
        Channel.State calldata to,
        uint256 signerIdx)
    external pure override
    {
        require(params.participants.length == numParts, "number of participants");

        uint8 actorIndex = uint8(from.appData[actorDataIndex]);
        // require(to.appData.length == appDataLength, "data length");
        require(actorIndex == signerIdx, "actor not signer");
        require((actorIndex + 1) % numParts == uint8(to.appData[actorDataIndex]), "next actor");


        require(from.appData.ModelCID == to.appData.ModelCID, string(abi.encodePacked("Cannot override model: expected ", string(from.appData.ModelCID), ", got ", string(to.appData.ModelCID))));
        require(from.appData.NumberOfRounds == to.appData.NumberOfRounds, string(abi.encodePacked("Cannot override number of rounds: expected ", string(from.appData.NumberOfRounds), ", got ", string(to.appData.NumberOfRounds))));
        require(to.appData.Round <= to.appData.NumberOfRounds, string(abi.encodePacked("Round out of bounds: ", string(to.appData.Round))));

        // Test valid action.
        if (from.appData.NextActor == uint8(1)) { // Client conditions
            require(from.appData.Round == to.appData.Round, string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override round: expected ", string(from.appData.Round), ", got ", string(to.appData.Round))));
            require(equalMaps(from.appData.Accuracy, to.appData.Accuracy, from.appData.NumberOfRounds), string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override accuracy")));
            require(equalMaps(from.appData.Loss, to.appData.Loss, from.appData.NumberOfRounds), string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override loss")));
            require(equalExcept(from.appData.Weights, to.appData.Weights, to.appData.Round), string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override weights outside current round")));
            require(to.appData.Weights[to.appData.Round] == 0, "Actor cannot skip weight");
        }

        if (from.appData.NextActor == uint8(0)) { // Server conditions
            require(to.appData.Round == from.appData.Round + 1, string(abi.encodePacked("Actor ", string(from.appData.NextActor), " must increment round: expected ", string(from.appData.Round + 1), ", got ", string(to.appData.Round))));
            require(equalMaps(from.appData.Weights, to.appData.Weights, from.appData.NumberOfRounds), string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override weights")));
            require(equalExcept(from.appData.Accuracy, to.appData.Accuracy, to.appData.Round, from.appData.NumberOfRounds), string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override accuracy outside current round")));
            require(to.appData.Accuracy[to.appData.Round] == 0, "Actor cannot skip accuracy");
            require(equalExcept(from.appData.Loss, to.appData.Loss, to.appData.Round), string(abi.encodePacked("Actor ", string(from.appData.NextActor), " cannot override loss outside current round")));
            require(to.appData.Loss[to.appData.Round] == 0, "Actor cannot skip loss");
        }

        require(from.appData.RoundPhase == to.appData.RoundPhase, string(abi.encodePacked("Cannot repeat round phase: ", string(from.appData.RoundPhase), " -> ", string(to.appData.RoundPhase))));


        // Test final state.
        (bool isFinal, bool hasWinner, uint8 winner) = checkFinal(to.appData);
        require(to.isFinal == isFinal, "final flag");
        Array.requireEqualAddressArray(to.outcome.assets, from.outcome.assets);
        Channel.requireEqualSubAllocArray(to.outcome.locked, from.outcome.locked);
        uint256[][] memory expectedBalances = from.outcome.balances;
        if (hasWinner) {
            uint8 loser = 1 - winner;
            expectedBalances = new uint256[][](expectedBalances.length);
            for (uint i = 0; i < expectedBalances.length; i++) {
                expectedBalances[i] = new uint256[](numParts);
                expectedBalances[i][winner] = from.outcome.balances[i][0] + from.outcome.balances[i][1];
                expectedBalances[i][loser] = 0;
            }
        }
        requireEqualUint256ArrayArray(to.outcome.balances, expectedBalances);
    }


    function equalMaps(mapping(uint8 => uint8) memory d1, mapping(uint8 => uint8) memory d2, uint8 numberOfRounds) internal view returns (bool) {
        for (uint8 k = 0; k < numberOfRounds; k++) {
            if (d2[k] != d1[k]) {
                return false;
            }
        }
        return true;
    }

    function equalExcept(mapping(uint8 => uint8) memory d1, mapping(uint8 => uint8) memory d2, uint8 key, uint8 numberOfRounds) internal view returns (bool) {
        for (uint8 k = 0; k < numberOfRounds; k++) {
            if (k != key && d2[k] != d1[k]) {
                return false;
            }
        }
        return true;
    }


    /// @dev Asserts that a and b are equal.
    function requireEqualAddressArray(
        address[] memory a,
        address[] memory b
    )
    internal
    pure
    {
        require(a.length == b.length, "address[]: unequal length");
        for (uint i = 0; i < a.length; i++) {
            require(a[i] == b[i], "address[]: unequal item");
        }
    }


    function checkFinal(bytes memory d) internal pure returns (bool isFinal, bool hasWinner, uint8 winner) {
        if (d.Round == d.NumberOfRounds - 1 && d.RoundPhase == 2) {
            if (d.Accuracy[d.Round] >= threshold) {
                return (true, true, 1); // FLClient wins
            } else {
                return (true, true, 0); // FLServer wins
            }
        }

        return (false, false, 0);
    }

    function requireEqualUint256ArrayArray(
        uint256[][] memory a,
        uint256[][] memory b
    )
    internal pure
    {
        require(a.length == b.length, "uint256[][]: unequal length");
        for (uint i = 0; i < a.length; i++) {
            Array.requireEqualUint256Array(a[i], b[i]);
        }
    }
}