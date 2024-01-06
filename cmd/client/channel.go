package client

import (
	"context"
	"fmt"
	"math/big"
	
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/perun-examples/app-channel/cmd/app"
	"perun.network/go-perun/log"

)

// FLChannel is a wrapper for a Perun channel for the Tic-tac-toe app use case.
type FLChannel struct {
	ch *client.Channel
	log     log.Logger
}

// newFLChannel creates a new tic-tac-toe app channel.
func newFLChannel(ch *client.Channel) *FLChannel {
	return &FLChannel{
		ch: 	   ch,
		log:       log.WithField("channel", ch.ID()),
	}
}


// func stateBals(state *channel.State) []channel.Bal {
// 	return state.Balances[0]
// }

// Set sends a game move to the channel peer.
func (g *FLChannel) Set(model, numberOfRounds, weight, accuracy, loss int) error {
	g.log.Debugf("Setting state: %s", model)
	ctx, cancel := context.WithTimeout(context.Background(), config.Channel.Timeout)
	defer cancel()

	err := g.ch.UpdateBy(ctx, func(state *channel.State) error {
		app, ok := state.App.(*app.FLApp)
		if !ok {
			return fmt.Errorf("invalid app type: %T", app)
		}

 		return app.Set(state, model, numberOfRounds, weight, accuracy, loss, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
	return err
}

// ForceSet registers a game move on-chain.
func (g *FLChannel) ForceSet(model, numberOfRounds, weight, accuracy, loss int) error {
	g.log.Debugf("Force setting state: %s", model)
	ctx, cancel := context.WithTimeout(context.Background(), config.Channel.Timeout)
	defer cancel()

	err := g.ch.ForceUpdate(ctx, func(state *channel.State) {
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

	return err
}

// Settle settles the app channel and withdraws the funds.
func (g *FLChannel) Settle() {
	// Channel should be finalized through last ("winning") move.
	// No need to set `isFinal` here.
	g.log.Debugf("Settle channel: %s", g.ch.ID())
	ctx, cancel := context.WithTimeout(context.Background(), config.Channel.Timeout)
	defer cancel()

	err := g.ch.Settle(ctx, false)
	if err != nil {
		panic(err)
	}

	// Cleanup.
	g.ch.Close()
}


func (g *FLChannel) GetBalances() (our, other *big.Int) {
	bals := stateBals(g.ch.State())
	if len(bals) != 2 {
		return new(big.Int), new(big.Int)
	}
	return bals[g.ch.Idx()], bals[1-g.ch.Idx()]
}