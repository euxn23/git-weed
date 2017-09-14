package main

import (
	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)


func Action(_ *cli.Context) error {
	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	w.Commit(commitMessage, &git.CommitOptions{
		Committer: &object.Signature{
		},
		Author: &object.Signature{
		},
	})

	return nil
}
