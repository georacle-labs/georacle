package main

import (
	"log"
	"os"

	"github.com/georacle-labs/georacle/cmd"
)

func main() {
	cli, err := cmd.NewCLI()
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
