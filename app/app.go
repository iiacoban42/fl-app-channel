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
	"log"

	"github.com/pkg/errors"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
)

// FLApp is a channel app.
type FLApp struct {
	Addr wallet.Address
}

func NewFLApp(addr wallet.Address) *FLApp {
	return &FLApp{
		Addr: addr,
	}
}

// Def returns the app address.
func (a *FLApp) Def() wallet.Address {
	return a.Addr
}

func (a *FLApp) InitData(firstActor channel.Index) *FLAppData {
	return &FLAppData{
		NextActor: uint8(firstActor),
	}
}

// DecodeData decodes the channel data.
func (a *FLApp) DecodeData(r io.Reader) (channel.Data, error) {
	d := FLAppData{}

	var err error
	d.NextActor, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading actor")
	}

	d.Model, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading model")
	}

	d.NumberOfRounds, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading numberOfRounds")
	}

	d.Round, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading round")
	}

	d.RoundPhase, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading roundPhase")
	}

	weight, err := readUInt8Array(r, len(d.Weight))
	if err != nil {
		return nil, errors.WithMessage(err, "reading weight")
	}
	copy(d.Weight[:], weight)

	accuracy, err := readUInt8Array(r, len(d.Accuracy))
	if err != nil {
		return nil, errors.WithMessage(err, "reading accuracy")
	}
	copy(d.Accuracy[:], accuracy)

	loss, err := readUInt8Array(r, len(d.Loss))
	if err != nil {
		return nil, errors.WithMessage(err, "reading loss")
	}
	copy(d.Loss[:], loss)

	// grid, err := readUInt8Array(r, len(d.Grid))
	// if err != nil {
	// 	return nil, errors.WithMessage(err, "reading grid")
	// }
	// copy(d.Grid[:], makeFieldValueArray(grid))
	return &d, nil
}

// ValidInit checks that the initial state is valid.
func (a *FLApp) ValidInit(p *channel.Params, s *channel.State) error {
	if len(p.Parts) != numParts {
		return fmt.Errorf("invalid number of participants: expected %d, got %d", numParts, len(p.Parts))
	}

	appData, ok := s.Data.(*FLAppData)
	if !ok {
		return fmt.Errorf("invalid data type: %T", s.Data)
	}

	zero := FLAppData{}
	// if appData.Grid != zero.Grid {
	// 	return fmt.Errorf("invalid starting grid: %v", appData.Grid)
	// }

	if appData.Model != zero.Model {
		return fmt.Errorf("invalid starting model: %v", appData.Model)
	}

	if appData.NumberOfRounds != zero.NumberOfRounds {
		return fmt.Errorf("invalid starting numberOfRounds: %v", appData.NumberOfRounds)
	}

	if appData.Round != zero.Round {
		return fmt.Errorf("invalid starting round: %v", appData.Round)
	}

	if appData.RoundPhase != zero.RoundPhase {
		return fmt.Errorf("invalid starting roundPhase: %v", appData.RoundPhase)
	}

	if appData.Weight != zero.Weight {
		return fmt.Errorf("invalid starting weight: %v", appData.Weight)
	}

	if appData.Accuracy != zero.Accuracy {
		return fmt.Errorf("invalid starting accuracy: %v", appData.Accuracy)
	}

	if appData.Loss != zero.Loss {
		return fmt.Errorf("invalid starting loss: %v", appData.Loss)
	}

	if s.IsFinal {
		return fmt.Errorf("must not be final")
	}

	if appData.NextActor >= numParts {
		return fmt.Errorf("invalid next actor: got %d, expected < %d", appData.NextActor, numParts)
	}
	return nil
}

// ValidTransition is called whenever the channel state transitions.
func (a *FLApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {
	err := channel.AssetsAssertEqual(from.Assets, to.Assets)
	if err != nil {
		return fmt.Errorf("Invalid assets: %v", err)
	}

	fromData, ok := from.Data.(*FLAppData)
	if !ok {
		panic(fmt.Sprintf("from state: invalid data type: %T", from.Data))
	}

	toData, ok := to.Data.(*FLAppData)
	if !ok {
		panic(fmt.Sprintf("to state: invalid data type: %T", from.Data))
	}

	// Check actor.
	if fromData.NextActor != uint8safe(uint16(idx)) {
		return fmt.Errorf("invalid actor: expected %v, got %v", fromData.NextActor, idx)
	}

	// Check next actor.
	if len(params.Parts) != numParts {
		panic("invalid number of participants")
	}
	expectedToNextActor := calcNextActor(fromData.NextActor)
	if toData.NextActor != expectedToNextActor {
		return fmt.Errorf("invalid next actor: expected %v, got %v", expectedToNextActor, toData.NextActor)
	}

	// if fromData.NextActor == 0 {
	// 	// Check model.
	// 	if toData.Model > maxFieldValue {
	// 		return fmt.Errorf("invalid model value: %d", toData.Model)
	// 	}


	// Check grid.
	// changed := false
	// for i, v := range toData.Grid {
	// 	if v > maxFieldValue {
	// 		return fmt.Errorf("invalid grid value at index %d: %d", i, v)
	// 	}
	// 	vFrom := fromData.Grid[i]
	// 	if v != vFrom {
	// 		if vFrom != notSet {
	// 			return fmt.Errorf("cannot overwrite field %d", i)
	// 		}
	// 		if changed {
	// 			return fmt.Errorf("cannot change two fields")
	// 		}
	// 		changed = true
	// 	}
	// }
	//
	// if !changed {
	// 	return fmt.Errorf("cannot skip turn")
	// }

	// check round inputs



		// Check model.
	// if toData.Model > maxFieldValue {
	// 		return fmt.Errorf("invalid model value: %d", toData.Model)
	// 	}


	// Check final and allocation.
	isFinal, winner := toData.CheckFinal()
	if to.IsFinal != isFinal {
		return fmt.Errorf("final flag: expected %v, got %v", isFinal, to.IsFinal)
	}
	expectedAllocation := from.Allocation.Clone()
	if winner != nil {
		expectedAllocation.Balances = computeFinalBalances(from.Allocation.Balances, *winner)
	}
	if err := expectedAllocation.Equal(&to.Allocation); err != nil {
		return errors.WithMessagef(err, "wrong allocation: expected %v, got %v", expectedAllocation, to.Allocation)
	}
	return nil
}


func (a *FLApp) Set(s *channel.State, model, numberOfRounds, weight, accuracy, loss int, actorIdx channel.Index) error {
	d, ok := s.Data.(*FLAppData)
	if !ok {
		return fmt.Errorf("invalid data type: %T", d)
	}

	d.Set(model, numberOfRounds, weight, accuracy, loss, actorIdx)
	log.Println("\n" + d.String())

	if isFinal, winner := d.CheckFinal(); isFinal {
		s.IsFinal = true
		if winner != nil {
			s.Balances = computeFinalBalances(s.Balances, *winner)
		}
	}
	return nil
}
