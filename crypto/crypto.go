package crypto

import (
	"encoding/json"
)

// Encrypt a byte slice
func Encrypt(data, password []byte) ([]byte, error) {
	var c Cipher
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
	var c Cipher

	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	data, err := c.Decrypt(password)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// EdDSAGen generates an ed25519 key pair
func EdDSAGen() (*KeyPair, error) {
	k := KeyPair{T: EdDSA}
	if err := k.Gen(); err != nil {
		return nil, err
	}
	return &k, nil
}
