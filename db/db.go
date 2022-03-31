package db

import (
	ctx "context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB encapsulates a mongo connection
type DB struct {
	Name     string            // the name of the db (typically matches network name)
	Ctx      ctx.Context       // calling context
	Client   *mongo.Client     // persistent mongo connection
	Services *mongo.Collection // collection of spatial queries
	Keys     *mongo.Collection // collection of encrypted keys
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
	db := client.Database(d.Name)
	d.Services = db.Collection("services")
	d.Keys = db.Collection("keys")

	// ensure unique key entries
	d.Keys.Indexes().CreateOne(
		d.Ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "raw", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	return nil
}

// Close terminates an existing mongo connection
func (d *DB) Close() error {
	return d.Client.Disconnect(d.Ctx)
}
