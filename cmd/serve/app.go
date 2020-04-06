package serve

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

var Runner = &cli.Command{
	Action: Run,
	Name:   "serve",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			DefaultText: "Port on with the HTTP-Server is listening",
			Value:       3000,
		},
		&cli.StringFlag{
			Name:  "host",
			Value: "localhost",
		},
		&cli.StringFlag{
			Name:        "config",
			DefaultText: "Pass a config file",
		},
	},
}

func Run(ctx *cli.Context) error {
	// http.Handle("/", http.FileServer(static.Box().HTTPBox()))
	address := fmt.Sprintf("%s:%d", ctx.String("host"), ctx.Int("port"))
	return http.ListenAndServe(address, nil)
}
