package node

import (
	"log"
	"os"
	"strings"

	"github.com/georacle-labs/georacle/peer"
)

var (
	// BootstrapNodes is a temporary placeholder for a set of bootstrap nodes
	BootstrapNodes = os.Getenv("BOOTSTRAP")
)

// Bootstrap connects to an initial set of peers
func (n *Node) Bootstrap() error {
	peers := strings.Split(BootstrapNodes, ",")
	if len(peers) <= 0 {
		log.Println("Failed to bootstrap network")
	}

	for _, p := range peers {
		var peer peer.Peer
		if err := peer.Parse(p); err != nil {
			return err
		}
		if err := n.Connect(peer); err != nil {
			return err
		}
	}

	return nil
}

// Connect to a peer once
func (n *Node) Connect(peer peer.Peer) error {
	if p, ok := n.Peers[peer.ID()]; ok {
		// reuse existing connection
		if p.Connected() {
			return nil
		}
		// stale connection...reconnect
		return p.Connect()
	}
	// new peer...connect
	err := peer.Connect()
	if err == nil {
		n.Peers[peer.ID()] = peer
		log.Printf("[+] Connected to peer %s\n", peer.ID())
	}
	return err
}
