package commands

import (
  "os/exec"
  "os"
  "github.com/urfave/cli"
)

func Test(c *cli.Context) error {
  cmd := exec.Command("/bin/sh", "-c", "go generate ./... && go test ./...")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()

  return nil
}
