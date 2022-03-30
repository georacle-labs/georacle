package cmd

import (
	"fmt"
	"reflect"

	"github.com/georacle-labs/go-georacle/chain"
	"github.com/georacle-labs/go-georacle/client"
	"github.com/urfave/cli/v2"
)

// Start parses a cli config and starts the underlying client
func Start(ctx *cli.Context) error {
	c := new(client.Config)
	c.Network = ctx.String("network")
	c.DBURI = ctx.String("db-uri")
	c.WSURI = ctx.String("ws-uri")

	// scan conf file for missing fields
	jsonConfig := new(client.Config)
	confFile := ctx.String("config")
	if err := jsonConfig.FromJSON(confFile); err == nil {
		if len(c.Network) <= 0 {
			c.Network = jsonConfig.Network
		}
		if len(c.WSURI) <= 0 {
			c.WSURI = jsonConfig.WSURI
		}
		if len(c.DBURI) <= 0 {
			c.DBURI = jsonConfig.DBURI
		}
	}

	client, err := c.NewClient()
	if err != nil {
		return err
	}

	return client.Chain.Run()
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
					Name:    "config",
					Value:   defaultConfigPath,
					Aliases: []string{"c"},
					Usage:   "Load json configuration from `FILE`",
				},
			},
			Action: Start,
		},
	}

	return app, nil
}
