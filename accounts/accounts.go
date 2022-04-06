package accounts

import (
	ctx "context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/georacle-labs/georacle/accounts/evm"
	"github.com/georacle-labs/georacle/chain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Account represents a generic account
type Account interface {
	Gen() error
	Export(string) ([]byte, error)
	Import([]byte, string) error
	String() string
}

// Entry represents a keystore entry
type Entry struct {
	ID      primitive.ObjectID
	Account Account
}

// Master holds a chain's master key
// This Key:
//  * Is responsible for encrypting/decrypting a keystore
//  * Is user-provided and loaded into memory at runtime
//  * Should never touch disk
type Master struct {
	Type     chain.Type
	Password string
	Entries  []Entry
	Ctx      ctx.Context
	Store    *mongo.Collection
}

// Init the master keystore
func (m *Master) Init() error {
	log.Println("Decrypting keystore...")
	entries, err := m.GetEntries()
	if err != nil {
		return err
	}

	m.Entries = entries

	log.Printf("Decrypted %d account(s)\n", len(m.Entries))

	return nil
}

// NewAccount generates a new account conditioned on the chain type
func (m *Master) NewAccount() error {
	var account Account
	switch m.Type {
	case chain.EVM:
		account = new(evm.Account)
		if err := account.Gen(); err != nil {
			return err
		}
	default:
		return errors.Errorf("Invalid Account type: %v", m.Type)
	}

	if err := m.InsertAccount(account); err != nil {
		return err
	}

	log.Printf(
		"Generated new account of type %s: %s\n", m.Type.String(), account.String(),
	)

	return nil
}

// RemoveAccount removes an account from the keystore by index
func (m *Master) RemoveAccount(index uint) error {
	if int(index) >= len(m.Entries) {
		return errors.New("InvalidIndexError")
	}

	toRemove := m.Entries[index]
	if err := m.RemoveEntry(toRemove.ID); err != nil {
		return err
	}

	m.Entries = append(m.Entries[:index], m.Entries[index+1:]...)

	log.Printf("Removed account: %s", toRemove.Account.String())

	return nil
}

// InsertAccount inserts a unique account into the keystore
func (m *Master) InsertAccount(account Account) error {
	for _, e := range m.Entries {
		if e.Account.String() == account.String() {
			return errors.New("DuplicateAccountError")
		}
	}

	encrypted, err := account.Export(m.Password)
	if err != nil {
		return err
	}

	id, err := m.AddEntry(encrypted)
	if err != nil {
		return err
	}

	m.Entries = append(m.Entries, Entry{id, account})
	return nil
}

// DecryptAccount decrypts a single account conditioned on the account type
func (m *Master) DecryptAccount(encryptedAccount []byte) (Account, error) {
	switch m.Type {
	case chain.EVM:
		a := new(evm.Account)
		err := a.Import(encryptedAccount, m.Password)
		return a, err
	default:
		return nil, errors.Errorf("Invalid Account type: %v", m.Type)
	}
}

// DumpAccounts to stdout
func (m *Master) DumpAccounts() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t", "id", "Account")
	fmt.Fprintf(w, "\n %s\t%s\t", "--", "-------")

	for i, e := range m.Entries {
		hexString := e.Account.String()
		fmt.Fprintf(w, "\n %d\t%s\t", i, hexString)
	}

	fmt.Fprintf(w, "\n")
}
