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
	Client   *mongo.Client
	Services *mongo.Collection
}

// Service represents a Generic service type
type Service struct {
	ID          primitive.ObjectID `bson:" _id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"created_at"`
	QueryType   int                `bson:"query_type"`
	Query       []byte             `bson:"query"`
	QueryDigest []byte             `bson:"query_digest"`
}

// NewDB returns a new mongo connection
func NewDB(uri string) *DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil
	}

	db := client.Database("georacle")
	services := db.Collection("services")

	return &DB{Client: client, Services: services}
}

// Close terminates an existing mongo connection
func (d *DB) Close() error {
	return d.Client.Disconnect(context.TODO())
}
