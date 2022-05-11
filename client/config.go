package client

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"reflect"

	"github.com/georacle-labs/georacle/accounts"
	"github.com/georacle-labs/georacle/chain"
	"github.com/georacle-labs/georacle/db"
	"github.com/georacle-labs/georacle/node"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidConf is thrown on an invalid config file
	ErrInvalidConf = errors.New("Error parsing config file")

	// ErrInvalidNetwork is thrown on an unsupported network type
	ErrInvalidNetwork = errors.New("Invalid network")
)

// Config represents the complete configuration necessary to run a node
type Config struct {
	Network  string `json:"network"`  // network type
	Password string `json:"password"` // master password (user prompted if empty)
	WSURI    string `json:"ws_uri"`   // chain-specific client websocket uri
	DBURI    string `json:"db_uri"`   // mongo connection string
	Addr     string `json:"addr"`     // public listening address
	Port     uint16 `json:"port"`     // public listening port
}

// NewClient returns a new client instance from an existing config
func (c *Config) NewClient() (*Client, error) {
	chainParams, err := c.ParseNetwork()
	if err != nil {
		return nil, err
	}

	db, err := c.ParseDB()
	if err != nil {
		return nil, err
	}

	chain, err := chain.New(chainParams, c.WSURI)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Params: chainParams,
		Chain:  chain,
		DB:     db,
	}

	client.Accounts = accounts.Master{
		Type:     chainParams.Type,
		Password: []byte(c.Password),
	}

	client.Node = node.Node{Addr: c.Addr, Port: c.Port}

	return client, nil
}

// ParseNetwork validates a config's `network` field
func (c *Config) ParseNetwork() (*chain.Config, error) {
	if err := ValidNetwork(c.Network); err != nil {
		return nil, err
	}
	return chain.Chains[c.Network], nil
}

// ParseDB validates a config's `db_uri` field
func (c *Config) ParseDB() (*db.DB, error) {
	// database name defaults to the network name
	db := &db.DB{Name: c.Network}
	if err := db.Open(c.DBURI); err != nil {
		return nil, err
	}
	return db, nil
}

// FromJSON populates a config from a json file
func (c *Config) FromJSON(confPath string) error {
	file, err := ioutil.ReadFile(confPath)
	if err != nil {
		return errors.Wrapf(ErrInvalidConf, confPath)
	}

	if err = json.Unmarshal([]byte(file), c); err != nil {
		return err
	}

	return nil
}

// DefaultConfig is a helper function to parse and return the default conf
func DefaultConfig() (*Config, error) {
	cfgPath, err := DefaultConfigPath()
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = cfg.FromJSON(cfgPath)
	return cfg, err
}

// DefaultConfigPath resolves to $HOME/.georacle/config.json
func DefaultConfigPath() (conf string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	conf = path.Join(home, ".georacle", "config.json")
	return
}

// ValidNetwork checks for a valid network
func ValidNetwork(network string) error {
	validNetworks := reflect.ValueOf(chain.Chains).MapKeys()
	for net := range chain.Chains {
		if net == network {
			return nil
		}
	}
	return errors.Wrapf(
		ErrInvalidNetwork, " `%s` must be one of %v\n", network, validNetworks,
	)
}
