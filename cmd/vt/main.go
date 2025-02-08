package main

import (
	"github.com/happyhackingspace/vulnerable-target/internal/cli"
	"github.com/happyhackingspace/vulnerable-target/internal/logger"
)

func init() {
	logger.Init()
}

func main() {
	cli.Execute()
}
