package cli

import "github.com/urfave/cli/v2"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		getVersionFlag(),
		getListTemplatesFlag(),
		getVerbosityFlag(),
		getProviderFlag(),
		getIDFlag(),
	}
}
