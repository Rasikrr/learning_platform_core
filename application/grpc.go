package application

import (
	"context"
	coreGrpc "github.com/Rasikrr/learning_platform_core/grpc"
)

// nolint: unparam
func (a *App) initGRPC(_ context.Context) error {
	if !a.config.GRPC.Required {
		return nil
	}
	a.grpcServer = coreGrpc.NewServer(a.config.GRPC.Port)
	a.starters.Add(a.grpcServer)
	a.closers.Add(a.grpcServer)
	return nil
}
