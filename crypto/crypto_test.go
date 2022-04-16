package crypto

import (
	"bytes"
	"crypto/rand"
	"io"
	mr "math/rand"
	"testing"
)

func TestEncrypt(t *testing.T) {
	k := KeyPair{T: EdDSA}
	for i := 0; i < (1 << 5); i++ {
		if err := k.Gen(); err != nil {
			t.Fatal(err)
		}
		data := append(k.Pub, k.Priv...)
		password := make([]byte, (mr.Uint32()%(1<<10))+1)
		if _, err := io.ReadFull(rand.Reader, password); err != nil {
			t.Fatal(err)
		}

		payload, err := Encrypt(data, password)
		if err != nil {
			t.Fatal(err)
		}

		decrypted, err := Decrypt(payload, password)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(data, decrypted) != 0 {
			t.Fatal("invalid decryption")
		}

		if _, err := io.ReadFull(rand.Reader, password); err != nil {
			t.Fatal(err)
		}
		decrypted, err = Decrypt(data, password)
		if err == nil {
			t.Fatal("Expected error")
		}
	}
}
