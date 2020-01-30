package hexo

import (
	"log"
	"os/exec"
)

func clean() error {
	cmd := exec.Command("hexo", "clean")

	err := cmd.Run()
	if err != nil {
		log.Println("Execute Command failed:" + err.Error())
		return err
	}

	log.Println("Execute git pull finished.")
	return nil
}
