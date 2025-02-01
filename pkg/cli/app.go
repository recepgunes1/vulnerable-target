package cli

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:            "vt",
		Usage:           "vulnerable target",
		HideHelpCommand: true,
		Flags:           GetFlags(),
	}
}
