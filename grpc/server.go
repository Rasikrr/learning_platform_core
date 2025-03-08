package grpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	port   int
	server *grpc.Server
}

func NewServer(port int) *Server {
	return &Server{
		port:   port,
		server: newGrpcServer(),
	}
}

func (s *Server) Start(_ context.Context) error {
	lis, err := net.Listen("tcp", s.addr("0.0.0.0"))
	if err != nil {
		return err
	}
	log.Println("starting grpc server")
	if err := s.server.Serve(lis); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			return nil
		}
		return err
	}
	return nil
}

func (s *Server) Close(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}

func (s *Server) Srv() *grpc.Server {
	return s.server
}

func (s *Server) addr(host string) string {
	return fmt.Sprintf("%s:%d", host, s.port)
}

func newGrpcServer() *grpc.Server {
	return grpc.NewServer()
}
