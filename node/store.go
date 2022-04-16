package node

import (
	ctx "context"
	"time"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store wraps a mongo connection to a node store
type Store struct {
	Ctx ctx.Context       // calling context
	ID  *mongo.Collection // collection of encrypted node identifiers
}

// Entry represents a node store entry
type Entry struct {
	ID  primitive.ObjectID
	Key crypto.KeyPair
}

// ID represents an encrypted node ID
type ID struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Raw       []byte             `bson:"raw"`
}

// Init the node store
func (s *Store) Init(db *mongo.Database) error {
	s.Ctx = ctx.Background()

	s.ID = db.Collection("id")

	// ensure unique id entries
	s.ID.Indexes().CreateOne(
		s.Ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "raw", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	return nil
}

// AddEntry adds an entry to the store and returns its id
func (s *Store) AddEntry(raw []byte) (primitive.ObjectID, error) {
	var id primitive.ObjectID
	i := ID{Raw: raw, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	// throws error on duplicate key entry
	result, err := s.ID.InsertOne(s.Ctx, i)
	if err != nil {
		return id, err
	}

	id = result.InsertedID.(primitive.ObjectID)
	return id, err
}

// RemoveEntry removes an entry from the store by id
func (s *Store) RemoveEntry(id primitive.ObjectID) error {
	result, err := s.ID.DeleteOne(s.Ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("DocumentDeleteError")
	}

	return err
}

// GetEntries decrypts the node store
func (s *Store) GetEntries(password []byte) ([]Entry, error) {
	entries := make([]Entry, 0)

	cursor, err := s.ID.Find(s.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(s.Ctx)

	var id ID
	var k crypto.KeyPair

	for cursor.Next(s.Ctx) {
		if err = cursor.Decode(&id); err != nil {
			return nil, err
		}
		if err := k.Decrypt(id.Raw, password); err != nil {
			// zero entries
			copy(entries, make([]Entry, len(entries)))
			return nil, err
		}
		entries = append(entries, Entry{id.ID, k})
	}

	return entries, nil
}
