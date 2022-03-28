package db

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbURI = os.Getenv("DB_URI")

func TestValidConnection(t *testing.T) {

	db := NewDB(dbURI)
	if db == nil {
		t.Fatal(db)
	}

	defer db.Close()

	err := db.Client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}
}
