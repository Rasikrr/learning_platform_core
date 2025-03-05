package application

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/http"
)

// nolint: unparam
func (a *App) initHTTP(ctx context.Context) error {
	if !a.config.HTTP.Required {
		return nil
	}
	a.httpServer = http.NewServer(ctx, a.config)
	a.starters.Add(a.httpServer)
	a.closers.Add(a.httpServer)
	return nil
}
