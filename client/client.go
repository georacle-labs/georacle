package client

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/georacle-labs/go-georacle/chain"
	"github.com/georacle-labs/go-georacle/db"
	"github.com/pkg/errors"
)

// Config represents the complete configuration necessary to run a client
type Config struct {
	Network string `json:"network"` // network type
	WSURI   string `json:"ws_uri"`  // node websocket uri
	DBURI   string `json:"db_uri"`  // mongo connection string
}

// Client represents a single client instance
type Client struct {
	Params *chain.Config   // chain params
	Wg     *sync.WaitGroup // client's running go routines
	Chain  chain.Chain     // corresponding node connection
	DB     *db.DB          // open db handle
}

// NewClient returns a new client instance from an existing config
func (c *Config) NewClient() (*Client, error) {
	chainParams, err := c.ParseNetwork()
	if err != nil {
		return nil, err
	}

	node, err := chain.New(chainParams, c.WSURI)
	if err != nil {
		return nil, err
	}

	db, err := c.ParseDB()
	if err != nil {
		return nil, err
	}

	return &Client{chainParams, new(sync.WaitGroup), node, db}, nil
}

// ParseNetwork validates a config's `network` field
func (c *Config) ParseNetwork() (*chain.Config, error) {
	params, ok := chain.Chains[c.Network]
	if !ok {
		return nil, errors.Errorf("InvalidNetworkError: `%v` is not a valid network\n", c.Network)
	}
	return params, nil
}

// ParseDB validates a config's `db_uri` field
func (c *Config) ParseDB() (*db.DB, error) {
	// database name is the network name
	db := &db.DB{Name: c.Network}
	if err := db.Open(c.DBURI); err != nil {
		return nil, err
	}
	return db, nil
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

// FromJSON populates a config from a json file
func (c *Config) FromJSON(confPath string) error {
	file, err := ioutil.ReadFile(confPath)
	if err != nil {
		return errors.Errorf("Error parsing config file: %v", confPath)
	}

	if err = json.Unmarshal([]byte(file), c); err != nil {
		return err
	}

	return nil
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
