package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"

	"github.com/georacle-labs/georacle/crypto/cipher"
)

// GenKeyPair returns a new ed25519 keypair
func GenKeyPair() (pub, priv []byte, err error) {
	pub, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return
}

// Encrypt a byte slice
func Encrypt(data, password []byte) ([]byte, error) {
	var c cipher.Cipher
	if err := c.Encrypt(data, password); err != nil {
		return nil, err
	}

	payload, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// Decrypt a byte slice
func Decrypt(data, password []byte) ([]byte, error) {
	var c cipher.Cipher

	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	data, err := c.Decrypt(password)
	if err != nil {
		return nil, err
	}

	return data, nil
}
