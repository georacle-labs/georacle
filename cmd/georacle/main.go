package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/client"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cli, err := NewCLI()
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// ParseConfig reads cli flags into a client configuration
func ParseConfig(ctx *cli.Context) (*client.Config, error) {
	config := new(client.Config)
	config.Network = ctx.String("network")
	config.DBURI = ctx.String("db-uri")
	config.WSURI = ctx.String("ws-uri")
	config.Addr = ctx.String("addr")
	config.Port = uint16(ctx.Uint("port"))
	config.Password = ctx.String("password")

	// scan conf file for missing fields
	jsonConfig := new(client.Config)
	confFile := ctx.String("config")

	if err := jsonConfig.FromJSON(confFile); err == nil {
		config.Password = jsonConfig.Password
		if len(config.Network) <= 0 {
			config.Network = jsonConfig.Network
		}
		if len(config.WSURI) <= 0 {
			config.WSURI = jsonConfig.WSURI
		}
		if len(config.DBURI) <= 0 {
			config.DBURI = jsonConfig.DBURI
		}
		if len(config.Addr) <= 0 {
			config.Addr = jsonConfig.Addr
		}
		if config.Port <= 0 {
			config.Port = jsonConfig.Port
		}
		if len(config.Password) <= 0 {
			config.Password = jsonConfig.Password
		}
	}

	if len(config.Password) <= 0 {
		pass, err := GetPass()
		if err != nil {
			return nil, err
		}
		config.Password = pass
	}

	return config, nil
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
					Name:    "network",
					Aliases: []string{"n"},
					Usage:   fmt.Sprintf("valid networks: %v", validNetworks),
				},
				&cli.StringFlag{
					Name:    "ws-uri",
					Aliases: []string{"w"},
					Usage:   "a websocket rpc endpoint",
				},
				&cli.StringFlag{
					Name:    "db-uri",
					Aliases: []string{"d"},
					Usage:   "a mongodb connection string",
				},
				&cli.StringFlag{
					Name:    "addr",
					Aliases: []string{"a"},
					Usage:   "public listening address",
				},
				&cli.StringFlag{
					Name:    "port",
					Aliases: []string{"p"},
					Usage:   "listening port",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "data store decryption key",
				},
				&cli.StringFlag{
					Name:    "config",
					Value:   defaultConfigPath,
					Aliases: []string{"c"},
					Usage:   "load json configuration from `FILE`",
				},
			},
			Action: Start,
		},
		{
			Name:  "accounts",
			Usage: "manage georacle accounts",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "network",
					Aliases: []string{"n"},
					Usage:   fmt.Sprintf("valid networks: %v", validNetworks),
				},
				&cli.StringFlag{
					Name:    "db-uri",
					Aliases: []string{"d"},
					Usage:   "a mongodb connection string",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "data store decryption key",
				},
				&cli.StringFlag{
					Name:    "config",
					Value:   defaultConfigPath,
					Aliases: []string{"c"},
					Usage:   "load json configuration from `FILE`",
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
