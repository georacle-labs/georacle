package chain

import (
	"github.com/georacle-labs/go-georacle/chain/evm"
	"github.com/pkg/errors"
)

// Chain defines the interface to a generic chain
type Chain interface {
	Open() error
	Close() error
	Run() error
}

// New returns a new chain corresponding to a chain config
func New(cfg *Config, uri string) (Chain, error) {
	switch cfg.Type {
	case EVM:
		return evm.NewClient(int64(cfg.ID), uri)
	default:
		return nil, errors.Errorf("Invalid chain: `%v`\n", cfg.Name)
	}
}
