package evm

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// ProvidersAddr is a temporary placeholder for the contract address
	ProvidersAddr = common.HexToAddress(os.Getenv("PROVIDERS"))
)

// ProviderManager subscribes to join/exit events
func (c *Client) ProviderManager() error {
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
	exitSub, err := c.Providers.WatchProviderExit(watchOpts, exits, nil)
	if err != nil {
		return err
	}
	defer exitSub.Unsubscribe()

	for {
		select {
		case join := <-joins:
			log.Printf("[%v] Provider Join: %s", c.ID, join.Pubkey)
		case exit := <-exits:
			log.Printf("[%v] Provider Exit: %s", c.ID, exit.P)
		}
	}
}
