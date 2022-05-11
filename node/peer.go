package node

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/georacle-labs/georacle/crypto"
	"github.com/georacle-labs/georacle/rpc"
	"github.com/pkg/errors"
)

const (
	// PortRange is the largest valid TCP port
	PortRange = (1 << 16) - 1
)

var (
	// ErrInvalidURI is thrown on a malformed URI string
	ErrInvalidURI = errors.New("InvalidURIError")
)

// Peer represents a connecting node
type Peer struct {
	Pubkey []byte     // 64 byte ed25519 pubkey
	Host   net.IP     // host ip
	Port   uint16     // tcp port
	client rpc.Client // rpc client
}

// Parse a URI string of the format:
// <64 byte hex-encoded pubkey>@<host>:<port>
func (p *Peer) Parse(id string) error {
	split := strings.Split(id, "@")
	if len(split) != 2 {
		return ErrInvalidURI
	}

	// 64 byte hex encoded ed25519 pubkey
	pk, err := hex.DecodeString(strings.Replace(split[0], "0x", "", 1))
	if err != nil || !crypto.ValidEdDSA(pk) {
		return ErrInvalidURI
	}

	split = strings.Split(split[1], ":")
	if len(split) <= 1 {
		return ErrInvalidURI
	}

	// TCP port
	port, err := strconv.Atoi(split[len(split)-1])
	if err != nil || port <= 0 || port > PortRange {
		return ErrInvalidURI
	}

	// Host IP must be one of the following
	// IPv4 ("192.0.2.1")
	// IPv6 ("2001:db8::68")
	// IPv4-mapped IPv6 ("::ffff:192.0.2.1")
	ip := net.ParseIP(strings.Join(split[:len(split)-1], ""))
	if ip == nil {
		return ErrInvalidURI
	}

	p.Pubkey = pk
	p.Host = ip
	p.Port = uint16(port)

	return nil
}

// Connected checks for an active connection
func (p *Peer) Connected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()
	up, err := p.client.Ping(ctx)
	return up && err == nil
}

// Connect to a peer over jsonrpc
func (p *Peer) Connect() error {
	return p.client.Init(p.Addr())
}

// Addr returns a peer's network address
func (p *Peer) Addr() string {
	return fmt.Sprintf("%s:%d", p.Host.String(), p.Port)
}

// ID returns a peer's ID
func (p *Peer) ID() string {
	return "0x" + hex.EncodeToString(p.Pubkey)
}
