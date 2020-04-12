package develop

import (
	net_http "net/http"

	"github.com/roderm/vuego-boilerplate/pkg/app"
	"github.com/roderm/vuego-boilerplate/pkg/grpc"
	"github.com/roderm/vuego-boilerplate/pkg/http"
	"github.com/roderm/vuego-boilerplate/pkg/web"
	"github.com/roderm/vuego-boilerplate/ricebox"
	"github.com/urfave/cli/v2"
)

var Runner = &cli.Command{
	Action: Run,
	Name:   "develop",
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
		&cli.BoolFlag{
			Name:        "web-log",
			Value:       true,
			DefaultText: "Writing the 'npm run dev' log into stdout",
		},
	},
}

func Run(ctx *cli.Context) error {
	a := app.New()
	http_svc := http.New(ctx)
	a.AddService(http.ServiceName, http_svc)
	a.AddService(grpc.ServiceName, grpc.New(ctx))
	a.AddService(web.ServiceName, web.New(ctx))
	http_svc.GetRouter().Handle("/static/",
		net_http.FileServer(ricebox.Static().HTTPBox()))
	return a.Run()
}
