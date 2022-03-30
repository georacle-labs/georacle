package chain

// ID is a chain's unique identifier
type ID uint64

// Type denotes a family of chains
type Type uint64

// Config represents a set of chain parameters
type Config struct {
	ID   ID     // chain id
	Name string // canonical name
	Type Type   // underlying runtime
	Test bool   // true if testnet
}

const (
	// Ethereum mainnet
	Ethereum ID = 1

	// Ropsten testnet
	Ropsten ID = 3

	// Rinkeby testnet
	Rinkeby ID = 4

	// Kovan testnet
	Kovan ID = 42
)

const (
	// EVM compatible chains
	EVM Type = 1
)

var (
	// Chains denotes all valid chains
	Chains = map[string]*Config{
		"eth": {
			ID:   Ethereum,
			Name: "mainnet",
			Type: EVM,
			Test: false,
		},

		"ropsten": {
			ID:   Ropsten,
			Name: "ropsten",
			Type: EVM,
			Test: true,
		},

		"rinkeby": {
			ID:   Rinkeby,
			Name: "rinkeby",
			Type: EVM,
			Test: true,
		},

		"kovan": {
			ID:   Kovan,
			Name: "kovan",
			Type: EVM,
			Test: true,
		},
	}
)
