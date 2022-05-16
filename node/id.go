package node

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/georacle-labs/georacle/crypto"
)

const (
	// PayloadSize is 50 bytes
	// 32 byte pubkey + 18 byte NetAddr
	PayloadSize = crypto.PublicKeySize + NetAddrSize

	// NetAddrSize is 18 bytes
	// 16 byte ipv6 address + 2 byte port
	NetAddrSize = 18
)

var (
	// ErrNodeID is thrown on parsing an invalid node ID
	ErrNodeID = errors.New("invalid node id")

	// ErrPubkey is thrown on an invalid pubkey
	ErrPubkey = errors.New("invalid public key")
)

// Encode hex encodes a node ID into the format:
// (32 byte public key || 18 byte host)
func Encode(pubkey []byte, addr net.IP, port uint16) (id string) {
	if host := EncodeAddr(addr, port); host != nil {
		id = hex.EncodeToString(append(pubkey, host...))
	}
	return fmt.Sprintf("0x%s", id)
}

// EncodeAddr encodes a host into the format:
// (16 byte ipv6 addr || 2 byte port)
func EncodeAddr(addr net.IP, port uint16) []byte {
	addrv6 := addr.To16()
	if addrv6 == nil {
		return nil
	}

	// 18 byte host
	host := make([]byte, 18)

	// 16 byte network addr
	copy(host[0:16], addrv6)

	// 2 byte be encoded port
	binary.BigEndian.PutUint16(host[16:18], port)

	return host
}

// Decode a hex encoded node ID
func Decode(id string) (pubkey []byte, ip net.IP, port uint16, err error) {
	var payload []byte
	payload, err = hex.DecodeString(strings.Replace(id, "0x", "", 1))
	if err != nil {
		return
	}

	// 50 byte node ID
	if len(payload) != PayloadSize {
		err = ErrNodeID
		return
	}

	if !crypto.ValidEdDSA(payload[:crypto.PublicKeySize]) {
		err = ErrPubkey
		return
	}
	netAddr := payload[crypto.PublicKeySize:PayloadSize]

	// 32 byte public key
	pubkey = payload[:crypto.PublicKeySize]

	// 16 byte host
	ip = netAddr[:16]

	// 2 byte port
	port = binary.BigEndian.Uint16(netAddr[16:18])

	return
}
