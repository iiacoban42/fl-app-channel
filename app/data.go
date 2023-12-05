package app

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
)

// FLAppData is the app data struct.
// Grid:
// 0 1 2
// 3 4 5
// 6 7 8
type FLAppData struct {
	NextActor uint8
	// Grid      [9]FieldValue
	Model 	  uint8
	NumberOfRounds uint8
	Round 	   uint8
	RoundPhase uint8
	Weight 	  [3]uint8
	Accuracy  [3]uint8
	Loss 	  [3]uint8
}

func (d *FLAppData) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "model: %v\n", d.Model)
	fmt.Fprintf(&b, "numberOfRounds: %v\n", d.NumberOfRounds)
	fmt.Fprintf(&b, "round: %v\n", d.Round)
	fmt.Fprintf(&b, "roundPhase: %v\n", d.RoundPhase)
	fmt.Fprintf(&b, "weight: %v\n", d.Weight)
	fmt.Fprintf(&b, "accuracy: %v\n", d.Accuracy)
	fmt.Fprintf(&b, "loss: %v\n", d.Loss)
	fmt.Fprintf(&b, "Next actor: %v\n", d.NextActor)
	return b.String()
}

// Encode encodes app data onto an io.Writer.
func (d *FLAppData) Encode(w io.Writer) error {
	err := writeUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}

	err = writeUInt8(w, d.Model)
	if err != nil {
		return errors.WithMessage(err, "writing model")
	}

	err = writeUInt8(w, d.NumberOfRounds)
	if err != nil {
		return errors.WithMessage(err, "writing numberOfRounds")
	}

	err = writeUInt8(w, d.Round)
	if err != nil {
		return errors.WithMessage(err, "writing round")
	}

	err = writeUInt8(w, d.RoundPhase)
	if err != nil {
		return errors.WithMessage(err, "writing roundPhase")
	}

	err = writeUInt8Array(w, makeUInt8Array(d.Weight[:]))
	if err != nil {
		return errors.WithMessage(err, "writing weight")
	}

	err = writeUInt8Array(w, makeUInt8Array(d.Accuracy[:]))
	if err != nil {
		return errors.WithMessage(err, "writing accuracy")
	}

	err = writeUInt8Array(w, makeUInt8Array(d.Loss[:]))
	return errors.WithMessage(err, "writing loss")

	// err = writeUInt8Array(w, makeUInt8Array(d.Grid[:]))
	// return errors.WithMessage(err, "writing grid")

}

// Clone returns a deep copy of the app data.
func (d *FLAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *FLAppData) Set(model, numberOfRounds, weight, accuracy, loss int, actorIdx channel.Index) {
	if d.NextActor != uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	// v := makeFieldValueFromPlayerIdx(actorIdx)
	// d.Grid[y*3+x] = v

	if d.NextActor != uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	// v := makeFieldValueFromPlayerRoundPhaserIdx(actorIdx)
	// d.Grid[y*3+x] = v

	if d.RoundPhase == 0{ // init then waiting for updates
		// require that the round is 0
		if d.Round != 0 {
			fmt.Printf("round: %v\n", d.Round)
			panic("invalid round")
		}
		d.Model = uint8safe(uint16(model)) // set the model
		d.NumberOfRounds = uint8safe(uint16(numberOfRounds)) // set the number of rounds
		d.RoundPhase = 1

	} else if d.RoundPhase == 1 { //update then waiting for aggregation
		d.Weight[d.Round] = uint8safe(uint16(weight))
		d.RoundPhase = 2

	}else if d.RoundPhase == 2 { // aggregate then waiting for updates
		d.Accuracy[d.Round] = uint8safe(uint16(accuracy))
		d.Loss[d.Round] = uint8safe(uint16(loss))

		if d.Round == d.NumberOfRounds - 1 {
			d.RoundPhase = 3
		} else if d.Round < d.NumberOfRounds - 1 {
		d.RoundPhase = 1
		} else {
			panic("invalid round")
		}
		d.Round = uint8safe(uint16(d.Round + 1))
		
	} else if d.RoundPhase == 3 { //waiting for termination
		fmt.Printf("waiting for termination")
		d.RoundPhase = 4
	} else {
		panic("game already over")
	}

	d.NextActor = calcNextActor(d.NextActor)

}

func calcNextActor(actor uint8) uint8 {
	return (actor + 1) % numParts
}
