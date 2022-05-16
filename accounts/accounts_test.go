package accounts

import (
	"crypto/rand"
	"io"
	mr "math/rand"
	"os"
	"testing"

	"github.com/georacle-labs/georacle/accounts/evm"
	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/db"
)

var (
	dbURI = os.Getenv("DB_URI")
)

const (
	NumAccounts = (1 << 5)
)

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

func TestInvalidPassword(t *testing.T) {
	db := &db.DB{Name: "test"}
	if err := db.Open(dbURI); err != nil {
		t.Fatal(err)
	}

	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	// valid insert should pass
	if err := ks.NewAccount(); err != nil {
		t.Fatal(err)
	}

	ksInvalid, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	// overwrite password
	if _, err := io.ReadFull(rand.Reader, ksInvalid.Password); err != nil {
		t.Fatal(err)
	}

	// decryption should fail
	if err := ksInvalid.Init(db.Accounts); err == nil {
		t.Fatal("expected authentication error")
	}

	// valid remove should pass
	if err := ks.RemoveAccount(0); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidChain(t *testing.T) {
	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	ks.Type = 0 // invalid chain type
	if err := ks.NewAccount(); err == nil {
		t.Fatal("expected invalid chain type")
	}
}

func TestDuplicateInsert(t *testing.T) {
	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	account := new(evm.Account)
	if err := account.Gen(); err != nil {
		t.Fatal(err)
	}

	// first insert should pass
	if err := ks.InsertAccount(account); err != nil {
		t.Fatal(err)
	}

	// duplicate insert should fail
	if err := ks.InsertAccount(account); err != ErrDuplicate {
		t.Fatal("expected duplicate insert error")
	}
}

func TestInvalidRemove(t *testing.T) {
	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	// out of bound index should fail
	if err := ks.RemoveAccount(1); err != ErrIndex {
		t.Fatal("expected index out of range error")
	}

	// first remove should pass
	if err := ks.RemoveAccount(0); err != nil {
		t.Fatal(err)
	}

	// next one should fail
	if err := ks.RemoveAccount(0); err != ErrIndex {
		t.Fatal("expected index out of range error")
	}
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

	ks.DumpAccounts()
}

func TestAccountRemove(t *testing.T) {
	ks, err := GenTestKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	for len(ks.Entries) > 0 {
		nEntries := len(ks.Entries)
		randEntry := mr.Intn(nEntries)
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

	ks.DumpAccounts()
}
