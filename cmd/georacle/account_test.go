package main

import (
	"os"
	"testing"

	"github.com/georacle-labs/georacle/client"
)

func TestNewClient(t *testing.T) {
	cfg, err := client.DefaultConfig()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Network) <= 0 {
		t.Fatal(cfg.Network)
	}
	if len(cfg.WSURI) <= 0 {
		t.Fatal(cfg.WSURI)
	}
	if len(cfg.DBURI) <= 0 {
		t.Fatal(cfg.DBURI)
	}
	if len(cfg.Addr) <= 0 {
		t.Fatal(cfg.Addr)
	}
	if cfg.Port <= 0 {
		t.Fatal(cfg.Port)
	}
}

func TestAccountNew(t *testing.T) {
	app, err := NewCLI()
	if err != nil {
		t.Fatal(err)
	}

	args := append(os.Args[0:1], "accounts", "new")
	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}

	args = append(os.Args[0:1], "accounts", "list")
	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}
}

func TestAccountRemove(t *testing.T) {
	app, err := NewCLI()
	if err != nil {
		t.Fatal(err)
	}

	args := append(os.Args[0:1], "accounts", "remove", "1")
	if err := app.Run(args); err == nil {
		t.Fatal("expected index out of range error")
	}

	args = append(os.Args[0:1], "accounts", "remove", "0")
	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}

	args = append(os.Args[0:1], "accounts", "remove", "0")
	if err := app.Run(args); err == nil {
		t.Fatal("expected index out of range error")
	}

	args = append(os.Args[0:1], "accounts", "list")
	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}
}
