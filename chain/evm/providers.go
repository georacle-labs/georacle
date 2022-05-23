package evm

import (
	"encoding/hex"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/georacle-labs/georacle/node"
)

const (
	// GasLimit is the default gas limit used to send transactions
	GasLimit = uint64(300000)
)

var (
	// ProvidersAddr is a temporary placeholder for the contract address
	ProvidersAddr = common.HexToAddress(os.Getenv("PROVIDERS"))
)

// ProviderManager subscribes to join/exit events
func (c *Client) ProviderManager(p chan node.Peer) error {
	joins := make(chan *ProvidersProviderJoin)
	exits := make(chan *ProvidersProviderExit)
	watchOpts := &bind.WatchOpts{Context: c.Ctx, Start: nil}

	// monitor join events
	joinSub, err := c.Providers.WatchProviderJoin(watchOpts, joins, nil, nil, nil)
	if err != nil {
		return err
	}
	defer joinSub.Unsubscribe()

	// monitor exit events
	exitSub, err := c.Providers.WatchProviderExit(watchOpts, exits, nil, nil, nil)
	if err != nil {
		return err
	}
	defer exitSub.Unsubscribe()

	for {
		select {
		case join := <-joins:
			id := hex.EncodeToString(append(join.Pubkey[:], join.NetAddr[:node.NetAddrSize]...))
			log.Printf("[%v] Provider Join (%s): 0x%s", c.ID, join.P, id)
			peer, err := node.NewPeer(join.Pubkey[:], join.NetAddr[:node.NetAddrSize])
			if err != nil {
				return err
			}
			p <- peer
		case exit := <-exits:
			id := hex.EncodeToString(append(exit.Pubkey[:], exit.NetAddr[:node.NetAddrSize]...))
			log.Printf("[%v] Provider Exit (%s): 0x%s", c.ID, exit.P, id)
			peer, err := node.NewPeer(exit.Pubkey[:], exit.NetAddr[:node.NetAddrSize])
			if err != nil {
				return err
			}
			p <- peer
		case joinErr := <-joinSub.Err():
			return joinErr
		case exitErr := <-exitSub.Err():
			return exitErr
		case <-c.Ctx.Done():
			return nil
		}
	}
}

// Join submits a tx to join the provider network
func (c *Client) Join(pubkey []byte, netAddr []byte) error {
	nonce, err := c.Client.PendingNonceAt(c.Ctx, c.Account.Address)
	if err != nil {
		return err
	}

	gasPrice, err := c.Client.SuggestGasPrice(c.Ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.Account.PrivateKey, c.ID)
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = GasLimit
	auth.GasPrice = gasPrice

	var pk, addr [32]byte
	copy(pk[:], pubkey[:])
	copy(addr[:], netAddr[:])

	tx, err := c.Providers.Join(auth, pk, addr)
	if err != nil {
		return err
	}

	log.Printf("[%v] Sent join tx %s\n", c.ID, tx.Hash().Hex())
	log.Printf("[%v] Joined provider network\n", c.ID)
	return nil
}

// Peers queries the provider network for a list of peers
func (c *Client) Peers() ([]node.Peer, error) {
	providerAddrs, err := c.Providers.GetProviders(nil)
	if err != nil {
		return nil, err
	}

	peers := make([]node.Peer, len(providerAddrs))
	for i, addr := range providerAddrs {
		p, err := c.Providers.Lookup(nil, addr)
		if err != nil {
			return nil, err
		}
		peer, err := node.NewPeer(p.Pubkey[:], p.NetAddr[:node.NetAddrSize])
		if err != nil {
			return nil, err
		}
		peers[i] = peer
		log.Printf("[Providers] Discovered peer (%d) %s\n", i, peer.ID())
	}

	return peers, nil
}
