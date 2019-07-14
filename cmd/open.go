package cmd

import (
	"errors"
	"os/exec"
	"runtime"

	"github.com/inabajunmr/treview/github/trending"
)

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = errors.New("unsupported platform")
	}
	return err
}

func open(repos []trending.Repository, idxs []int) error {
	for _, idx := range idxs {
		err := openbrowser(repos[idx].URL)
		if err != nil {
			return err
		}
	}
	return nil
}
