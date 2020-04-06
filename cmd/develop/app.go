package develop

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/roderm/vuego-boilerplate/ricebox"
	log "github.com/sirupsen/logrus"
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
	},
}

func Run(ctx *cli.Context) error {
	npm, _ := context.WithCancel(ctx.Context)
	npmLogger := log.New()

	go func(ctx context.Context) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cmd := exec.CommandContext(ctx, "bash", "-c", "npm run dev")
		cmd.Dir = dir + "/web"
		cmd.Stdout = npmLogger.WriterLevel(log.InfoLevel)
		cmd.Stderr = npmLogger.WriterLevel(log.WarnLevel)
		panic(cmd.Run())
	}(npm)
	http.Handle("/", http.FileServer(ricebox.Static().HTTPBox()))
	address := fmt.Sprintf("%s:%d", ctx.String("host"), ctx.Int("port"))
	return http.ListenAndServe(address, nil)
}
