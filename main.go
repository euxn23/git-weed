package main

import (
	"os"

	"github.com/urfave/cli"
)


func main() {
	newApp().Run(os.Args)
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
