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
	"reflect"

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

func (a *FLApp) InitData(firstActor channel.Index, modelCID, numberOfRounds uint8) *FLAppData {
	return &FLAppData{
		NextActor: uint8(firstActor),
		ModelCID: uint8(modelCID),
		RoundPhase: uint8(0),
		NumberOfRounds: uint8(numberOfRounds),
		Round: uint8(0),
		Weights: makeUInt8UInt8Map(uint8(numberOfRounds)),
		Accuracy: makeUInt8UInt8Map(uint8(numberOfRounds)),
		Loss: makeUInt8UInt8Map(uint8(numberOfRounds)),
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

	d.ModelCID, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading model")
	}

	d.NumberOfRounds, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading number of rounds")
	}

	d.Round, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading round")
	}

	d.RoundPhase, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading round phase")
	}

	d.Weights, err = readUInt8UInt8Map(r, int(d.NumberOfRounds))
	if err != nil {
		return nil, errors.WithMessage(err, "reading weight")
	}

	d.Accuracy, err = readUInt8UInt8Map(r, int(d.NumberOfRounds))
	if err != nil {
		return nil, errors.WithMessage(err, "reading accuracy")
	}

	d.Loss, err = readUInt8UInt8Map(r, int(d.NumberOfRounds))
	if err != nil {
		return nil, errors.WithMessage(err, "reading loss")
	}

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

	// zero := FLAppData{}
	// if appData.Grid != zero.Grid {
	// 	return fmt.Errorf("invalid starting grid: %v", appData.Grid)
	// }

	if appData.ModelCID <= uint8(0) {
		return fmt.Errorf("invalid starting model: %v", appData.ModelCID)
	}

	if appData.NumberOfRounds <= uint8(0) {
		return fmt.Errorf("invalid starting number of rounds: %v", appData.NumberOfRounds)
	}

	if appData.Round != uint8(0) {
		return fmt.Errorf("invalid starting round: %v", appData.Round)
	}

	if !reflect.DeepEqual(appData.Weights, makeUInt8UInt8Map(uint8(appData.NumberOfRounds))){
		return fmt.Errorf("invalid starting weights: %v", appData.Weights)
	}

	if !reflect.DeepEqual(appData.Accuracy, makeUInt8UInt8Map(uint8(appData.NumberOfRounds))){
		return fmt.Errorf("invalid starting accuracy: %v", appData.Accuracy)
	}

	if !reflect.DeepEqual(appData.Loss, makeUInt8UInt8Map(uint8(appData.NumberOfRounds))){
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

	// Check data.
	if fromData.ModelCID != toData.ModelCID {
		return fmt.Errorf("cannot override model: expected %v, got %v", fromData.ModelCID, toData.ModelCID)
	}

	if fromData.NumberOfRounds != toData.NumberOfRounds {
		return fmt.Errorf("cannot override number of rounds: expected %v, got %v", fromData.NumberOfRounds, toData.NumberOfRounds)
	}

	if toData.Round > toData.NumberOfRounds {
		return fmt.Errorf("round out of bounds: %v", toData.Round)
	}

	if fromData.NextActor == uint8(1){ //Client conditions
		if fromData.Round != toData.Round{
			return fmt.Errorf("actor: %v cannot override round: expected %v, got %v", fromData.NextActor, fromData.Round, toData.Round)
		}
		if !reflect.DeepEqual(fromData.Accuracy, toData.Accuracy){
			return fmt.Errorf("actor: %v cannot override accuracy: expected %v, got %v", fromData.NextActor, mapToString(fromData.Accuracy), mapToString(toData.Accuracy))
		}
		if !reflect.DeepEqual(fromData.Loss, toData.Loss){
			return fmt.Errorf("actor: %v cannot override loss: expected %v, got %v", fromData.NextActor, mapToString(fromData.Loss), mapToString(toData.Loss))
		}

		if !equalExcept(fromData.Weights, toData.Weights, toData.Round){
			return fmt.Errorf("actor: %v cannot override weights outside current round: expected %v, got %v", fromData.NextActor, mapToString(fromData.Weights), mapToString(toData.Weights))
		}

		if toData.Weights[fromData.Round] == 0 { //weight is not set
			return fmt.Errorf("actor: %v cannot skip weight %v -> %v", fromData.NextActor, mapToString(fromData.Weights), mapToString(toData.Weights))
		}
	}

	if fromData.NextActor == uint8(0){ //Server conditions
		if toData.Round != fromData.Round+1{
			return fmt.Errorf("actor: %v must increment round: expected %v, got %v", fromData.NextActor, fromData.Round+1, toData.Round)
		}
		if !reflect.DeepEqual(fromData.Weights, toData.Weights){
			return fmt.Errorf("actor: %v cannot override weights: expected %v, got %v", fromData.NextActor, fromData.Weights, toData.Weights)
		}
		if !equalExcept(fromData.Accuracy, toData.Accuracy, toData.Round){
			return fmt.Errorf("actor: %v cannot override accuracy outside current round: expected %v, got %v", fromData.NextActor, mapToString(fromData.Accuracy), mapToString(toData.Accuracy))
		}

		if toData.Accuracy[fromData.Round] == 0 { //accuracy is not set
			return fmt.Errorf("actor: %v cannot skip accuracy", fromData.NextActor)
		}

		if !equalExcept(fromData.Loss, toData.Loss, toData.Round){
			return fmt.Errorf("actor: %v cannot override loss outside current round: expected %v, got %v", fromData.NextActor, mapToString(fromData.Loss), mapToString(toData.Loss))
		}

		if toData.Loss[fromData.Round] == 0{ //loss is not set
			return fmt.Errorf("actor: %v cannot skip loss", fromData.NextActor)
		}
	}

	// if fromData.RoundPhase != toData.RoundPhase {
	// 	return fmt.Errorf("cannot repeat round phase: %v -> %v", fromData.RoundPhase, toData.RoundPhase)
	// }

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


func (a *FLApp) Set(s *channel.State, weight, accuracy, loss int, actorIdx channel.Index) error {
	d, ok := s.Data.(*FLAppData)
	if !ok {
		return fmt.Errorf("invalid data type: %T", d)
	}

	d.Set(weight, accuracy, loss, actorIdx)
	log.Println("\n" + d.String())

	if isFinal, winner := d.CheckFinal(); isFinal {
		s.IsFinal = true
		if winner != nil {
			s.Balances = computeFinalBalances(s.Balances, *winner)
		}
	}
	return nil
}