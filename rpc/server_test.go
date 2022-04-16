package rpc

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	ServerPort = 9000
	Clients    = (1 << 9)
)

var (
	Timeout = time.Duration(10 * time.Second)
)

func TestServer(t *testing.T) {
	serverAddr := fmt.Sprintf("0.0.0.0:%d", ServerPort)

	s := Server{}
	s.Init(ServerPort)
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
				errs <- ErrNoResponse
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
