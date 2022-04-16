package crypto

import (
	"bytes"
	"crypto/rand"
	"io"
	mr "math/rand"
	"testing"
)

const (
	Iter = (1 << 10)
)

func TestCipher(t *testing.T) {
	c := Cipher{N: ScryptN >> 10}

	for i := 0; i < Iter; i++ {
		payload := make([]byte, (mr.Uint32()%(1<<12))+1)
		if _, err := io.ReadFull(rand.Reader, payload); err != nil {
			t.Fatal(err)
		}
		password := make([]byte, (mr.Uint32()%(1<<10))+1)
		if _, err := io.ReadFull(rand.Reader, password); err != nil {
			t.Fatal(err)
		}

		if err := c.Encrypt(payload, password); err != nil {
			t.Fatal(err)
		}
		decrypted, err := c.Decrypt(password)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(payload, decrypted) != 0 {
			t.Fatal("invalid decryption")
		}

		if _, err := io.ReadFull(rand.Reader, password); err != nil {
			t.Fatal(err)
		}
		decrypted, err = c.Decrypt(password)
		if err == nil {
			t.Fatal("Expected error")
		}
	}
}
