package peer

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	mr "math/rand"
	"net"
	"strconv"
	"testing"

	"github.com/georacle-labs/georacle/crypto"
)

func TestURI(t *testing.T) {
	var p Peer
	v4 := make([]byte, 4)

	for i := 0; i < (1 << 10); i++ {

		if _, err := io.ReadFull(rand.Reader, v4); err != nil {
			t.Fatal(err)
		}
		host := net.IPv4(v4[0], v4[1], v4[2], v4[3]).String()
		port := strconv.Itoa(int((mr.Uint32() % PortRange) + 1))

		keyPair, err := crypto.EdDSAGen()
		if err != nil {
			t.Fatal(err)
		}
		pk := hex.EncodeToString(keyPair.Pub)

		// parse valid ipv4
		if err := p.Parse(pk + "@" + host + ":" + port); err != nil {
			t.Fatal(err)
		}

		// parse invalid host
		if err := p.Parse(pk + "@" + string(v4) + ":" + port); err == nil {
			t.Fatal("Invalid host")
		}

		// parse invalid port
		if err := p.Parse(pk + "@" + host + ":" + string(v4)); err == nil {
			t.Fatal("Invalid port")
		}

		// parse invalid id
		if err := p.Parse(hex.EncodeToString(keyPair.Priv) + "@" + host + ":" + port); err == nil {
			t.Fatal("Invalid id")
		}

		// parse invalid id separator
		if err := p.Parse(hex.EncodeToString(keyPair.Priv) + host + ":" + port); err == nil {
			t.Fatal("Invalid separator")
		}

		// parse invalid port separator
		if err := p.Parse(hex.EncodeToString(keyPair.Priv) + "@" + host + port); err == nil {
			t.Fatal("Invalid separator")
		}
	}
}
