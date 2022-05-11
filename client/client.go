package client

import (
	"sync"

	"github.com/georacle-labs/georacle/accounts"
	ea "github.com/georacle-labs/georacle/accounts/evm"
	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/chain/evm"
	"github.com/georacle-labs/georacle/db"
	"github.com/georacle-labs/georacle/node"
	"github.com/pkg/errors"
)

var (
	// ErrType is thrown on an invalid type cast
	ErrType = errors.New("Invalid type cast")

	// ErrChain is thrown on an inivalid chain type
	ErrChain = errors.New("Invalid chain type")
)

// Client represents a single client instance
type Client struct {
	Params   *chain.Config   // chain params
	Chain    chain.Chain     // ws node connection
	DB       *db.DB          // open db handle
	Alive    bool            // true iff client initialized
	Accounts accounts.Master // global account store
	Node     node.Node       // RPC node
}

// Init Client
func (c *Client) Init() error {
	if !c.Alive {
		// decrypt account store
		if err := c.Accounts.Init(c.DB.Accounts); err != nil {
			return err
		}

		// set the default signing account
		if len(c.Accounts.Entries) > 0 {
			defaultAccount := c.Accounts.Entries[0].Account
			switch c.Params.Type {
			case chain.EVM:
				ethClient, ok := c.Chain.(*evm.Client)
				if !ok {
					return ErrType
				}
				ethAccount, ok := defaultAccount.(*ea.Account)
				if !ok {
					return ErrType
				}
				ethClient.Account = ethAccount
			default:
				return errors.Wrapf(ErrChain, " %v\n", c.Params.Type)
			}
		}

		// init the rpc node
		if err := c.Node.Init(c.DB.Node, c.Accounts.Password); err != nil {
			return err
		}

		// init the underlying chain
		if err := c.Chain.Open(); err != nil {
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

	var wg sync.WaitGroup
	errs := make(chan error, 2)

	// start underlying chain
	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- c.Chain.Run()
	}()

	// start RPC node
	wg.Add(1)
	go func() {
		defer wg.Done()
		errs <- c.Node.Start()
	}()

	wg.Wait()

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
