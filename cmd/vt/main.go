package main

import (
	"github.com/happyhackingspace/vulnerable-target/internal/cli"
	"github.com/happyhackingspace/vulnerable-target/internal/config"
	"github.com/happyhackingspace/vulnerable-target/internal/logger"
	"github.com/happyhackingspace/vulnerable-target/pkg/providers"
	"github.com/happyhackingspace/vulnerable-target/pkg/templates"
)

func init() {
	logger.Init()
	templates.Init()
	config.LoadEnv()
}

func main() {
	cli.Execute()
	providers.Start()
}
