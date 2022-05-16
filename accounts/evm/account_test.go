package evm

import (
	"crypto/rand"
	"reflect"
	"testing"
)

func TestAccountGen(t *testing.T) {
	a := new(Account)
	if err := a.Gen(); err != nil {
		t.Fatal(err)
	}
}

func TestAccountMarshal(t *testing.T) {
	a := new(Account)
	if err := a.Gen(); err != nil {
		t.Fatal(err)
	}

	password := make([]byte, 32)
	_, err := rand.Read(password)
	if err != nil {
		t.Fatal(err)
	}

	enc, err := a.Export(password)
	if err != nil {
		t.Fatal(err)
	}

	b := new(Account)
	err = b.Import(enc, password)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(a, b) || a.String() != b.String() {
		t.Fatalf("%v != %v\n", a.String(), b.String())
	}
}
