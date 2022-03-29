package accounts

// AccountType specifies the type of account (i.e. which chain it's intended for)
type AccountType uint64

const (
	// EVM represents an EIP55 compatible address
	EVM AccountType = iota
)

// Account represents a generic account
type Account interface {
	Export(string) ([]byte, error)
	Import([]byte, string) error
	String() string
}
