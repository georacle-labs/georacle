package evm

import (
	"crypto/rand"
	"reflect"
	"testing"
)

func TestAccountGen(t *testing.T) {
	if _, err := Gen(); err != nil {
		t.Fatal(err)
	}
}

func TestAccountMarshal(t *testing.T) {
	a, err := Gen()
	if err != nil {
		t.Fatal(err)
	}

	r := make([]byte, 32)
	_, err = rand.Read(r)
	if err != nil {
		t.Fatal(err)
	}
	password := string(r)

	enc, err := a.Export(password)
	if err != nil {
		t.Fatal(err)
	}

	b := new(Account)
	err = b.Import(enc, password)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v != %v\n", a.String(), b.String())
	}
}
