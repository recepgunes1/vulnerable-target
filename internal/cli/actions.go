package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var ValidLogLevels = map[string]bool{
	"info":  true,
	"warn":  true,
	"error": true,
	"fatal": true,
}

var ValidProviders = map[string]bool{
	"aws":          true,
	"azure":        true,
	"google-cloud": true,
	"docker":       true,
}

func getVersionFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:               "version",
		Aliases:            []string{"V"},
		Category:           "system",
		DisableDefaultText: true,
		Usage:              "show the current version of the tool",
		Action: func(ctx *cli.Context, b bool) error {
			if b {
				fmt.Println("version 1.0.0")
				os.Exit(0)
			}
			os.Exit(0)
			return nil
		},
	}
}

func getListTemplatesFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:               "list-templates",
		Aliases:            []string{"lt"},
		Category:           "system",
		DisableDefaultText: true,
		Usage:              "list all available templates with descriptions",
		Action: func(ctx *cli.Context, b bool) error {
			if b {
				log.Print("list-templates")
				os.Exit(0)
			}
			os.Exit(0)
			return nil
		},
	}
}

func getVerbosityFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "verbosity",
		Category: "logging",
		Aliases:  []string{"v"},
		Value:    "info",
		Usage:    "set the verbosity level for logs (info, warn, error, fatal)",
		Action: func(ctx *cli.Context, s string) error {
			if !ValidLogLevels[s] {
				log.Fatalf("invalid log level: %s. Must be one of: info, warn, error, fatal", s)
			}
			log.Printf("verbosity => %s \n", s)
			os.Exit(0)
			return nil
		},
	}
}

func getProviderFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "provider",
		Category: "target",
		Aliases:  []string{"p"},
		Usage:    "specify the target environment for vulnerability checks (aws, azure, google-cloud, docker)",
		Action: func(ctx *cli.Context, s string) error {
			if !ValidProviders[s] {
				log.Fatalf("invalid provider: %s. Must be one of: aws, azure, google-cloud, docker", s)
			}
			log.Printf("provider => %s \n", s)
			os.Exit(0)
			return nil
		},
	}
}

func getIDFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "id",
		Category: "target",
		Usage:    "specify a template ID for targeted operations",
		Action: func(ctx *cli.Context, s string) error {
			log.Printf("id => %s \n", s)
			os.Exit(0)
			return nil
		},
	}
}
