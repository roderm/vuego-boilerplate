package web

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const ServiceName = "Node Dev"

type Server struct {
	cliCtx *cli.Context
	logger *log.Logger
}

func New(ctx *cli.Context) *Server {
	return &Server{
		cliCtx: ctx,
		logger: log.New(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(ctx, "bash", "-c", "npm run dev")
	cmd.Dir = dir + "/web"
	if s.cliCtx.Bool("web-log") {
		cmd.Stdout = s.logger.WriterLevel(log.InfoLevel)
		cmd.Stderr = s.logger.WriterLevel(log.WarnLevel)
	}
	err = cmd.Run()
	return err
}
