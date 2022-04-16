package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/scrypt"
)

const (

	// BlockSize for aes256-gcm is 32 bytes
	BlockSize = (1 << 5)

	// ScryptN is the CPU/memory cost parameter for Scrypt
	ScryptN = (1 << 18)

	// ScryptR is the blocksize parameter for Scrypt
	ScryptR = (1 << 3)

	// ScryptP is the parallelization parameter for Scrypt
	ScryptP = 1
)

// Cipher wraps an aes256-gcm block cipher
// with scrypt for key derivation
type Cipher struct {
	Payload []byte `json:"payload"`
	IV      []byte `json:"iv"`
	Salt    []byte `json:"salt"`
	Nonce   []byte `json:"nonce"`
	N       int    `json:"n"`
	R       int    `json:"r"`
	P       int    `json:"p"`
}

// Encrypt a byte slice with aes256-gcm
func (c *Cipher) Encrypt(data, password []byte) error {
	c.Salt = make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, c.Salt); err != nil {
		return err
	}

	c.Nonce = make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, c.Nonce); err != nil {
		return err
	}

	dk, err := c.KDF(password)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(dk)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	c.Payload = gcm.Seal(nil, c.Nonce, data, nil)
	return nil
}

// Decrypt a byte slice with aes256-gcm
func (c *Cipher) Decrypt(password []byte) ([]byte, error) {
	dk, err := c.KDF(password)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(dk)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	data, err := gcm.Open(nil, c.Nonce, c.Payload, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// KDF runs Scrypt with standard params by default
// See https://www.tarsnap.com/scrypt/scrypt.pdf
func (c *Cipher) KDF(password []byte) (dk []byte, err error) {
	// fallback to standard params if empty
	if c.N <= 0 {
		c.N = ScryptN
	}
	if c.R <= 0 {
		c.R = ScryptR
	}
	if c.P <= 0 {
		c.P = ScryptP
	}
	return scrypt.Key(password, c.Salt, c.N, c.R, c.P, BlockSize)
}
