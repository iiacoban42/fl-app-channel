// Copyright 2021 PolyCrypt GmbH, Germany
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

package app

import (
	"fmt"
	"io"
	"math/big"

	"perun.network/go-perun/channel"
)

const numParts = 2

type FieldValue uint8

const (
	notSet FieldValue = iota
	player1
	player2
	maxFieldValue = player2
)

func (v FieldValue) String() string {
	switch v {
	case notSet:
		return " "
	case player1:
		return "FLServer"
	case player2:
		return "FLClient"
	default:
		panic(fmt.Sprintf("unsupported value: %d", v))
	}
}

func makeFieldValueFromPlayerIdx(idx channel.Index) FieldValue {
	switch idx {
	case 0:
		return player1
	case 1:
		return player2
	default:
		panic("invalid")
	}
}

func (v FieldValue) PlayerIndex() channel.Index {
	switch v {
	case player1:
		return 0
	case player2:
		return 1
	default:
		panic("invalid")
	}
}

// func (d FLAppData) CheckFinal() (isFinal bool, winner *channel.Index) {
// 	// 0 1 2
// 	// 3 4 5
// 	// 6 7 8

// 	// Check winner.
// 	v := [][]int{
// 		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
// 		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
// 		{0, 4, 8}, {2, 4, 6}, // diagonals
// 	}

// 	for _, _v := range v {
// 		ok, idx := d.samePlayer(_v...)
// 		if ok {
// 			return true, &idx
// 		}
// 	}

// 	// Check all set.
// 	for _, v := range d.Grid {
// 		if v != notSet {
// 			return false, nil
// 		}
// 	}
// 	return true, nil
// }

func (d FLAppData) CheckFinal() (isFinal bool, winner *channel.Index) {
	if d.Round == d.NumberOfRounds - 1 && d.RoundPhase == 2 {
		if d.Accuracy[d.Round] >= 60 { // FLClient wins
				index := makeFieldValueFromPlayerIdx(1)
				playerIndex := index.PlayerIndex()
				return true, &playerIndex
			}
			// FLServer wins
			index := makeFieldValueFromPlayerIdx(0)
			playerIndex := index.PlayerIndex()
			return true, &playerIndex
		}

	return false, nil
}


// func (d FLAppData) samePlayer(gridIndices ...int) (ok bool, player channel.Index) {
// 	if len(gridIndices) < 2 {
// 		panic("expecting at least two inputs")
// 	}

// 	first := d.Grid[gridIndices[0]]
// 	if first == notSet {
// 		return false, 0
// 	}
// 	for _, i := range gridIndices {
// 		if d.Grid[i] != first {
// 			return false, 0
// 		}
// 	}
// 	return true, 0
// }

func uint8safe(a uint16) uint8 {
	b := uint8(a)
	if uint16(b) != a {
		panic("unsafe")
	}
	return b
}

func readUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

func writeUInt8(w io.Writer, v uint8) error {
	_, err := w.Write([]byte{v})
	return err
}

func readUInt8Array(r io.Reader, n uint8) ([]uint8, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

// read a string from reader
func readString(r io.Reader, n int) (string, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return string(buf), err
}

// read a int to string map from reader
func readUInt8StringMap(r io.Reader, n uint8, strlen uint8) (map[uint8]string, error) {
	m := make(map[uint8]string)
	for i := uint8(0); i < n; i++ {
		v, err := readString(r, int(strlen))
		if err != nil {
			return nil, err
		}
		m[uint8(i)] = v
	}
	return m, nil
}

// read a int to int map from reader
func readUInt8UInt8Map(r io.Reader, n int) (map[uint8]uint8, error) {
	m := make(map[uint8]uint8)
	for i := 0; i < n; i++ {
		v, err := readUInt8(r)
		if err != nil {
			return nil, err
		}
		m[uint8(i)] = v
	}
	return m, nil
}

func writeUInt8Array(w io.Writer, v []uint8) error {
	_, err := w.Write(v)
	return err
}

// write a int to string map to writer
func writeUInt8StringMap(w io.Writer, m map[uint8]string) error {
	for _, v := range m {
		_, err := w.Write([]byte(v))
		if err != nil {
			return err
		}
	}
	return nil
}

// write a int to int map to writer
func writeUInt8UInt8Map(w io.Writer, m map[uint8]uint8) error {
	for _, v := range m {
		_, err := w.Write([]byte{v})
		if err != nil {
			return err
		}
	}
	return nil
}

func makeFieldValueArray(a []uint8) []FieldValue {
	b := make([]FieldValue, len(a))
	for i := range b {
		b[i] = FieldValue(a[i])
	}
	return b
}

func makeUInt8Array(a []FieldValue) []uint8 {
	b := make([]uint8, len(a))
	for i := range b {
		b[i] = uint8(a[i])
	}
	return b
}

func makeStringArray(a []FieldValue) []string {
	b := make([]string, len(a))
	for i := range b {
		b[i] = a[i].String()
	}
	return b
}

// make an int to string map
func makeUInt8StringMap(n uint8) map[uint8]string {
	b := make(map[uint8]string)
	for i := uint8(0); i < n; i++ {
		b[i] = ""
	}
	return b
}

// make an int to int map
func makeUInt8UInt8Map(n uint8) map[uint8]uint8 {
	b := make(map[uint8]uint8)
	for i := uint8(0); i < n; i++ {
		b[i] = 0
	}
	return b
}


func computeFinalBalances(bals channel.Balances, winner channel.Index) channel.Balances {
	loser := 1 - winner
	finalBals := bals.Clone()
	for i := range finalBals {
		finalBals[i][winner] = new(big.Int).Add(bals[i][0], bals[i][1])
		finalBals[i][loser] = big.NewInt(0)
	}
	return finalBals
}

// check if two maps are equal except for the given key
func equalExcept(d1, d2 map[uint8]uint8, key uint8) bool {
	for k, v := range d1 {
		if k != key && d2[k] != v {
			return false
		}
	}
	return true
}

// map to string
func mapToString(m map[uint8]uint8) string {
	s := ""
	for k, v := range m {
		s += fmt.Sprintf("%v: %v, ", k, v)
	}
	return s
}
