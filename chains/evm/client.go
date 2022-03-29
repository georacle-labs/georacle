package evm

import (
	ctx "context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/georacle-labs/go-georacle/chains"
)

// Client wraps a connection to an evm compatible node
type Client struct {
	// Config is the chain config
	Config *chains.ChainParams

	// URI is a valid endpoint
	URI *url.URL

	// Connection is a WebSocket connection to an Ethereum node
	RPC *rpc.Client

	// Client is an Ethereum client
	Client *ethclient.Client

	// Context is the client's calling context
	Ctx ctx.Context

	// Wg is the client's running go routines
	Wg *sync.WaitGroup

	// Alive is true if the node is connected
	Alive bool
}

// NewClient returns an empty RPC client
func NewClient(c *chains.ChainParams, uri string) *Client {
	wsURI, err := url.Parse(uri)
	if err != nil {
		return nil
	}

	client := &Client{Config: c, URI: wsURI, Alive: false}
	client.Ctx = ctx.Background()
	client.Wg = new(sync.WaitGroup)

	return client
}

// Open initiaites an RPC connection
func (c *Client) Open() error {
	if c.Alive {
		return errors.New("ClientAlreadyConnectedError")
	}

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
	if chains.ChainID(cid.Int64()) != c.Config.ID {
		return errors.New("InvalidChainIDError")
	}

	c.RPC = conn
	c.Client = client
	c.Alive = true

	return nil
}

// Close terminates an RPC connection
func (c *Client) Close() error {
	if !c.Alive {
		return errors.New("ClientDisconnectedError")
	}

	c.Wg.Wait()

	c.RPC.Close()
	c.Client = nil
	c.Alive = false

	return nil
}

// Run is the client's main tx loop
func (c *Client) Run() error {
	headers := make(chan *types.Header)

	sub, err := c.Client.SubscribeNewHead(c.Ctx, headers)
	if err != nil {
		return err
	}

	c.Wg.Add(1)

	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		for {
			select {
			case err := <-sub.Err():
				return err
			case header := <-headers:
				log.Printf("Received Block: %v\n", header.Hash().Hex())
				block, err := c.Client.BlockByNumber(c.Ctx, header.Number)
				if err != nil {
					fmt.Println("err = ", err)
					return err
				}
				log.Println("Block Number: ", block.Number().Uint64())
				log.Println("Time Stamp: ", block.Time())
				log.Println("Nonce: ", block.Nonce())
				log.Println("Num Tx: ", len(block.Transactions()))
			}
		}
	}(c.Wg)

	return nil
}
