package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

const ServiceName = "HTTP"

type Server struct {
	cliCtx *cli.Context
	hndlr  *mux.Router
}

func New(ctx *cli.Context) *Server {
	return &Server{
		cliCtx: ctx,
		hndlr:  mux.NewRouter(),
	}
}

func (s *Server) GetRouter() *mux.Router {
	return s.hndlr
}
func (s *Server) Run(ctx context.Context) error {
	var srv http.Server
	srv.BaseContext = func(l net.Listener) context.Context {
		return ctx
	}
	srv.Handler = s.hndlr
	addr := fmt.Sprintf("%s:%d", s.cliCtx.String("host"), s.cliCtx.Int("port"))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(l)
}
