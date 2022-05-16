package client

import (
	"testing"
)

func GenTestClient() (*Client, error) {
	cfg, err := DefaultConfig()
	if err != nil {
		return nil, err
	}
	return cfg.NewClient()

}

func TestNewClient(t *testing.T) {
	_, err := GenTestClient()
	if err != nil {
		t.Fatal(err)
	}
}
