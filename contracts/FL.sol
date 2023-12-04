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

pragma solidity ^0.7.0;
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

    // uint8 constant gridDataIndex = actorDataIndex + actorDataLength;
    uint8 constant rounds = 1;
    // uint8 constant appDataLength = gridDataIndex + gridDataLength; // Actor index + grid.
    uint8 constant numParts = 2;
    // uint8 constant notSet = 0;
    // uint8 constant firstPlayer = 1;
    // uint8 constant secondPlayer = 2;
    uint8 constant modelIndex = 1;
    uint8 constant numberOfRoundsIndex = modelIndex + 1;
    uint8 constant roundIndex = numberOfRoundsIndex + 1;
    uint8 constant roundPhaseIndex = roundIndex + 1;
    uint8 constant weightIndex = roundPhaseIndex + 1;
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
        require(actorIndex == signerIdx, "actor not signer");
        require((actorIndex + 1) % numParts == uint8(to.appData[actorDataIndex]), "next actor");

        if (uint8(from.appData[roundPhaseIndex]) != 0){
            require(from.appData[modelIndex] == to.appData[modelIndex], "model changed");
            require(from.appData[numberOfRoundsIndex] == to.appData[numberOfRoundsIndex], "round changed");
        }

        uint8 numRounds = uint8(to.appData[numberOfRoundsIndex]);

        uint8  accuracyIndex = weightIndex + numRounds;
        uint8  lossIndex = accuracyIndex + numRounds;
        uint8  appDataLength = lossIndex + numRounds;
        require(to.appData.length == appDataLength, "data length");


        require(uint8(from.appData[numberOfRoundsIndex]) <= uint8(to.appData[numberOfRoundsIndex]), string(abi.encodePacked("roundIndex out of bounds: ", to.appData[numberOfRoundsIndex])));


        // check server constraints
        if (actorIndex == 0) {
            require(!(from.appData[roundPhaseIndex] != 0 && uint8(to.appData[roundIndex]) != uint8(from.appData[roundIndex]) + uint8(1)), string(abi.encodePacked("actor must increment round: expected ", uint8(from.appData[roundIndex]) + 1, ", got ", to.appData[roundIndex])));

            require(from.appData[weightIndex+numRounds] == to.appData[weightIndex+numRounds], string(abi.encodePacked("actor cannot override weights: expected", from.appData[weightIndex+numRounds], ", got ", to.appData[weightIndex+numRounds])));

            // require(!(!equalExcept(from.appData[accuracyIndex:accuracyIndex+numRounds], to.appData[accuracyIndex:accuracyIndex+numRounds], int(from.appData[roundIndex]))), string(abi.encodePacked("actor cannot override accuracy outside current round: expected ", from.appData[accuracyIndex], ", got ", to.appData.accuracyIndex)));
            if (uint8(from.appData[roundPhaseIndex]) != 0){
                require(uint8(to.appData[accuracyIndex+uint8(from.appData[roundIndex])]) != 0, "actor cannot skip accuracy");
                require(uint8(to.appData[lossIndex+uint8(from.appData[roundIndex])]) != 0, "actor cannot skip loss");
            }
            // require(!(!equalExcept(from.appData[lossIndex:lossIndex+numRounds], to.appData[lossIndex:lossIndex+numRounds], int(from.appData[roundIndex]))), string(abi.encodePacked("actor cannot override loss outside current round: expected ", from.appData[lossIndex], ", got ", to.appData[lossIndex])));

        }

        // check client constraints

        if (actorIndex == 1) {
            require(from.appData[roundIndex] == to.appData[roundIndex], "actor cannot increment round");

            require(from.appData[accuracyIndex+numRounds] == to.appData[accuracyIndex+numRounds], string(abi.encodePacked("actor cannot override accuracy: expected", from.appData[accuracyIndex+numRounds], ", got ", to.appData[accuracyIndex+numRounds])));

            require(from.appData[lossIndex+numRounds] == to.appData[lossIndex+numRounds], string(abi.encodePacked("actor cannot override loss: expected", from.appData[lossIndex+numRounds], ", got ", to.appData[lossIndex+numRounds])));

            // require(!(!equalExcept(from.appData[weightIndex:weightIndex+numRounds], to.appData[weightIndex:weightIndex+numRounds], int(to.appData[roundIndex]))), string(abi.encodePacked("actor cannot override weight outside current round: expected ", from.appData[accuracyIndex], ", got ", to.appData[accuracyIndex])));

            require(uint8(to.appData[weightIndex+uint8(from.appData[roundIndex])]) != 0, "actor cannot skip weight");


        }

        // Test final state.
        (bool isFinal, bool hasWinner, uint8 winner) = checkFinal(to.appData, accuracyIndex);
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

    function checkFinal(bytes memory d, uint8 accuracyIndex) internal pure returns (bool isFinal, bool hasWinner, uint8 winner) {
        if (d[numberOfRoundsIndex] == d[roundIndex] && uint8(d[numberOfRoundsIndex]) == 3) {
            if (uint8(d[accuracyIndex]) >= threshold) {
                return (true, true, 1);
            }
                return (true, true, 0);
        }
        return (false, false, 0);

    }


    // check if 2 arrays are equal except for one element at index idx
    function equalExcept(uint256[] memory a, uint256[] memory b, int idx) internal pure returns (bool) {
        if (a.length != b.length) {
            return false;
        }
        for (uint i = 0; i < a.length; i++) {
            if (int(i) == idx) {
                continue;
            }
            if (a[i] != b[i]) {
                return false;
            }
        }
        return true;
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