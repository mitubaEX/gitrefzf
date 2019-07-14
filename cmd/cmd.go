package cmd

// Execute cmd
func Execute() error {
	repos, err := findRepo()
	if err != nil {
		return err
	}

	idxs, err := fzf(repos)
	if err != nil {
		return err
	}

	err = open(repos, idxs)
	if err != nil {
		return err
	}

	return nil
}
