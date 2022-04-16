package db

import (
	ctx "context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// AccountStore is the default name of the account store database
	AccountStore = "georacle_account_store"

	// NodeStore is the default name of the node ID database
	NodeStore = "georacle_node_store"
)

// DB encapsulates a mongo connection for persistent storage
type DB struct {
	Name     string          // the name of the db (typically matches network name)
	Ctx      ctx.Context     // calling context
	Client   *mongo.Client   // mongo connection
	Accounts *mongo.Database // client account store
	Node     *mongo.Database // node identifiers
}

// Service represents a generic service type
type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
	QueryType   int                `bson:"query_type"`
	Query       []byte             `bson:"query"`
	QueryDigest []byte             `bson:"query_digest"`
}

// Open initiates a new mongo connection
func (d *DB) Open(uri string) error {
	d.Ctx = ctx.Background()
	client, err := mongo.Connect(d.Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	d.Client = client
	d.Accounts = client.Database(AccountStore)
	d.Node = client.Database(NodeStore)

	return nil
}

// Close terminates an existing mongo connection
func (d *DB) Close() error {
	return d.Client.Disconnect(d.Ctx)
}
