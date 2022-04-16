package accounts

import (
	"math/rand"
	"os"
	"testing"

	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/db"
)

var dbURI = os.Getenv("DB_URI")

const NumAccounts = 100

func GenTestKeyStore() (*Master, error) {
	db := &db.DB{Name: "test"}
	if err := db.Open(dbURI); err != nil {
		return nil, err
	}

	master := &Master{
		Type:     chain.EVM,
		Password: []byte("SuperSecurePassword"),
	}

	err := master.Init(db.Accounts)
	return master, err
}

func TestAccountGen(t *testing.T) {
	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < NumAccounts; i++ {
		if err := ks.NewAccount(); err != nil {
			t.Fatal(err)
		}

		numAccounts := len(ks.Entries)
		if numAccounts != i+1 {
			t.Fatalf("%v != %v\n", numAccounts, i+1)
		}
	}
}

func TestAccountRemove(t *testing.T) {
	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	for len(ks.Entries) > 0 {
		nEntries := len(ks.Entries)
		randEntry := rand.Intn(nEntries)
		toRemove := ks.Entries[randEntry]

		if err := ks.RemoveAccount(uint(randEntry)); err != nil {
			t.Fatal(err)
		}

		if len(ks.Entries) != nEntries-1 {
			t.Fatalf("%v != %v\n", len(ks.Entries), nEntries-1)
		}

		for _, e := range ks.Entries {
			if e.Account.String() == toRemove.Account.String() {
				t.Fatalf("Expected removal of entry: %v\n", e)
			}
		}
	}
}
