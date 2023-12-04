package client

import (
	"context"
	"fmt"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/perun-examples/app-channel/app"
)

// FLChannel is a wrapper for a Perun channel for the Tic-tac-toe app use case.
type FLChannel struct {
	ch *client.Channel
}

// newFLChannel creates a new tic-tac-toe app channel.
func newFLChannel(ch *client.Channel) *FLChannel {
	return &FLChannel{ch: ch}
}

// Set sends a game move to the channel peer.

func (g *FLChannel) Set(model, numberOfRounds, weight, accuracy, loss int) {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		app, ok := state.App.(*app.FLApp)
		if !ok {
			return fmt.Errorf("invalid app type: %T", app)
		}

		return app.Set(state, model, numberOfRounds, weight, accuracy, loss, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// ForceSet registers a game move on-chain.
func (g *FLChannel) ForceSet(model, numberOfRounds, weight, accuracy, loss int) {
	err := g.ch.ForceUpdate(context.TODO(), func(state *channel.State) {
		err := func() error {
			app, ok := state.App.(*app.FLApp)
			if !ok {
				return fmt.Errorf("invalid app type: %T", app)
			}

			return app.Set(state, model, numberOfRounds, weight, accuracy, loss, g.ch.Idx())
		}()
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
}

// Settle settles the app channel and withdraws the funds.
func (g *FLChannel) Settle() {
	// Channel should be finalized through last ("winning") move.
	// No need to set `isFinal` here.
	err := g.ch.Settle(context.TODO(), false)
	if err != nil {
		panic(err)
	}

	// Cleanup.
	g.ch.Close()
}
