package node

import (
	"context"
	"log"
	"net"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/georacle-labs/georacle/rpc"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// ErrHost is thrown on an invalid listening address
	ErrHost = errors.New("invalid host")

	// ErrDisconnect is thrown when disconnecting from an invalid peer
	ErrDisconnect = errors.New("already disconnected")
)

// Node represents a host on the network
type Node struct {
	ID          string          // unique node ID
	Addr        net.IP          // public listening addr
	Port        uint16          // listening port
	KeyPair     *crypto.KeyPair // Node identifier
	Server      rpc.Server      // RPC server
	Store       Store           // persistent store
	Ctx         context.Context // calling context
	Peers       map[string]Peer // connected peers
	PeerUpdates chan Peer       // provider manager
}

// New returns an empty node
func New(host string, port uint16) (*Node, error) {
	res, err := net.LookupIP(host)
	if err != nil {
		return nil, errors.Wrapf(ErrHost, err.Error())
	}
	if len(res) > 0 {
		return &Node{Addr: res[0], Port: port}, nil
	}
	return nil, ErrHost
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

	if len(entries) < 1 {
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

	n.Peers = make(map[string]Peer, 0)
	n.PeerUpdates = make(chan Peer, 0)
	n.Ctx = context.Background()
	n.ID = n.Hex()

	return n.Server.Init(n.Addr.String(), n.Port)
}

// Start the rpc server and listen for connections
func (n *Node) Start() error {
	log.Printf("Started P2P networking with ID %v\n", n.ID)

	// start managing peers
	go func() {
		for {
			select {
			case p := <-n.PeerUpdates:
				var err error
				peerID := p.ID()
				if peer, connected := n.Peers[peerID]; connected {
					err = n.Disconnect(peer)
				} else {
					err = n.Connect(p)
				}
				if err != nil {
					log.Println("[Node] failed to connect to peer: ", peerID)
				}
			case <-n.Ctx.Done():
				return
			}
		}
	}()

	return n.Server.Run(n.Ctx)
}

// Stop the node and terminate the rpc server
func (n *Node) Stop() error {
	return n.Server.Close()
}

// Connect to a peer once
func (n *Node) Connect(p Peer) (err error) {
	// avoid self connection
	if n.Addr.Equal(p.Host) && n.Port == p.Port {
		return
	}

	// reuse existing connection
	if p, ok := n.Peers[p.ID()]; ok {
		if p.Connected() {
			return
		}
		// stale connection
		return p.Connect()
	}

	// new connection
	if err = p.Connect(); err == nil {
		n.Peers[p.ID()] = p
		log.Printf("[+] Connected to peer %s\n", p.ID())
	}

	return
}

// Disconnect from a peer if connected
func (n *Node) Disconnect(p Peer) error {
	peerID := p.ID()
	if peer, connected := n.Peers[peerID]; connected {
		if err := peer.client.Close(); err != nil {
			return err
		}
		delete(n.Peers, peerID)
		return nil
	}
	return ErrDisconnect
}

// Hex encode a node ID
func (n *Node) Hex() string {
	return Encode(n.KeyPair.Pub, n.Addr, n.Port)
}

// Host returns the node's public address
func (n *Node) Host() []byte {
	return EncodeAddr(n.Addr, n.Port)
}
