package accounts

import (
	ctx "context"
	"fmt"
	"time"

	"github.com/georacle-labs/georacle/accounts/evm"
	"github.com/georacle-labs/georacle/chain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// ErrDelete is thrown on an unsuccessful account delete
	ErrDelete = errors.New("unsuccessful delete")

	// ErrAccount is thrown on an unsupported account type
	ErrAccount = errors.New("invalid account type")
)

// Store wraps a mongo connection to an account store
type Store struct {
	AccountType chain.Type        // the account type
	Ctx         ctx.Context       // calling context
	Accounts    *mongo.Collection // collection of encrypted accounts for this chain
}

// Key represents an encrypted member of the account store
type Key struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Raw       []byte             `bson:"raw"`
}

// Init the account store for this chain
func (s *Store) Init(db *mongo.Database) error {
	s.Ctx = ctx.Background()

	dbName := fmt.Sprintf("georacle_%s", s.AccountType)
	s.Accounts = db.Collection(dbName)

	// ensure unique account entries
	s.Accounts.Indexes().CreateOne(
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
	k := Key{Raw: raw, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	// throws error on duplicate key entry
	result, err := s.Accounts.InsertOne(s.Ctx, k)
	if err != nil {
		return id, err
	}

	id = result.InsertedID.(primitive.ObjectID)
	return id, err
}

// RemoveEntry removes an entry from the store by id
func (s *Store) RemoveEntry(id primitive.ObjectID) error {
	result, err := s.Accounts.DeleteOne(s.Ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return ErrDelete
	}

	return err
}

// GetEntries decrypts the account store
func (s *Store) GetEntries(password []byte) ([]Entry, error) {
	entries := make([]Entry, 0)

	cursor, err := s.Accounts.Find(s.Ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(s.Ctx)

	for cursor.Next(s.Ctx) {
		var key Key
		if err = cursor.Decode(&key); err != nil {
			return nil, err
		}
		account, err := s.DecryptAccount(password, key.Raw)
		if err != nil {
			// zero entries
			copy(entries, make([]Entry, len(entries)))
			return nil, err
		}
		entries = append(entries, Entry{key.ID, account})
	}

	return entries, nil
}

// DecryptAccount decrypts a single account conditioned on the account type
func (s *Store) DecryptAccount(password, encryptedAccount []byte) (Account, error) {
	switch s.AccountType {
	case chain.EVM:
		a := new(evm.Account)
		err := a.Import(encryptedAccount, password)
		return a, err
	default:
		return nil, errors.Wrapf(ErrAccount, " %s", s.AccountType)
	}
}
