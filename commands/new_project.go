package commands

import (
  "os"
  "path/filepath"
  "fmt"
  "github.com/urfave/cli"
)

func NewProject(c *cli.Context) error {
  fmt.Println("Created", c.Args().First())
  os.Mkdir("." + string(filepath.Separator) + c.Args().First(),0777);

  return nil
}
