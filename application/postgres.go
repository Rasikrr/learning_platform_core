package application

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/database"
)

func (a *App) initPostgres(ctx context.Context) error {
	if !a.config.Postgres.Required {
		return nil
	}
	var err error
	a.postgres, err = database.NewPostgres(ctx, a.config)
	if err != nil {
		return err
	}
	a.closers.Add(a.postgres)
	return nil
}
