package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

const ServiceName = "gRPC"

type Server struct {
	cliCtx *cli.Context
	srv    *grpc.Server
}

func New(ctx *cli.Context) *Server {
	return &Server{
		cliCtx: ctx,
		srv:    grpc.NewServer(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", s.cliCtx.String("grpc-host"), s.cliCtx.Int("grpc-host"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.srv.Serve(lis)
}
