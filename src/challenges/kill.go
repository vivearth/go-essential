package challenge

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func KillServer(pidFile string) error {

	fdata, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer fdata.Close()
	var pid int

	if _, err := fmt.Fscanf(fdata, "%d", &pid); err != nil {
		return errors.Wrap(err, "bad process id")
	}
	fmt.Printf("Killed server with pid %d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warning: cannot remove pid file - %s", err)
	}
	return nil
}
