package main

import (
	"os"

	"github.com/masonforest/ethel/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "create ethereum accounts that ethel can deploy from",
			Action:  commands.Init,
		},
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create a new contract",
			Action:  commands.NewProject,
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test a contract",
			Action:  commands.Test,
		},
		{
			Name:    "deploy",
			Aliases: []string{"deploy"},
			Usage:   "deploy [contract name] [params]",
			Action:  commands.Deploy,
		},
	}

	app.Run(os.Args)
}
