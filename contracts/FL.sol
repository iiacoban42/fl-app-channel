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


contract FLApp is App {
    uint8 constant actorDataIndex = 0;
    uint8 constant actorDataLength = 1;
    uint8 constant numParts = 2;

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

        uint8 accuracyIndex = weightIndex + numRounds;
        uint8 lossIndex = accuracyIndex + numRounds;
        uint8 appDataLength = lossIndex + numRounds;

        require(to.appData.length == appDataLength, "data length");

        require(uint8(to.appData[roundIndex]) <= uint8(to.appData[numberOfRoundsIndex]), "round out of bounds");


        // // check server constraints
        if (actorIndex == 0) {
            require(equalBytes(from.appData[weightIndex:weightIndex+numRounds], to.appData[weightIndex:weightIndex+numRounds]), "actor cannot override weights");

            require(equalExcept(from.appData[accuracyIndex:accuracyIndex+numRounds], to.appData[accuracyIndex:accuracyIndex+numRounds], uint8(from.appData[roundIndex])), "actor cannot override accuracy outside current round");

            require(equalExcept(from.appData[lossIndex:lossIndex+numRounds], to.appData[lossIndex:lossIndex+numRounds], uint8(from.appData[roundIndex])), "actor cannot override loss outside current round ");


            if (uint8(from.appData[roundPhaseIndex]) != 0){ // unless we are in init phase
                require(uint8(to.appData[roundIndex]) == uint8(from.appData[roundIndex]) + uint8(1), "actor must increment round");

                require(uint8(to.appData[accuracyIndex+uint8(from.appData[roundIndex])]) != 0, "actor cannot skip accuracy");

                require(uint8(to.appData[lossIndex+uint8(from.appData[roundIndex])]) != 0, "actor cannot skip loss");
            }

        }

        // check client constraints

        if (actorIndex == 1) {
            require(uint8(from.appData[roundIndex]) == uint8(to.appData[roundIndex]), "actor cannot increment round");

            require(equalBytes(from.appData[accuracyIndex:accuracyIndex+numRounds], to.appData[accuracyIndex:accuracyIndex+numRounds]), "actor cannot override accuracy");

            require(equalBytes(from.appData[lossIndex:lossIndex+numRounds], to.appData[lossIndex:lossIndex+numRounds]), "actor cannot override loss");

            require(equalExcept(from.appData[weightIndex:weightIndex+numRounds], to.appData[weightIndex:weightIndex+numRounds], uint8(to.appData[roundIndex])), "actor cannot override weight outside current round");

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


    // check if 2 byte arrays are equal except for one element at index idx
    function equalExcept(bytes memory a, bytes memory b, uint8 idx) internal pure returns (bool) {
        if (a.length != b.length) {
            return false;
        }
        for (uint i = 0; i < a.length; i++) {
            if (i == idx) {
                continue;
            }
            if (a[i] != b[i]) {
                return false;
            }
        }
        return true;
    }

    function equalBytes(bytes memory a, bytes memory b) internal pure returns (bool) {
        if (a.length != b.length) {
            return false;
        }
        for (uint i = 0; i < a.length; i++) {
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