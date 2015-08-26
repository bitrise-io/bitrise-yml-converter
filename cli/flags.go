package cli

import "github.com/codegangsta/cli"

const (
	// PathKey ...
	PathKey      = "path"
	pathKeyShort = "p"

	// SourceKey ...
	SourceKey      = "source"
	sourceKeyShort = "s"

	// DestinationKey ...
	DestinationKey      = "destination"
	destinationKeyShort = "d"
)

var (
	flags = []cli.Flag{}

	flPath = cli.StringFlag{
		Name:  PathKey + ", " + pathKeyShort,
		Usage: ".",
	}

	flSource = cli.StringFlag{
		Name:  SourceKey + ", " + sourceKeyShort,
		Usage: ".",
	}

	flDestination = cli.StringFlag{
		Name:  DestinationKey + ", " + destinationKeyShort,
		Usage: ".",
	}
)
