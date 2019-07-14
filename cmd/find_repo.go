package cmd

import "github.com/inabajunmr/treview/github/trending"

func findRepo() ([]trending.Repository, error) {
	span := trending.GetSpanByString("today")

	repos, err := trending.FindTrending("", span)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
