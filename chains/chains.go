package chains

// ChainID is a chain's unique identifier
type ChainID uint64

// ChainType denotes a family of chains
type ChainType uint64

// ChainParams represents a generic chain
type ChainParams struct {
	ID   ChainID
	Type ChainType
	Name string
}

const (
	// Ethereum denotes the Ethereum mainnet
	Ethereum ChainID = 1

	// Ropsten denotes the Ropsten testnet
	Ropsten = 3

	// Rinkeby denotes the Rinkeby testnet
	Rinkeby = 4

	// Kovan denotes the Kovan testnet
	Kovan = 42
)

const (
	// EVM denotes EVM compatible chains
	EVM ChainType = iota
)

var (
	// EthereumChainParams is the chain parameters for the Ethereum mainnet
	EthereumChainParams = &ChainParams{
		ID:   Ethereum,
		Type: EVM,
		Name: "eth",
	}

	// RopstenChainParams is the chain parameters for the Ropsten testnet
	RopstenChainParams = &ChainParams{
		ID:   Ropsten,
		Type: EVM,
		Name: "ropsten",
	}

	// RinkebyChainParams is the chain parameters for the Rinkeby testnet
	RinkebyChainParams = &ChainParams{
		ID:   Rinkeby,
		Type: EVM,
		Name: "rinkeby",
	}

	// KovanChainParams is the chain parameters for the Kovan testnet
	KovanChainParams = &ChainParams{
		ID:   Kovan,
		Type: EVM,
		Name: "kovan",
	}
)
