package chain

import (
	"os"
	"testing"
)

var (
	wsURI = os.Getenv("WS_URI")
)

func TestNewValidChain(t *testing.T) {
	for _, c := range Chains {
		_, err := New(c, wsURI)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNewInvalidChain(t *testing.T) {
	cfg := &Config{}
	c, err := New(cfg, wsURI)
	if err == nil || c != nil {
		t.Fatal("expected error thrown")
	}
}
