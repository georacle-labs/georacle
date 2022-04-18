package node

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/georacle-labs/georacle/peer"
	"github.com/georacle-labs/georacle/rpc"
	"go.mongodb.org/mongo-driver/mongo"
)

// Node represents a host on the network
type Node struct {
	ID      string               // node ID
	Port    uint16               // listening port
	KeyPair *crypto.KeyPair      // Node identifier
	Server  rpc.Server           // RPC server
	Store   Store                // persistent store
	Ctx     context.Context      // calling context
	Peers   map[string]peer.Peer // connected peers
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

	n.Peers = make(map[string]peer.Peer, 0)
	n.Ctx = context.Background()
	n.ID = fmt.Sprintf("0x%s", hex.EncodeToString(n.KeyPair.Pub))

	log.Printf("Initialized node: %s\n", n.ID)
	if err := n.Server.Init(n.Port); err != nil {
		return err
	}

	// bootstrap network
	return n.Bootstrap()
}

// Start the underlying server and listen for connections
func (n *Node) Start() error {
	log.Printf("Started Node: %v\n", n.ID)
	return n.Server.Run(n.Ctx)
}

// Stop the node and terminate the server
func (n *Node) Stop() error {
	return n.Server.Close()
}
