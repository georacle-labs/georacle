package evm

import (
	"os"
	"testing"
)

var (
	wsURI = os.Getenv("WS_URI")
)

func TestConnection(t *testing.T) {
	c, err := NewClient(42, wsURI)
	if err != nil {
		t.Fatal(c)
	}

	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	defer c.Close()
}
