package rpc

import (
	"context"
	"errors"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var (
	// ErrResponse is thrown on a connection time out
	ErrResponse = errors.New("no response from server")
)

// Client wraps an RPC client
type Client struct {
	client *rpc.Client // RPC client
}

// Init an RPC connection
func (c *Client) Init(addr string) (err error) {
	c.client, err = jsonrpc.Dial("tcp", addr)
	return
}

// Close an RPC connection
func (c *Client) Close() error {
	return c.client.Close()
}

// RPC services

// Ping an RPC server
func (c *Client) Ping(ctx context.Context) (ack bool, err error) {
	err = c.client.Call("Ping.Ping", nil, &ack)
	return
}
