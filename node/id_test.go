package node

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"net"
	"testing"

	"github.com/georacle-labs/georacle/crypto"
)

func TestEncode(t *testing.T) {
	randAddr := make([]byte, 16)
	randPort := make([]byte, 2)

	for i := 0; i < (1 << 10); i++ {
		keyPair, err := crypto.EdDSAGen()
		if err != nil {
			t.Fatal(err)
		}

		if _, err := io.ReadFull(rand.Reader, randAddr); err != nil {
			t.Fatal(err)
		}
		if _, err := io.ReadFull(rand.Reader, randPort); err != nil {
			t.Fatal(err)
		}

		port := binary.BigEndian.Uint16(randPort)
		ipv6 := net.IP(randAddr)
		ipv4 := net.IP(randAddr[:4])

		v6 := Encode(keyPair.Pub, ipv6, port)
		v4 := Encode(keyPair.Pub, ipv4, port)

		// check v6 encoding
		decodedPK, decodedIP, decodedPort, err := Decode(v6)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(decodedPK, keyPair.Pub) {
			t.Fatal("pubkey encoding failed")
		}
		if !decodedIP.Equal(ipv6) {
			t.Fatal("ipv6 encoding failed")
		}
		if decodedPort != port {
			t.Fatal("port encoding failed")
		}

		// check v4 encoding
		decodedPK, decodedIP, decodedPort, err = Decode(v4)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(decodedPK, keyPair.Pub) {
			t.Fatal("pubkey encoding failed")
		}
		if !decodedIP.Equal(ipv4) {
			t.Fatal("ipv4 encoding failed")
		}
		if decodedPort != port {
			t.Fatal("port encoding failed")
		}
	}
}
