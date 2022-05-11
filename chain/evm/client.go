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
	"github.com/pkg/errors"
)

var (
	// ErrInvalidChainID is thrown on an invalid ETH chain ID
	ErrInvalidChainID = errors.New("Invalid Chain ID ")

	// ErrInvalidClientState is thrown on an invalid client state
	ErrInvalidClientState = errors.New("Invalid Client State")
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
			return ErrInvalidChainID
		}
		if c.ID.Cmp(cid) != 0 {
			return errors.Wrapf(ErrInvalidChainID, "%v != %v", c.ID.Int64(), cid.Int64())
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

	return ErrInvalidClientState
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
	return ErrInvalidClientState
}

// Run is the client's main blocking tx loop
func (c *Client) Run() (err error) {
	var wg sync.WaitGroup

	// start the provider manager
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = c.ProviderManager()
	}()

	log.Printf("[%v] Started Provider manager %s", c.ID, ProvidersAddr.Hex())
	wg.Wait()
	return
}
