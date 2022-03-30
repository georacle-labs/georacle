package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB encapsulates a mongo connection
type DB struct {
	Name     string // chain dependent
	Client   *mongo.Client
	Services *mongo.Collection
}

// Service represents a generic service type
type Service struct {
	ID          primitive.ObjectID `bson:" _id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"created_at"`
	QueryType   int                `bson:"query_type"`
	Query       []byte             `bson:"query"`
	QueryDigest []byte             `bson:"query_digest"`
}

// Open initiates a new mongo connection
func (d *DB) Open(uri string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	db := client.Database(d.Name)
	d.Services = db.Collection("services")
	d.Client = client

	return nil
}

// Close terminates an existing mongo connection
func (d *DB) Close() error {
	return d.Client.Disconnect(context.TODO())
}
