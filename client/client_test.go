package client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	cfg, err := DefaultConfig()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := cfg.NewClient(); err != nil {
		t.Fatal(err)
	}
}
