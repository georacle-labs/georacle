package chain

import (
	"testing"
)

func TestChainType(t *testing.T) {
	for _, c := range Chains {
		if len(c.Type.String()) <= 0 {
			t.Fatal("invalid chain type")
		}
	}

	c := Config{}
	if len(c.Type.String()) > 0 {
		t.Fatal("expected empty chain type")
	}
}
