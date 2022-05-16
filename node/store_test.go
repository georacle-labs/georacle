package node

import (
	"os"
	"testing"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/georacle-labs/georacle/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NumEntries = (1 << 5)
)

var (
	TestPassword = []byte("Sup3rS3cureP@ssword!")
	Entries      [NumEntries]primitive.ObjectID
	dbURI        = os.Getenv("DB_URI")
)

func GenTestNodeStore() (*Store, error) {
	db := &db.DB{Name: "test"}
	if err := db.Open(dbURI); err != nil {
		return nil, err
	}

	store := &Store{}
	err := store.Init(db.Node)
	return store, err
}

func TestAddEntry(t *testing.T) {
	ns, err := GenTestNodeStore()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < NumEntries; i++ {
		k, err := crypto.EdDSAGen()
		if err != nil {
			t.Fatal(err)
		}

		payload, err := k.Encrypt(TestPassword)
		if err != nil {
			t.Fatal(err)
		}

		id, err := ns.AddEntry(payload)
		if err != nil {
			t.Fatal(err)
		}
		Entries[i] = id
	}

	count := 0
	cursor, err := ns.ID.Find(ns.Ctx, bson.M{})
	if err != nil {
		t.Fatal(err)
	}
	defer cursor.Close(ns.Ctx)

	for cursor.Next(ns.Ctx) {
		count++
	}

	if count != NumEntries {
		t.Fatalf("Insufficient Count: %v != %v", count, NumEntries)
	}
}

func TestGetEntries(t *testing.T) {
	ns, err := GenTestNodeStore()
	if err != nil {
		t.Fatal(err)
	}

	entries, err := ns.GetEntries(TestPassword)
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != NumEntries {
		t.Fatalf("expected %v entries. got %v", NumEntries, len(entries))
	}

	if _, err = ns.GetEntries([]byte{}); err == nil {
		t.Fatal("expected decryption error")
	}
}

func TestRemoveEntry(t *testing.T) {
	ns, err := GenTestNodeStore()
	if err != nil {
		t.Fatal(err)
	}

	removed := 0
	for _, entry := range Entries {
		if err := ns.RemoveEntry(entry); err != nil {
			t.Fatal(err)
		}
		removed++
	}

	if removed != NumEntries {
		t.Fatalf("Insufficient Count: %v != %v", removed, NumEntries)
	}
}
