package node

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/georacle-labs/georacle/rpc"
	"go.mongodb.org/mongo-driver/mongo"
)

// Node facilitates external communication
// through authenticated channels
type Node struct {
	Port    uint16          // listening port
	KeyPair *crypto.KeyPair // Node identifier
	Server  rpc.Server      // RPC server
	Store   Store           // persistent store
	Ctx     context.Context // calling context
}

// Init a node and generate a host ID
func (n *Node) Init(db *mongo.Database, password []byte) error {
	if err := n.Store.Init(db); err != nil {
		return err
	}

	entries, err := n.Store.GetEntries(password)
	if err != nil {
		return err
	}

	if len(entries) <= 0 {
		// node not yet initialized -> generate a new identity
		log.Println("Generating new identity...")

		// generate a new ed25519 key pair
		n.KeyPair, err = crypto.EdDSAGen()
		if err != nil {
			return err
		}

		// encrypt and store the key for subsequent use
		bytes, err := n.KeyPair.Encrypt(password)
		if err != nil {
			return err
		}
		if _, err = n.Store.AddEntry(bytes); err != nil {
			return err
		}
	} else {
		// node previously initialized -> use existing ID
		n.KeyPair = &entries[0].Key
	}

	log.Printf("Initialized node: %s\n", n.ID())

	n.Ctx = context.Background()
	return n.Server.Init(n.Port)
}

// Start the underlying server and listen for connections
func (n *Node) Start() error {
	log.Printf("Started Node: %v\n", n.ID())
	return n.Server.Run(n.Ctx)
}

// Stop the node and terminate the server
func (n *Node) Stop() error {
	return n.Server.Close()
}

// ID is a node's hex-encoded 32 byte public key
func (n *Node) ID() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(n.KeyPair.Pub))
}
