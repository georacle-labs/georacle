package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"

	"filippo.io/edwards25519"
	"github.com/pkg/errors"
)

var (
	// ErrKeyLen is thrown on an invalid key length
	ErrKeyLen = errors.New("KeyLengthError")

	// ErrKeyType is thrown on an invalid key type
	ErrKeyType = errors.New("InvalidKeyTypeError")
)

// KeyType is used to specify a signing algorithm
type KeyType uint16

const (
	// EdDSA is used to generate ed25519 keys
	EdDSA = 1
)

// KeyPair represents a generic key pair
type KeyPair struct {
	T    KeyType `json:"type"`
	Pub  []byte  `json:"pub"`
	Priv []byte  `json:"priv"`
}

// Gen generates a fresh key pair
func (k *KeyPair) Gen() (err error) {
	switch k.T {
	case EdDSA:
		k.Pub, k.Priv, err = ed25519.GenerateKey(rand.Reader)
	default:
		return ErrKeyType
	}
	return
}

// FromSeed initilizes a key pair from a 32 byte seed
func (k *KeyPair) FromSeed(priv []byte) (err error) {
	switch k.T {
	case EdDSA:
		if len(priv) != ed25519.SeedSize {
			return ErrKeyLen
		}
		k.Priv = ed25519.NewKeyFromSeed(priv)
		k.Pub = k.Priv[ed25519.PublicKeySize:]
	default:
		return ErrKeyType
	}
	return
}

// Decrypt an encrypted key pair
func (k *KeyPair) Decrypt(payload, password []byte) (err error) {
	var data []byte
	if data, err = Decrypt(payload, password); err == nil {
		err = json.Unmarshal(data, k)
	}
	return
}

// Encrypt a key pair
func (k *KeyPair) Encrypt(password []byte) ([]byte, error) {
	data, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	return Encrypt(data, password)
}

// EdDSAGen generates an ed25519 key pair
func EdDSAGen() (*KeyPair, error) {
	k := KeyPair{T: EdDSA}
	if err := k.Gen(); err != nil {
		return nil, err
	}
	return &k, nil
}

// ValidEdDSA checks for a valid ed25519 pubkey
func ValidEdDSA(pubkey []byte) bool {
	// ensure 64 byte encoding
	if len(pubkey) != ed25519.PublicKeySize {
		return false
	}

	// ensure point lies on the curve
	_, err := (&edwards25519.Point{}).SetBytes(pubkey)
	return err == nil
}
