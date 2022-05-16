package cmd

import (
	"os"
	"testing"
)

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
