package evm

import (
	ctx "context"
	"log"
	"math/big"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/georacle-labs/georacle/accounts/evm"
	"github.com/georacle-labs/georacle/node"
	"github.com/pkg/errors"
)

var (
	// ErrChainID is thrown on an invalid ETH chain ID
	ErrChainID = errors.New("invalid chain id")

	// ErrClientState is thrown on an invalid client state
	ErrClientState = errors.New("invalid client state")
)

// Client wraps a connection to an evm compatible node
type Client struct {
	// ID is the chain ID
	ID *big.Int

	// Account is the eth account used for signing transactions
	Account *evm.Account

	// URI is a valid endpoint
	URI *url.URL

	// Connection is a WebSocket connection to an Ethereum node
	RPC *rpc.Client

	// Client is an Ethereum client
	Client *ethclient.Client

	// Context is the client's calling context
	Ctx ctx.Context

	// Alive is true if the node is connected
	Alive bool

	// Providers smart contract instance
	Providers *Providers
}

// NewClient returns an empty evm client
func NewClient(chainID int64, uri string) (*Client, error) {
	wsURI, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	client := &Client{ID: big.NewInt(chainID), URI: wsURI, Alive: false}
	client.Ctx = ctx.Background()
	return client, nil
}

// Open connects to an eth client
func (c *Client) Open() error {
	if !c.Alive {
		// dial eth client
		uri := c.URI.String()
		conn, err := rpc.DialWebsocket(c.Ctx, uri, "")
		if err != nil {
			return err
		}
		client := ethclient.NewClient(conn)

		// validate chain ID
		cid := new(big.Int)
		if cid, err = client.ChainID(c.Ctx); err != nil {
			return ErrChainID
		}
		if c.ID.Cmp(cid) != 0 {
			return errors.Wrapf(ErrChainID, "%v != %v", c.ID.Int64(), cid.Int64())
		}

		// load providers contract instance
		c.Providers, err = NewProviders(ProvidersAddr, client)
		if err != nil {
			return err
		}

		c.RPC = conn
		c.Client = client
		c.Alive = true

		log.Printf("Started chain with ID %v", c.ID)
		return nil
	}

	return ErrClientState
}

// Close terminates an existing RPC connection
func (c *Client) Close() error {
	if c.Alive {
		c.RPC.Close()
		c.RPC = nil
		c.Client = nil
		c.Alive = false
		return nil

	}
	return ErrClientState
}

// Run is the client's main blocking tx loop
func (c *Client) Run(p chan node.Peer) (err error) {
	if !c.Alive {
		return ErrClientState
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = c.ProviderManager(p)
	}()

	log.Printf("[%v] Started Provider manager %s", c.ID, ProvidersAddr.Hex())

	wg.Wait()
	return
}
