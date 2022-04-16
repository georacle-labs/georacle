package client

import (
	"sync"

	"github.com/georacle-labs/georacle/accounts"
	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/db"
	"github.com/georacle-labs/georacle/node"
)

// Client represents a single client instance
type Client struct {
	Params   *chain.Config   // chain params
	Chain    chain.Chain     // ws node connection
	DB       *db.DB          // open db handle
	Alive    bool            // true iff client initialized
	Accounts accounts.Master // global account store
	Node     node.Node       // RPC node
	Wg       sync.WaitGroup  // Running GoRoutines
}

// Init Client
func (c *Client) Init() error {
	if !c.Alive {
		// init the account store
		if err := c.Accounts.Init(c.DB.Accounts); err != nil {
			return err
		}
		// init the node store
		if err := c.Node.Init(c.DB.Node, c.Accounts.Password); err != nil {
			return err
		}
		c.Alive = true
	}
	return nil
}

// Start the client
func (c *Client) Start() (err error) {
	if err := c.Init(); err != nil {
		return err
	}

	errs := make(chan error, 2)

	// start underlying chain
	c.Wg.Add(1)
	go func() {
		defer c.Wg.Done()
		errs <- c.Chain.Run()
	}()

	// start RPC node
	c.Wg.Add(1)
	go func() {
		defer c.Wg.Done()
		errs <- c.Node.Start()
	}()

	c.Wg.Wait()

	close(errs)
	for e := range errs {
		if e != nil {
			return e
		}
	}

	return nil
}

// Stop the client
func (c *Client) Stop() {
	c.Node.Stop()
	c.Chain.Close()
	c.DB.Close()
}

// ListAccounts lists accounts to stdout
func (c *Client) ListAccounts() error {
	if err := c.Init(); err != nil {
		return err
	}
	c.Accounts.DumpAccounts()
	return nil
}

// NewAccount generates a new account and updates the keystore
func (c *Client) NewAccount() error {
	if err := c.Init(); err != nil {
		return err
	}
	return c.Accounts.NewAccount()
}

// RemoveAccount removes an account from the keystore by index
func (c *Client) RemoveAccount(index uint) error {
	if err := c.Init(); err != nil {
		return err
	}
	return c.Accounts.RemoveAccount(index)
}
