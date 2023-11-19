package app

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
)

// FLAppData is the app data struct.
//
// RoundPhase: 0 = Init, 1 = WaitForUpdates, 2 = WaitForAggregation, 3 = WaitForTermination, 4 = Done
type FLAppData struct {
	RoundPhase uint8
	NextActor uint8
	NumberOfRounds uint8
	Round uint8
	ModelCID uint8
	Weights map[uint8]uint8
	Accuracy map[uint8]uint8
	Loss map[uint8]uint8

}

func (d *FLAppData) String() string {
	var b bytes.Buffer
	// print Model
	fmt.Fprintf(&b, "model: %v\n", d.ModelCID)
	// print Round
	fmt.Fprintf(&b, "round: %v out of: %v\n", d.Round, d.NumberOfRounds)
	// print RoundPhase
	fmt.Fprintf(&b, "roundPhase: %v\n", d.RoundPhase)
	// print Weights map
	fmt.Fprintf(&b, "weights: %v\n", d.Weights)
	// print Accuracy map
	fmt.Fprintf(&b, "accuracy: %v\n", d.Accuracy)
	// print Loss map
	fmt.Fprintf(&b, "loss: %v\n", d.Loss)
	// print NextActor
	fmt.Fprintf(&b, "nextActor: %v\n", d.NextActor)
	return b.String()
}

// Encode encodes app data onto an io.Writer.
func (d *FLAppData) Encode(w io.Writer) error {
	err := writeUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}

	err = writeUInt8(w, d.ModelCID)
	if err != nil {
		return errors.WithMessage(err, "writing model")
	}

	err = writeUInt8(w, d.NumberOfRounds)
	if err != nil {
		return errors.WithMessage(err, "writing number of rounds")
	}

	err = writeUInt8(w, d.Round)
	if err != nil {
		return errors.WithMessage(err, "writing round")
	}

	err = writeUInt8(w, d.RoundPhase)
	if err != nil {
		return errors.WithMessage(err, "writing round phase")
	}

	err = writeUInt8UInt8Map(w, d.Weights)
	if err != nil {
		return errors.WithMessage(err, "writing weight")
	}

	err = writeUInt8UInt8Map(w, d.Accuracy)
	if err != nil {
		return errors.WithMessage(err, "writing accuracy")
	}

	err = writeUInt8UInt8Map(w, d.Loss)
	return errors.WithMessage(err, "writing loss")

	// err = writeUInt8Array(w, makeUInt8Array(d.Grid[:]))
	// return errors.WithMessage(err, "writing grid")

}

// Clone returns a deep copy of the app data.
func (d *FLAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *FLAppData) Set(weight, accuracy, loss int, actorIdx channel.Index) {
	if d.NextActor != uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	// v := makeFieldValueFromPlayerRoundPhaserIdx(actorIdx)
	// d.Grid[y*3+x] = v

	if d.RoundPhase == 0{ //waiting for updates
		// require that the round is 0
		if d.Round != 0 {
			fmt.Printf("round: %v\n", d.Round)
			panic("invalid round")
		}
		d.Weights[d.Round] = uint8safe(uint16(weight)) // set the weights
		d.RoundPhase = 1

	} else if d.RoundPhase == 1 { //waiting for aggregation
		d.Accuracy[d.Round] = uint8safe(uint16(accuracy))
		d.Loss[d.Round] = uint8safe(uint16(loss))

		if d.Round == d.NumberOfRounds - 1 {
			d.RoundPhase = 2
		} else if d.Round < d.NumberOfRounds - 1 {
		d.RoundPhase = 0
		} else {
			panic("invalid round")
		}
		d.Round = uint8safe(uint16(d.Round + 1))
	} else if d.RoundPhase == 2 { //waiting for termination
		fmt.Printf("waiting for termination")
		d.RoundPhase = 3
	} else {
		panic("game already over")
	}

	d.NextActor = calcNextActor(d.NextActor)
}

func calcNextActor(actor uint8) uint8 {
	return (actor + 1) % numParts
}
