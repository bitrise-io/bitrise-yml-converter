package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   ".",
			Action:  convert,
			Flags: []cli.Flag{
				flSource,
				flDestination,
			},
		},
	}
)
