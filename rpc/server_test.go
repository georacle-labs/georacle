package rpc

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

const (
	Addr    = "0.0.0.0"
	Port    = 4242
	Clients = (1 << 9)
)

func TestServer(t *testing.T) {
	serverAddr := fmt.Sprintf("%s:%d", Addr, Port)

	s := Server{}
	if err := s.Init(Addr, Port); err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	go s.Run(context.Background())

	var wg sync.WaitGroup
	errs := make(chan error, Clients)

	for i := 0; i < Clients; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := new(Client)
			if err := c.Init(serverAddr); err != nil {
				errs <- err
				return
			}
			defer c.Close()
			ack, err := c.Ping(context.Background())
			if !ack {
				errs <- ErrResponse
			}
			errs <- err
		}()
	}

	wg.Wait()

	close(errs)
	for e := range errs {
		if e != nil {
			t.Fatal(e)
		}
	}
}
