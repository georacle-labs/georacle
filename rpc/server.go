package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Server wraps a JSON RPC server
type Server struct {
	listener net.Listener
}

// Init an RPC server
func (s *Server) Init(port uint16) (err error) {
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		for _, service := range Services {
			rpc.Register(service)
		}
	}
	return
}

// Run the server's main loop
func (s *Server) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			con, err := s.listener.Accept()
			if err != nil {
				return err
			}

			remotePeer := con.RemoteAddr().String()
			log.Printf("[+] Received connection from %v\n", remotePeer)

			go jsonrpc.ServeConn(con)
		}
	}
}

// Close an RPC connection
func (s *Server) Close() error {
	return s.listener.Close()
}
