// Copyright (c) 2019 Chair of Applied Cryptography, Technische Universität
// Darmstadt, Germany. All rights reserved. This file is part of
// perun-eth-demo. Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	_ "perun.network/go-perun/backend/ethereum" // backend init
	ethchannel "perun.network/go-perun/backend/ethereum/channel"
	ethwallet "perun.network/go-perun/backend/ethereum/wallet"
	phd "perun.network/go-perun/backend/ethereum/wallet/hd"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
	"perun.network/go-perun/wallet"
	"perun.network/go-perun/wire"
	wirenet "perun.network/go-perun/wire/net"
	"perun.network/go-perun/wire/net/simple"

	"perun.network/perun-examples/app-channel/cmd/app"
)

type peer struct {
	alias   string
	perunID wire.Address
	ch      *FLChannel
	log     log.Logger
}

type node struct {
	log log.Logger

	bus    *wirenet.Bus
	client *client.Client
	dialer *simple.Dialer

	// Account for signing on-chain TX. Currently also the Perun-ID.
	onChain *phd.Account
	// Account for signing off-chain TX. Currently one Account for all
	// state channels, later one we want one Account per Channel.
	offChain wallet.Account
	wallet   *phd.Wallet

	adjudicator channel.Adjudicator
	adjAddr     common.Address
	asset       channel.Asset
	assetAddr   common.Address
	funder      channel.Funder
	appAddr     common.Address

	// App client
	stake       channel.Bal    // The amount we put at stake.
	app         *app.FLApp     // The app definition.

	// Needed to deploy contracts.
	cb ethchannel.ContractBackend

	// Protects peers
	mtx   sync.Mutex
	peers map[string]*peer
}

func getOnChainBal(ctx context.Context, addrs ...wallet.Address) ([]*big.Int, error) {
	bals := make([]*big.Int, len(addrs))
	var err error
	for idx, addr := range addrs {
		bals[idx], err = ethereumBackend.BalanceAt(ctx, ethwallet.AsEthAddr(addr), nil)
		if err != nil {
			return nil, errors.Wrap(err, "querying on-chain balance")
		}
	}
	return bals, nil
}

func (n *node) Connect(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	return n.connect(args[0])
}

func (n *node) connect(alias string) error {
	n.log.Traceln("Connecting...")
	if n.peers[alias] != nil {
		return errors.New("Peer already connected")
	}
	peerCfg, ok := config.Peers[alias]
	if !ok {
		return errors.Errorf("Alias '%s' unknown. Add it to 'config/network.yaml'.", alias)
	}

	n.dialer.Register(peerCfg.perunID, peerCfg.Hostname+":"+strconv.Itoa(int(peerCfg.Port)))

	n.peers[alias] = &peer{
		alias:   alias,
		perunID: peerCfg.perunID,
		log:     log.WithField("peer", peerCfg.perunID),
	}

	fmt.Printf("📡 Connected to %v. Ready to open channel.\n", alias)

	return nil
}

// peer returns the peer with the address `addr` or nil if not found.
func (n *node) peer(addr wire.Address) *peer {
	for _, peer := range n.peers {
		if peer.perunID.Equals(addr) {
			return peer
		}
	}
	return nil
}

func (n *node) channelPeer(ch *client.Channel) *peer {
	perunID := ch.Peers()[1-ch.Idx()] // assumes two-party channel
	return n.peer(perunID)
}

func (n *node) setupChannel(ch *client.Channel) {
	if len(ch.Peers()) != 2 {
		log.Fatal("Only channels with two participants are currently supported")
	}

	perunID := ch.Peers()[1-ch.Idx()] // assumes two-party channel
	p := n.peer(perunID)

	if p == nil {
		log.WithField("peer", perunID).Warn("Opened channel to unknown peer")
		return
	} else if p.ch != nil {
		p.log.Warn("Peer tried to open more than one channel")
		return
	}

	p.ch = newFLChannel(ch)

	// Start watching.
	go func() {
		l := log.WithField("channel", ch.ID())
		l.Debug("Watcher started")
		err := ch.Watch(n)
		l.WithError(err).Debug("Watcher stopped")
	}()

	bals := weiToEther(ch.State().Balances[0]...)
	fmt.Printf("🆕 Channel established with %s. Initial balance: [My: %v Ξ, Peer: %v Ξ]\n",
		p.alias, bals[ch.Idx()], bals[1-ch.Idx()]) // assumes two-party channel
}

func (n *node) HandleAdjudicatorEvent(e channel.AdjudicatorEvent) {
	if _, ok := e.(*channel.ConcludedEvent); ok {
		PrintfAsync("🎭 Received concluded event\n")
		func() {
			n.mtx.Lock()
			defer n.mtx.Unlock()
			ch := n.channel(e.ID())
			if ch == nil {
				// If we initiated the channel closing, then the channel should
				// already be removed and we return.
				return
			}
			peer := n.channelPeer(ch.ch)
			if err := n.settle(peer); err != nil {
				PrintfAsync("🎭 error while settling: %v\n", err)
			}
			PrintfAsync("🏁 Settled channel with %s.\n", peer.alias)
		}()
	}
}

type balTuple struct {
	My, Other *big.Int
}

func (n *node) GetBals() map[string]balTuple {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	bals := make(map[string]balTuple)
	for alias, peer := range n.peers {
		if peer.ch != nil {
			my, other := peer.ch.GetBalances()
			bals[alias] = balTuple{my, other}
		}
	}
	return bals
}

func findConfig(id wallet.Address) (string, *netConfigEntry) {
	for alias, e := range config.Peers {
		if e.perunID.String() == id.String() {
			return alias, e
		}
	}
	return "", nil
}

func (n *node) HandleUpdate(_ *channel.State, update client.ChannelUpdate, resp *client.UpdateResponder) {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	log := n.log.WithField("channel", update.State.ID)
	log.Debug("Channel update")

	// ch := n.channel(update.State.ID)
	// if ch == nil {
	// 	log.Error("Channel for ID not found")
	// 	return
	// }
	// ch.ch.Handle(update, resp)

	ctx, cancel := context.WithTimeout(context.Background(), config.Node.HandleTimeout)
	defer cancel()

	err := resp.Accept(ctx)
	if err != nil {
		panic(err)
	}
}

func (n *node) channel(id channel.ID) *FLChannel {
	for _, p := range n.peers {
		if p.ch != nil && p.ch.ch.ID() == id {
			return p.ch
		}
	}
	return nil
}

func (n *node) HandleProposal(prop client.ChannelProposal, res *client.ProposalResponder) {
	req, ok := prop.(*client.LedgerChannelProposal)
	if !ok {
		log.Fatal("Can handle only ledger channel proposals.")
	}

	if len(req.Peers) != 2 {
		log.Fatal("Only channels with two participants are currently supported")
	}

	n.mtx.Lock()
	defer n.mtx.Unlock()
	id := req.Peers[0]
	n.log.Debug("Received channel proposal")

	// Find the peer by its perunID and create it if not present
	p := n.peer(id)
	alias, cfg := findConfig(id)

	if p == nil {
		if cfg == nil {
			ctx, cancel := context.WithTimeout(context.Background(), config.Node.HandleTimeout)
			defer cancel()
			if err := res.Reject(ctx, "Unknown identity"); err != nil {
				n.log.WithError(err).Warn("rejecting")
			}
			return
		}
		p = &peer{
			alias:   alias,
			perunID: id,
			log:     log.WithField("peer", id),
		}
		n.peers[alias] = p
		n.log.WithField("channel", id).WithField("alias", alias).Debug("New peer")
	}
	n.log.WithField("peer", id).Debug("Channel proposal")

	bals := weiToEther(req.InitBals.Balances[0]...)
	theirBal := bals[0] // proposer has index 0
	ourBal := bals[1]   // proposal receiver has index 1

	ctx, cancel := context.WithTimeout(context.Background(), config.Node.HandleTimeout)
	defer cancel()
	// Check that the channel has sufficient funding balances.
	stakeFloat := new(big.Float).SetInt(n.stake) // Convert n.stake to *big.Float
	fmt.Printf("🔁 Incoming channel proposal from %v with funding [My: %v Ξ, Peer: %v Ξ].\n", alias, ourBal, theirBal)
	fmt.Printf("Stake: %v Ξ\n", stakeFloat)
	if theirBal.Cmp(stakeFloat) > 0 && ourBal.Cmp(stakeFloat) > 0 {
		fmt.Printf("✅ Channel proposal accepted. Opening channel...\n")
		a := req.Accept(n.offChain.Address(), client.WithRandomNonce())
		if _, err := res.Accept(ctx, a); err != nil {
			n.log.Error(errors.WithMessage(err, "accepting channel proposal"))
			return
		}
		return
	}

	fmt.Printf("❌ Channel proposal rejected\n")
	if err := res.Reject(ctx, "Invalid funding balance"); err != nil {
		n.log.Error(errors.WithMessage(err, "rejecting channel proposal"))
		return
	}

}

func (n *node) Open(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	peerName := args[0]
	peer := n.peers[peerName]
	if peer == nil {
		// try to connect to peer
		if err := n.connect(peerName); err != nil {
			return err
		}
		peer = n.peers[peerName]
	}
	myBalEth, _ := new(big.Float).SetString(args[1]) // Input was already validated by command parser.
	peerBalEth, _ := new(big.Float).SetString(args[2])

	initBals := &channel.Allocation{
		Assets:   []channel.Asset{n.asset},
		Balances: [][]*big.Int{etherToWei(myBalEth, peerBalEth)},
	}

	firstActorIdx := channel.Index(0)
	withApp := client.WithApp(n.app, n.app.InitData(firstActorIdx))

	prop, err := client.NewLedgerChannelProposal(
		config.Channel.ChallengeDurationSec,
		n.offChain.Address(),
		initBals,
		[]wire.Address{n.onChain.Address(), peer.perunID},
		client.WithRandomNonce(),
		withApp,
	)
	if err != nil {
		return errors.WithMessage(err, "creating channel proposal")
	}

	fmt.Printf("💭 Proposing channel to %v...\n", peerName)

	ctx, cancel := context.WithTimeout(context.Background(), config.Channel.FundTimeout)
	defer cancel()
	n.log.Debug("Proposing channel")
	ch, err := n.client.ProposeChannel(ctx, prop)
	if err != nil {
		return errors.WithMessage(err, "proposing channel failed")
	}
	if n.channel(ch.ID()) == nil {
		return errors.New("OnNewChannel handler could not setup channel")
	}
	return nil
}

// func (n *node) Send(args []string) error {
// 	n.mtx.Lock()
// 	defer n.mtx.Unlock()
// 	n.log.Traceln("Sending...")

// 	peer := n.peers[args[0]]
// 	if peer == nil {
// 		return errors.Errorf("peer not found %s", args[0])
// 	} else if peer.ch == nil {
// 		return errors.Errorf("connect to peer first")
// 	}
// 	amountEth, _ := new(big.Float).SetString(args[1]) // Input was already validated by command parser.
// 	return peer.ch.sendMoney(etherToWei(amountEth)[0])
// }

func (n *node) Set(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("Sending...")

	peer := n.peers[args[0]]
	if peer == nil {
		return errors.Errorf("peer not found %s", args[0])
	} else if peer.ch == nil {
		return errors.Errorf("connect to peer first")
	}

	model, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.WithMessage(err, "converting model string to int")
	}

	numberOfRounds, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.WithMessage(err, "converting numberOfRounds string to int")
	}

	weight, err := strconv.Atoi(args[3])
	if err != nil {
		return errors.WithMessage(err, "converting weight string to int")
	}

	accuracy, err := strconv.Atoi(args[4])
	if err != nil {
		return errors.WithMessage(err, "converting accuracy string to int")
	}

	loss, err := strconv.Atoi(args[5])
	if err != nil {
		return errors.WithMessage(err, "converting loss string to int")
	}

	err = peer.ch.Set(model, numberOfRounds, weight, accuracy, loss)
	return err
}


func (n *node) ForceSet(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("Sending...")

	peer := n.peers[args[0]]
	if peer == nil {
		return errors.Errorf("peer not found %s", args[0])
	} else if peer.ch == nil {
		return errors.Errorf("connect to peer first")
	}

	model, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.WithMessage(err, "converting model string to int")
	}

	numberOfRounds, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.WithMessage(err, "converting numberOfRounds string to int")
	}

	weight, err := strconv.Atoi(args[3])
	if err != nil {
		return errors.WithMessage(err, "converting weight string to int")
	}

	accuracy, err := strconv.Atoi(args[4])
	if err != nil {
		return errors.WithMessage(err, "converting accuracy string to int")
	}

	loss, err := strconv.Atoi(args[5])
	if err != nil {
		return errors.WithMessage(err, "converting loss string to int")
	}

	err = peer.ch.ForceSet(model, numberOfRounds, weight, accuracy, loss)
	return err
}

func (n *node) Close(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("Closing...")

	alias := args[0]
	peer := n.peers[alias]
	if peer == nil {
		return errors.Errorf("Unknown peer: %s", alias)
	}
	// if err := peer.ch.sendFinal(); err != nil {
	// 	return errors.WithMessage(err, "sending final state for state closing")
	// }

	if err := n.settle(peer); err != nil {
		return errors.WithMessage(err, "settling")
	}
	fmt.Printf("\r🏁 Settled channel with %s.\n", peer.alias)
	return nil
}

func (n *node) settle(p *peer) error {
	p.ch.log.Debug("Settling")
	ctx, cancel := context.WithTimeout(context.Background(), config.Channel.SettleTimeout)
	defer cancel()

	if err := p.ch.ch.Settle(ctx, p.ch.ch.Idx() == 0); err != nil {
		return errors.WithMessage(err, "settling the channel")
	}

	if err := p.ch.ch.Close(); err != nil {
		return errors.WithMessage(err, "channel closing")
	}
	p.ch.log.Debug("Removing channel")
	p.ch = nil
	return nil
}

// Info prints the phase of all channels.
func (n *node) Info(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("Info...")

	ctx, cancel := context.WithTimeout(context.Background(), config.Chain.TxTimeout)
	defer cancel()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "Peer\tPhase\tVersion\tMy Ξ\tPeer Ξ\tMy On-Chain Ξ\tPeer On-Chain Ξ\t\n")
	for alias, peer := range n.peers {
		onChainBals, err := getOnChainBal(ctx, n.onChain.Address(), peer.perunID)
		if err != nil {
			return err
		}
		onChainBalsEth := weiToEther(onChainBals...)
		if peer.ch == nil {
			fmt.Fprintf(w, "%s\t%s\t \t \t \t%v\t%v\t\n", alias, "Connected", onChainBalsEth[0], onChainBalsEth[1])
		} else {
			bals := weiToEther(peer.ch.GetBalances())
			fmt.Fprintf(w, "%s\t%v\t%d\t%v\t%v\t%v\t%v\t\n",
				alias, peer.ch.ch.Phase(), peer.ch.ch.State().Version, bals[0], bals[1], onChainBalsEth[0], onChainBalsEth[1])
		}
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}

func (n *node) Exit([]string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("Exiting...")

	return n.client.Close()
}

func (n *node) ExistsPeer(alias string) bool {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	return n.peers[alias] != nil
}
