package node

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/georacle-labs/georacle/rpc"
)

// Peer wraps an RPC client used to connect to a peer
type Peer struct {
	Pubkey []byte     // public key
	Host   net.IP     // host ip
	Port   uint16     // tcp port
	client rpc.Client // rpc client
}

// NewPeer returns an empty peer
func NewPeer(pubkey, addr []byte) (Peer, error) {
	if len(addr) != NetAddrSize {
		return Peer{}, ErrInvalidHost
	}
	if !crypto.ValidEdDSA(pubkey) {
		return Peer{}, ErrInvalidPubkey
	}
	p := Peer{Pubkey: pubkey, Host: addr[:16]}
	p.Port = binary.BigEndian.Uint16(addr[16:NetAddrSize])
	return p, nil
}

// Connected checks for an active connection
func (p *Peer) Connected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()
	up, err := p.client.Ping(ctx)
	return up && err == nil
}

// Connect to a peer over JSON-RPC
func (p *Peer) Connect() error {
	return p.client.Init(p.Addr())
}

// Addr returns a peer's network address
func (p *Peer) Addr() string {
	return fmt.Sprintf("%s:%d", p.Host.String(), p.Port)
}

// ID returns a peer's ID
func (p *Peer) ID() string {
	return Encode(p.Pubkey, p.Host, p.Port)
}
