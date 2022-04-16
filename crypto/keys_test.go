package crypto

import (
	"crypto/rand"
	"io"
	mr "math/rand"
	"reflect"
	"testing"
)

func TestKeyGen(t *testing.T) {
	k1 := KeyPair{T: EdDSA}
	k2 := KeyPair{T: EdDSA}

	for i := 0; i < (1 << 10); i++ {
		if err := k1.Gen(); err != nil {
			t.Fatal(err)
		}

		if err := k2.FromSeed(k1.Priv[:32]); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(k1, k2) {
			t.Fatalf("%v != %v\n", k1, k2)
		}
	}
}

func TestKeyEncrypt(t *testing.T) {
	for i := 0; i < 100; i++ {
		password := make([]byte, (mr.Uint32()%(1<<10))+1)
		if _, err := io.ReadFull(rand.Reader, password); err != nil {
			t.Fatal(err)
		}

		key, err := EdDSAGen()
		if err != nil {
			t.Fatal(err)
		}

		payload, err := key.Encrypt(password)
		if err != nil {
			t.Fatal(err)
		}

		k := new(KeyPair)
		if err := k.Decrypt(payload, password); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(key, k) {
			t.Fatalf("%v != %v\n", k, key)
		}
	}
}
