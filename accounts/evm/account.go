package evm

import (
	"crypto/ecdsa"

	ks "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

// Account represents an Ethereum address
type Account struct {
	ID         uuid.UUID         `json:"id" `         // 16 byte random UUID (version 4)
	Address    common.Address    `json:"address"`     // 20 byte eip-55 address
	PrivateKey *ecdsa.PrivateKey `json:"private_key"` // 32 byte ecdsa private key
}

// Gen generates a new account
func (a *Account) Gen() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	priv, err := crypto.GenerateKey()
	if err != nil {
		return err
	}

	pubKey := priv.PublicKey
	addr := crypto.PubkeyToAddress(pubKey)

	a.ID = id
	a.Address = addr
	a.PrivateKey = priv
	return nil
}

// Export returns the account as encrypted JSON using standard scrypt params
func (a *Account) Export(password []byte) ([]byte, error) {
	key := &ks.Key{Id: a.ID, Address: a.Address, PrivateKey: a.PrivateKey}
	return ks.EncryptKey(key, string(password), ks.StandardScryptN, ks.StandardScryptP)
}

// Import populates an account from an encrypted JSON payload
func (a *Account) Import(payload, password []byte) error {
	key, err := ks.DecryptKey(payload, string(password))
	if err != nil {
		return err
	}

	a.ID = key.Id
	a.Address = key.Address
	a.PrivateKey = key.PrivateKey

	return nil
}

// String dumps a mixed-case checksum encoded address
func (a *Account) String() string {
	return a.Address.Hex()
}
