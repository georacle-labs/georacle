package chain

import (
	"github.com/georacle-labs/georacle/chain/evm"
	"github.com/georacle-labs/georacle/node"
	"github.com/pkg/errors"
)

var (
	// ErrChain is thrown on an unsupported chain
	ErrChain = errors.New("invalid chain")
)

// Chain defines the interface for a generic chain
type Chain interface {
	// Open the chain
	Open() error

	// Close the chain
	Close() error

	// Run the on chain protocol
	Run() error

	// Join the provider network
	Join([]byte, []byte) error

	// Peers queries the provider network for a list of peers
	Peers() ([]node.Peer, error)
}

// New returns a new chain corresponding to a chain config
func New(cfg *Config, uri string) (Chain, error) {
	switch cfg.Type {
	case EVM:
		return evm.NewClient(int64(cfg.ID), uri)
	default:
		return nil, errors.Wrapf(ErrChain, ": %v", cfg.Name)
	}
}
