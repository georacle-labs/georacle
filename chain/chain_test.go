package chain

import (
	"os"
	"testing"
)

var (
	WSURI = os.Getenv("ETH_WS_URI")
)

func TestNewValidChain(t *testing.T) {
	for _, c := range Chains {
		_, err := New(c, WSURI)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNewInvalidChain(t *testing.T) {
	cfg := &Config{}
	c, err := New(cfg, WSURI)
	if err == nil || c != nil {
		t.Fatal("expected error thrown")
	}
}
