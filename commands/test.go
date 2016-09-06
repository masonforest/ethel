package commands

import (
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func Test(c *cli.Context) error {
	cmd := exec.Command("/bin/sh", "-c", "go generate ./... && go test ./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
