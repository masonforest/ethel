package main

import (
  "os"

  "github.com/urfave/cli"
  "github.com/masonforest/ethel/commands"
)

func main() {
  app := cli.NewApp()

  app.Commands = []cli.Command{
    {
      Name:    "new",
      Aliases: []string{"n"},
      Usage:   "create a new contract",
      Action:  commands.NewProject,
    },
  }

  app.Run(os.Args)
}
