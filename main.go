package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/inabajunmr/treview/github/trending"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	span := trending.GetSpanByString("today")

	repos, err := trending.FindTrending("", span)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	idx, err := fuzzyfinder.FindMulti(
		repos,
		func(i int) string {
			return repos[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("%s",
				repos[i].ToString())
		}))
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idx {
		openbrowser(repos[i].URL)
	}
}
