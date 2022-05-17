package main

import (
	"errors"
	"strconv"

	"github.com/urfave/cli/v2"
)

var (
	//ErrAccount is thrown on an unspecified account
	ErrAccount = errors.New("specify an account")
)

// AccountList lists accounts to stdout
func AccountList(ctx *cli.Context) error {
	config, err := ParseConfig(ctx)
	if err != nil {
		return err
	}

	client, err := config.NewClient()
	if err != nil {
		return err
	}

	if err := client.ListAccounts(); err != nil {
		return err
	}

	return nil
}

// AccountNew generates a new account and updates the keystore
func AccountNew(ctx *cli.Context) error {
	config, err := ParseConfig(ctx)
	if err != nil {
		return err
	}

	client, err := config.NewClient()
	if err != nil {
		return err
	}

	if err := client.NewAccount(); err != nil {
		return err
	}

	return nil
}

// AccountRemove deletes an account from the keystore by id
func AccountRemove(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return ErrAccount
	}

	id, err := strconv.Atoi(ctx.Args().Get(0))
	if err != nil {
		return err
	}

	config, err := ParseConfig(ctx)
	if err != nil {
		return err
	}

	client, err := config.NewClient()
	if err != nil {
		return err
	}

	if err := client.RemoveAccount(uint(id)); err != nil {
		return err
	}

	return nil
}
