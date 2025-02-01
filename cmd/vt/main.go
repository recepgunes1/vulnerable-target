package main

import (
	"log"
	"os"

	"github.com/happyhackingspace/vulnerable-target/pkg/cli"
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
