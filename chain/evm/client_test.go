package evm

import (
	"os"
	"sync"
	"testing"
	"time"
)

var (
	wsURI = os.Getenv("WS_URI")
)

func TestInvalidOpen(t *testing.T) {
	// empty URI should fail
	if c, err := NewClient(42, ""); err == nil {
		if err := c.Open(); err == nil {
			t.Fatal("expected failure on empty uri")
		}
	} else {
		t.Fatal(err)
	}

	// mismatched chain ID should fail
	if c, err := NewClient(0, wsURI); err == nil {
		if err := c.Open(); err == nil {
			t.Fatal("expected failure on mismatched URI")
		}
	} else {
		t.Fatal(err)
	}
}

func TestInvalidState(t *testing.T) {
	c, err := NewClient(42, wsURI)
	if err != nil {
		t.Fatal(c)
	}

	// first open should pass
	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	// subsequent open should fail
	if err := c.Open(); err != ErrClientState {
		t.Fatal("expected invalid state on open")
	}

	// first close should pass
	if err := c.Close(); err != nil {
		t.Fatal(err)
	}

	// subsequent close should fail
	if err := c.Close(); err == nil {
		t.Fatal("expected invalid state on close")
	}

	// open after close should pass
	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	// close after open should pass
	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestOpen(t *testing.T) {
	c, err := NewClient(42, wsURI)
	if err != nil {
		t.Fatal(c)
	}

	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestRun(t *testing.T) {
	c, err := NewClient(42, wsURI)
	if err != nil {
		t.Fatal(c)
	}

	if err := c.Open(); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = c.Run()
	}()

	// wait for run call
	time.Sleep(1 * time.Second)

	if err := c.Close(); err != nil {
		t.Fatal(err)
	}

	wg.Wait()
	if err != nil {
		t.Fatal(err)
	}
}
