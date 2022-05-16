package node

import (
	"bytes"
	"testing"

	"github.com/georacle-labs/georacle/db"
)

func TestNewNode(t *testing.T) {
	if _, err := New("0.0.0.0", 9000); err != nil {
		t.Fatal(err)
	}
	if _, err := New("invalid host", 9000); err == nil {
		t.Fatal("expected invalid host")
	}
}

func TestNodeInit(t *testing.T) {
	db := &db.DB{Name: "test"}
	if err := db.Open(dbURI); err != nil {
		t.Fatal(err)
	}

	node, err := New("0.0.0.0", 9000)
	if err != nil {
		t.Fatal(err)
	}

	if err := node.Init(db.Node, TestPassword); err != nil {
		t.Fatal(err)
	}

	nodeEntries, err := node.Store.GetEntries(TestPassword)
	if err != nil {
		t.Fatal(err)
	}

	if len(nodeEntries) < 1 {
		t.Fatal("expected id generation on init")
	}

	if len(node.Hex()) <= 0 {
		t.Fatal("expected valid node ID")
	}

	if err := node.Store.RemoveEntry(nodeEntries[0].ID); err != nil {
		t.Fatal(err)
	}
}

func TestNodeConnect(t *testing.T) {
	db1 := &db.DB{Name: "node1"}
	if err := db1.Open(dbURI); err != nil {
		t.Fatal(err)
	}
	db2 := &db.DB{Name: "node2"}
	if err := db2.Open(dbURI); err != nil {
		t.Fatal(err)
	}

	n1, err := New("0.0.0.0", 9000)
	if err != nil {
		t.Fatal(err)
	}
	n2, err := New("0.0.0.0", 9001)
	if err != nil {
		t.Fatal(err)
	}

	if err := n1.Init(db1.Node, TestPassword); err != nil {
		t.Fatal(err)
	}
	if err := n2.Init(db2.Node, TestPassword); err != nil {
		t.Fatal(err)
	}

	go n1.Start()
	go n2.Start()
	if bytes.Equal(n1.Host(), n2.Host()) {
		t.Fatal("expected different hosts")
	}

	p1, err := NewPeer(n1.KeyPair.Pub, n1.Host())
	if err != nil {
		t.Fatal(err)
	}
	p2, err := NewPeer(n2.KeyPair.Pub, n2.Host())
	if err != nil {
		t.Fatal(err)
	}

	if err := n1.Connect(p2); err != nil {
		t.Fatal(err)
	}
	if err := n2.Connect(p1); err != nil {
		t.Fatal(err)
	}

	if err := p1.Connect(); err != nil {
		t.Fatal(err)
	}
	if err := p2.Connect(); err != nil {
		t.Fatal(err)
	}

	if !p1.Connected() || !p2.Connected() {
		t.Fatal("expected reachable peers")
	}

	n1.Stop()
	n2.Stop()
}
