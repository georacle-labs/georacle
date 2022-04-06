package cmd

import (
	"testing"

	"github.com/georacle-labs/georacle/client"
)

func TestNewClient(t *testing.T) {
	cfg, err := client.DefaultConfig()
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Network) <= 0 {
		t.Fatal(cfg.Network)
	}

	if len(cfg.WSURI) <= 0 {
		t.Fatal(cfg.WSURI)
	}

	if len(cfg.DBURI) <= 0 {
		t.Fatal(cfg.DBURI)
	}
}
