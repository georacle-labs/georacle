package evm

import (
	"os"
	"testing"

	"github.com/georacle-labs/go-georacle/chains"
)

var (
	URI = os.Getenv("ETH_WS_URI")
)

func TestConnection(t *testing.T) {
	c := NewClient(chains.KovanChainParams, URI)
	if c == nil {
		t.Fatal(c)
	}

	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	defer c.Close()
}
