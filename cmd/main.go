package main

import (
	"os"

	"github.com/roderm/vuego-boilerplate/cmd/develop"
	"github.com/roderm/vuego-boilerplate/cmd/serve"
	log "github.com/sirupsen/logrus"
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
	if err != nil {
		log.Fatal(err)
	}
}
