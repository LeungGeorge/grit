package update

import (
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

// Update 更新 grit
var Update = cli.Command{
	Name:   "update",
	Usage:  "upgrade to newest version of master.",
	Action: update,
}

// update executes update command and return exit code.
func update(ctx *cli.Context) error {
	// go get -u github.com/LeungGeorge/grit
	cmd := exec.Command("go", "get", "-u", "github.com/LeungGeorge/grit")

	err := cmd.Run()
	if err != nil {
		log.Println("Execute Command failed:" + err.Error())
		return err
	}

	log.Println("Execute grit update finished.")
	return nil
}
