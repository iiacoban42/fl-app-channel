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
    // uint8 constant actorDataLength = 1;
    // uint8 constant gridDataIndex = actorDataIndex + actorDataLength;
    // uint8 constant gridDataLength = 9;
    // uint8 constant appDataLength = gridDataIndex + gridDataLength; // Actor index + grid.
    uint8 constant modelCIDIndex = 1;
    uint8 constant numberOfRoundsIndex = 2;
    uint8 constant roundIndex = 3;
    uint8 constant roundPhaseIndex = 4;
    uint8 constant weightsIndex = 5;

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
        require(from.appData[numberOfRoundsIndex] == to.appData[numberOfRoundsIndex], string(abi.encodePacked("Cannot change number of rounds: expected", from.appData[numberOfRoundsIndex], " got ", to.appData[numberOfRoundsIndex])));
        // require(from.appData[roundPhaseIndex] != to.appData[roundPhaseIndex],  string(abi.encodePacked("Cannot repeat round phase: ", from.appData[roundPhaseIndex], " -> ", to.appData[roundPhaseIndex])));
        require(from.appData[modelCIDIndex] == to.appData[modelCIDIndex],
                string(abi.encodePacked("Cannot override model: expected ", from.appData[modelCIDIndex], ", got ", to.appData[modelCIDIndex])));
        require(from.appData.length == to.appData.length, "data length");

        uint8 actorIndex = uint8(from.appData[actorDataIndex]);
        require(actorIndex == signerIdx, "actor not signer");
        require((actorIndex + 1) % numParts == uint8(to.appData[actorDataIndex]), "next actor");
        //parse from data state
        uint8 numberOfRounds = uint8(from.appData[numberOfRoundsIndex]);
        uint8 round = uint8(from.appData[roundIndex]);
        uint8 toRound = uint8(to.appData[roundIndex]);

        uint8 accuracyIndex = weightsIndex + numberOfRounds;
        uint8 lossIndex = accuracyIndex + numberOfRounds;



        require(uint8(from.appData[numberOfRoundsIndex]) <= uint8(to.appData[numberOfRoundsIndex]), string(abi.encodePacked("Round out of bounds: ", to.appData[numberOfRoundsIndex])));

        uint8 fromData_NextActor = uint8(to.appData[actorDataIndex]);
        // Test valid action.
        if (fromData_NextActor == uint8(1)) { // Client conditions
                require(uint8(from.appData[roundIndex]) == uint8(to.appData[roundIndex]), string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override round: expected ", from.appData[roundIndex], ", got ", to.appData[roundIndex])));
                // require(from.appData[accuracyIndex : accuracyIndex+numberOfRounds] == to.appData[accuracyIndex: accuracyIndex+numberOfRounds], string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override accuracy")));
                // require(from.appData[lossIndex : lossIndex+numberOfRounds] == to.appData[lossIndex : lossIndex + numberOfRounds], string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override loss")));
                // require(equalExcept(fromData_Weights, toData_Weights, toData_Round, toData_NumberOfRounds), string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override weights outside current round")));
                require(to.appData[weightsIndex+round] != 0, "Actor cannot skip weight");
            }

            if (fromData_NextActor == uint8(0)) { // Server conditions
                require(toRound == round + 1, string(abi.encodePacked("Actor ", fromData_NextActor, " must increment round: expected ", round + 1, ", got ", toRound)));
                // require(from.appData[weightsIndex : weightsIndex+numberOfRounds] == to.appData[weightsIndex: weightsIndex+numberOfRounds], string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override weights")));
                // require(equalExcept(fromData_Accuracy, toData_Accuracy, toData_Round, fromData_NumberOfRounds), string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override accuracy outside current round")));
                require(to.appData[accuracyIndex+round] != 0, "Actor cannot skip accuracy");
                // require(equalExcept(fromData_Loss, toData_Loss, toData_Round, toData_NumberOfRounds), string(abi.encodePacked("Actor ", fromData_NextActor, " cannot override loss outside current round")));
                require(to.appData[lossIndex+round] != 0, "Actor cannot skip weight");

            }

        // Test final state.
        // uint8 finalAccuracy = uint8(to.appData[accuracyIndex+round]);
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
        uint8 numberOfRounds = uint8(d[numberOfRoundsIndex]);
        uint8 round = uint8(d[roundIndex]);
        // uint8 roundPhase = uint8(d[roundPhaseIndex]);
        // if (round == numberOfRounds - 1 && roundPhase == 2) {
        if (round == numberOfRounds - 1) {

            // if (accuracy >= threshold) {
                return (true, true, 1); // FLClient wins
            // } else {
            //     return (true, true, 0); // FLServer wins
            // }
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