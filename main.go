package main

import (
	"log"
	"os"

	"github.com/georacle-labs/go-georacle/chains"
	"github.com/georacle-labs/go-georacle/chains/evm"
)

var (
	// URI is the ws uri of an eth node
	URI = os.Getenv("ETH_WS_URI")
)

func main() {
	c := evm.NewClient(chains.KovanChainParams, URI)

	if err := c.Open(); err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	c.Run()
}
