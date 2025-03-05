package application

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/configs"
	"github.com/Rasikrr/learning_platform_core/database"
	coreGrpc "github.com/Rasikrr/learning_platform_core/grpc"
	"github.com/Rasikrr/learning_platform_core/http"
	"github.com/Rasikrr/learning_platform_core/redis"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
)

type App struct {
	name       string
	config     *configs.Config
	redis      redis.Cache
	postgres   *database.Postgres
	httpServer *http.Server
	grpcServer *coreGrpc.Server
	starters   *Starters
	closers    *Closers
}

func NewApp(ctx context.Context, name string) *App {
	cfg, err := configs.Parse(name)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	return NewAppWithConfig(ctx, name, &cfg)
}

func NewAppWithConfig(ctx context.Context, name string, cfg *configs.Config) *App {
	app := &App{
		name:     name,
		config:   cfg,
		starters: NewStarters(),
		closers:  NewClosers(),
	}
	if err := app.initPostgres(ctx); err != nil {
		log.Fatalf("failed to init postgres: %v", err)
	}
	if err := app.initRedis(ctx); err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}
	if err := app.initGRPC(ctx); err != nil {
		log.Fatalf("failed to init grpc: %v", err)
	}

	if err := app.initHTTP(ctx); err != nil {
		log.Fatalf("failed to init http: %v", err)
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	stopChan := make(chan struct{})
	go a.GracefulShutdown(ctx, stopChan)
	if err := a.start(ctx); err != nil {
		return err
	}
	<-stopChan

	return nil
}

func (a *App) start(ctx context.Context) error {
	g := errgroup.Group{}
	for _, s := range a.starters.starters {
		g.Go(func() error {
			return s.Start(ctx)
		})
	}
	return g.Wait()
}

func (a *App) Close(ctx context.Context) error {
	for _, c := range a.closers.closers {
		if err := c.Close(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) GracefulShutdown(ctx context.Context, stopChan chan struct{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)
	<-sigChan
	if err := a.Close(ctx); err != nil {
		log.Printf("error while closing app: %v", err)
	}
	close(stopChan)
}

func (a *App) GrpcServer() *coreGrpc.Server {
	return a.grpcServer
}

func (a *App) Postgres() *database.Postgres {
	return a.postgres
}

func (a *App) HTTPServer() *http.Server {
	return a.httpServer
}
