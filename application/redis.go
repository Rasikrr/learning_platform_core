package application

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/redis"
)

func (a *App) initRedis(ctx context.Context) error {
	if !a.config.Redis.Required {
		return nil
	}
	var err error
	a.redis, err = redis.NewRedisCache(ctx, a.config)
	if err != nil {
		return err
	}
	a.closers.Add(a.redis)
	return nil
}
