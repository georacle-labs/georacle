package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/client"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

// ParseConfig reads cli opts into a client config
func ParseConfig(ctx *cli.Context) (*client.Config, error) {
	config := new(client.Config)
	config.Network = ctx.String("network")
	config.Password = ctx.String("password")
	config.DBURI = ctx.String("db-uri")
	config.WSURI = ctx.String("ws-uri")

	// scan conf file for missing fields
	jsonConfig := new(client.Config)
	confFile := ctx.String("config")

	if err := jsonConfig.FromJSON(confFile); err == nil {
		if len(config.Network) <= 0 {
			config.Network = jsonConfig.Network
		}
		if len(config.Password) <= 0 {
			config.Password = jsonConfig.Password
		}
		if len(config.WSURI) <= 0 {
			config.WSURI = jsonConfig.WSURI
		}
		if len(config.DBURI) <= 0 {
			config.DBURI = jsonConfig.DBURI
		}
	}

	if len(config.Password) <= 0 {
		pass, err := GetPass()
		if err != nil {
			return nil, err
		}
		config.Password = pass
	} else {
		// zero password
		jsonConfig.Password = ""
	}

	return config, nil
}

// Start parses a cli config and starts the underlying client
func Start(ctx *cli.Context) error {
	config, err := ParseConfig(ctx)
	if err != nil {
		return err
	}

	client, err := config.NewClient()
	if err != nil {
		return err
	}

	return client.Start()
}

// NewCLI inits a new cli instance of the client
func NewCLI() (*cli.App, error) {
	app := &cli.App{Name: "georacle", Usage: "the georacle client"}
	app.EnableBashCompletion = true

	validNetworks := reflect.ValueOf(chain.Chains).MapKeys()
	defaultConfigPath, err := client.DefaultConfigPath()
	if err != nil {
		return nil, err
	}

	app.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "start the georacle node",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "network",
					Usage: fmt.Sprintf("valid networks: %v", validNetworks),
				},
				&cli.StringFlag{
					Name:  "ws-uri",
					Usage: "a websocket rpc endpoint",
				},
				&cli.StringFlag{
					Name:  "db-uri",
					Usage: "a mongodb connection string",
				},
				&cli.StringFlag{
					Name:    "password",
					Aliases: []string{"p"},
					Usage:   "master password for unlocking keystore",
				},
				&cli.StringFlag{
					Name:    "config",
					Value:   defaultConfigPath,
					Aliases: []string{"c"},
					Usage:   "Load json configuration from `FILE`",
				},
			},
			Action: Start,
		},
		{
			Name:  "accounts",
			Usage: "manage georacle accounts",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "network",
					Usage: fmt.Sprintf("valid networks: %v", validNetworks),
				},
				&cli.StringFlag{
					Name:  "ws-uri",
					Usage: "a websocket rpc endpoint",
				},
				&cli.StringFlag{
					Name:  "db-uri",
					Usage: "a mongodb connection string",
				},
				&cli.StringFlag{
					Name:    "password",
					Aliases: []string{"p"},
					Usage:   "master password for unlocking keystore",
				},
				&cli.StringFlag{
					Name:    "config",
					Value:   defaultConfigPath,
					Aliases: []string{"c"},
					Usage:   "Load json configuration from `FILE`",
				},
			},
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "list available accounts",
					Action:  AccountList,
				},
				{
					Name:    "new",
					Aliases: []string{"n"},
					Usage:   "generate a new account",
					Action:  AccountNew,
				},
				{
					Name:    "remove",
					Aliases: []string{"rm"},
					Usage:   "remove an existing account",
					Action:  AccountRemove,
				},
			},
		},
	}

	return app, nil
}

// GetPass prompts the user for a password without echoing
func GetPass() (password string, err error) {
	var bytes []byte

	fmt.Print("Enter Password: ")
	bytes, err = terminal.ReadPassword(0)
	if err == nil {
		password = strings.TrimSpace(string(bytes))
	}
	fmt.Printf("\n")

	return
}
