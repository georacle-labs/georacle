package evm

import (
	ctx "context"
	"log"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// Client wraps a connection to an evm compatible node
type Client struct {
	// ID is the chain ID
	ID *big.Int

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
}

// NewClient returns an empty evm client
func NewClient(id int64, uri string) (*Client, error) {
	wsURI, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	client := &Client{ID: big.NewInt(id), URI: wsURI, Alive: false}
	client.Ctx = ctx.Background()

	return client, nil
}

// Open initiaites a new RPC connection
func (c *Client) Open() error {
	// subsequent calls here are idempotent
	if !c.Alive {
		uri := c.URI.String()

		conn, err := rpc.DialWebsocket(c.Ctx, uri, "")
		if err != nil {
			return err
		}

		client := ethclient.NewClient(conn)

		cid := new(big.Int)
		if cid, err = client.ChainID(c.Ctx); err != nil {
			return errors.New("ChainIDValidationError")
		}

		// validate chain ID
		if c.ID.Cmp(cid) != 0 {
			return errors.Errorf("InvalidChainIDError: %v != %v", c.ID.Int64(), cid.Int64())
		}

		c.RPC = conn
		c.Client = client
		c.Alive = true
	}

	return nil
}

// Close terminates an existing RPC connection
func (c *Client) Close() error {
	// subsequent calls here are idempotent
	if c.Alive {
		c.RPC.Close()
		c.RPC = nil
		c.Client = nil
		c.Alive = false
	}
	return nil
}

// Run is this client's main blocking tx loop
func (c *Client) Run() error {
	c.Open()
	headers := make(chan *types.Header)

	sub, err := c.Client.SubscribeNewHead(c.Ctx, headers)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case header := <-headers:
			log.Printf("Received Block: %v\n", header.Hash().Hex())
			block, err := c.Client.BlockByNumber(c.Ctx, header.Number)
			if err != nil {
				return err
			}
			log.Println("Block Number: ", block.Number().Uint64())
			log.Println("Time Stamp: ", block.Time())
			log.Println("Nonce: ", block.Nonce())
			log.Println("Num Tx: ", len(block.Transactions()))
		}
	}
}
