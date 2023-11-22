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
	Weight 	  [1]uint8
	Accuracy  [1]uint8
	Loss 	  [1]uint8
}

func (d *FLAppData) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "model: %v\n", d.Model)
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

func (d *FLAppData) Set(model, weight, accuracy, loss int, actorIdx channel.Index) {
	if d.NextActor != uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	// v := makeFieldValueFromPlayerIdx(actorIdx)
	// d.Grid[y*3+x] = v

	if d.NextActor == 0 {
		d.Model = uint8(model)
		d.Accuracy[0] = uint8(accuracy)
		d.Loss[0] = uint8(loss)

	} else {
		d.Weight[0] = uint8(weight)
	}

	d.NextActor = calcNextActor(d.NextActor)
}

func calcNextActor(actor uint8) uint8 {
	return (actor + 1) % numParts
}
