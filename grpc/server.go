package grpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime/debug"
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
	return grpc.NewServer(
		grpc.StreamInterceptor(streamPanicRecoveryInterceptor),
		grpc.UnaryInterceptor(unaryPanicRecoveryInterceptor),
	)
}

func streamPanicRecoveryInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	_ *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in stream: %v\n%s", r, debug.Stack())
		}
	}()
	return handler(srv, ss)
}

func unaryPanicRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in unary: %v\n%s", r, debug.Stack())
			err = status.Errorf(codes.Internal, "internal server error")
		}
	}()
	return handler(ctx, req)
}
