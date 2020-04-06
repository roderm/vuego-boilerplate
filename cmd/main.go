package main

import (
	"os"

	"github.com/roderm/vuego-boilerplate/cmd/develop"
	"github.com/roderm/vuego-boilerplate/cmd/serve"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Commands: []*cli.Command{
			develop.Runner,
			serve.Runner,
		},
	}
	err := app.Run(os.Args)
	panic(err)
}
