package main

import (
	"os"

	"github.com/urfave/cli"
	"log"
)

func main() {
	err := newApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()

	app.Name = "git-weed"
	app.Usage = "Manage git commit time."
	app.Version = VERSION
	app.Author = "yutaszk"
	app.Email = "yutaszk23@gmail.com"
	app.Flags = Flags
	app.Action = Action

	return app
}
