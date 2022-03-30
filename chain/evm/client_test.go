package evm

import (
	"os"
	"testing"
)

var (
	URI = os.Getenv("ETH_WS_URI")
)

func TestConnection(t *testing.T) {
	c, err := NewClient(42, URI)
	if err != nil {
		t.Fatal(c)
	}

	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	defer c.Close()
}
