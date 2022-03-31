package cmd

import (
	"strconv"
	"syscall"

	"github.com/urfave/cli/v2"
)

// AccountList lists accounts to stdout
func AccountList(ctx *cli.Context) error {
	config, err := ParseConfig(ctx)
	if err != nil {
		return err
	}

	client, err := config.NewClient()
	if err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	if err := client.ListAccounts(); err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	return nil
}

// AccountNew generates a new account and updates the keystore
func AccountNew(ctx *cli.Context) error {
	config, err := ParseConfig(ctx)
	if err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	client, err := config.NewClient()
	if err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	if err := client.NewAccount(); err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	return nil
}

// AccountRemove deletes an account from the keystore by id
func AccountRemove(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return cli.Exit("specify an account", int(syscall.EINVAL))
	}

	id, err := strconv.Atoi(ctx.Args().Get(0))
	if err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	config, err := ParseConfig(ctx)
	if err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	client, err := config.NewClient()
	if err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	if err := client.RemoveAccount(uint(id)); err != nil {
		return cli.Exit(err.Error(), int(syscall.EINVAL))
	}

	return nil
}
