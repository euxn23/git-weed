package main

import (
	"github.com/motemen/go-gitconfig"
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

	userName, err := gitconfig.Default.GetString("user.name")
	if err != nil {
		return err
	}
	userEmail, err := gitconfig.Default.GetString("user.email")
	if err != nil {
		return err
	}

	w.Commit(commitMessage, &git.CommitOptions{
		Committer: &object.Signature{
			Name:  userName,
			Email: userEmail,
		},
		Author: &object.Signature{
			Name:  userName,
			Email: userEmail,
		},
	})

	return nil
}
