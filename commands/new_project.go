package commands

import (
  "fmt"
  "github.com/urfave/cli"
)

func NewProject(c *cli.Context) error {
  RestoreAssets(c.Args().First(), "")

  fmt.Println("Created", c.Args().First())
  return nil
}
