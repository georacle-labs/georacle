package accounts

import (
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Key represents an encrypted member of the keystore
type Key struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Raw       []byte             `bson:"raw"`
}

// AddEntry adds a key to the keystore and returns its id
func (m *Master) AddEntry(raw []byte) (primitive.ObjectID, error) {
	var id primitive.ObjectID
	k := Key{Raw: raw, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	// throws error on duplicate key entry
	result, err := m.Store.InsertOne(m.Ctx, k)
	if err != nil {
		return id, err
	}

	id = result.InsertedID.(primitive.ObjectID)
	return id, err
}

// RemoveEntry removes a key from the keystore by id
func (m *Master) RemoveEntry(id primitive.ObjectID) error {
	result, err := m.Store.DeleteOne(m.Ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return errors.New("DocumentDeleteError")
	}

	return err
}

// GetEntries decrypts the keystore
func (m *Master) GetEntries() ([]Entry, error) {
	entries := make([]Entry, 0)

	cursor, err := m.Store.Find(m.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(m.Ctx)

	for cursor.Next(m.Ctx) {
		var key Key
		if err = cursor.Decode(&key); err != nil {
			return nil, err
		}
		account, err := m.DecryptAccount(key.Raw)
		if err != nil {
			// zero entries
			copy(entries, make([]Entry, len(entries)))
			return nil, err
		}
		entries = append(entries, Entry{key.ID, account})
	}

	return entries, nil
}
