package accounts

import (
	"context"
	"crypto/rand"
	"io"
	mr "math/rand"
	"testing"

	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NumEntries  = (1 << 10)
	ElementSize = (1 << 12) // 4K
)

var (
	Entries [NumEntries]primitive.ObjectID
)

func GenTestAccountStore() (*Store, error) {
	db := &db.DB{Name: "test"}
	if err := db.Open(dbURI); err != nil {
		return nil, err
	}

	store := &Store{
		AccountType: chain.EVM,
		Ctx:         context.TODO(),
	}

	err := store.Init(db.Accounts)
	return store, err
}

func TestAddEntry(t *testing.T) {
	as, err := GenTestAccountStore()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < NumEntries; i++ {
		eLen := (mr.Uint32() % ElementSize) + 1
		entry := make([]byte, eLen)
		io.ReadFull(rand.Reader, entry)

		id, err := as.AddEntry(entry)
		if err != nil {
			t.Fatal(err)
		}
		Entries[i] = id
	}

	count := 0
	cursor, err := as.Accounts.Find(as.Ctx, bson.M{})
	if err != nil {
		t.Fatal(err)
	}
	defer cursor.Close(as.Ctx)

	for cursor.Next(as.Ctx) {
		count++
	}

	if count != NumEntries {
		t.Fatalf("Insufficient Count: %v != %v", count, NumEntries)
	}
}

func TestRemoveEntry(t *testing.T) {
	as, err := GenTestAccountStore()
	if err != nil {
		t.Fatal(err)
	}

	removed := 0
	for _, entry := range Entries {
		if err := as.RemoveEntry(entry); err != nil {
			t.Fatal(err)
		}
		removed++
	}

	if removed != NumEntries {
		t.Fatalf("Insufficient Count: %v != %v", removed, NumEntries)
	}
}
