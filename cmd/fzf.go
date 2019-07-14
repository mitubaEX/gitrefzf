package cmd

import (
	"fmt"

	"github.com/inabajunmr/treview/github/trending"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

func fzf(repos []trending.Repository) ([]int, error) {
	idxs, err := fuzzyfinder.FindMulti(
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
		return nil, err
	}

	return idxs, nil
}
