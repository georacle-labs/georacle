package client

import (
	"sync"

	"github.com/georacle-labs/georacle/accounts"
	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/db"
)

// Client represents a single client instance
type Client struct {
	Params   *chain.Config   // chain params
	Wg       *sync.WaitGroup // client's running go routines
	Chain    chain.Chain     // corresponding node connection
	DB       *db.DB          // open db handle
	Alive    bool            // true if client initialized
	Accounts accounts.Master // client's accounts
}

// Init Client
func (c *Client) Init() error {
	if !c.Alive {
		if err := c.Accounts.Init(); err != nil {
			return err
		}
		c.Alive = true
	}
	return nil
}

// Start the client
func (c *Client) Start() error {
	if err := c.Init(); err != nil {
		return err
	}
	return c.Chain.Run()
}

// Stop the client
func (c *Client) Stop() {
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
