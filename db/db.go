package db

import (
	ctx "context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	accountStoreName := fmt.Sprintf("georacle_%s_account_store", d.Name)
	d.Accounts = client.Database(accountStoreName)

	nodeStoreName := fmt.Sprintf("georacle_%s_node_store", d.Name)
	d.Node = client.Database(nodeStoreName)

	return nil
}

// Close terminates an existing mongo connection
func (d *DB) Close() error {
	return d.Client.Disconnect(d.Ctx)
}
