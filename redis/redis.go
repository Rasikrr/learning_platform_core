package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/Rasikrr/learning_platform_core/configs"
	"github.com/Rasikrr/learning_platform_core/configs/appenv"
	redis "github.com/redis/go-redis/v9"
	"log"
	"strings"
	"time"
)

type cache struct {
	client *redis.Client
	prefix string
}

func NewRedisCache(ctx context.Context, cfg *configs.Config) (Cache, error) {
	addr := hostPort(
		cfg.Env.Get(appenv.RedisHost).GetString(),
		cfg.Env.Get(appenv.RedisPort).GetInt(),
	)

	opt := &redis.Options{
		Addr:         addr,
		Username:     cfg.Env.Get(appenv.RedisUser).GetString(),
		Password:     cfg.Env.Get(appenv.RedisPassword).GetString(),
		DB:           cfg.Env.Get(appenv.RedisDB).GetInt(),
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdle,
		MaxIdleConns: cfg.Redis.MaxIdle,
		ReadTimeout:  cfg.Redis.ReadTimeout,
	}

	client := redis.NewClient(opt)
	
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &cache{
		client: client,
		prefix: prefixKey(cfg.Name),
	}, nil
}

func (r *cache) Close(_ context.Context) error {
	log.Println("closing redis")
	return r.client.Close()
}

func (r *cache) Get(ctx context.Context, key string) (any, error) {
	k := r.getKey(key)
	val, err := r.redisStringCmd(ctx, k).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return val, nil
}

func (r *cache) Set(ctx context.Context, key string, value any) error {
	k := r.getKey(key)
	return r.client.Set(ctx, k, value, 0).Err()
}

func (r *cache) SetWithExpiration(ctx context.Context, key string, value any, expiration time.Duration) error {
	k := r.getKey(key)
	return r.client.Set(ctx, k, value, expiration).Err()
}

func (r *cache) Exists(ctx context.Context, key string) (bool, error) {
	k := r.getKey(key)
	exists, err := r.client.Exists(ctx, k).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (r *cache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	k := r.getKey(key)
	return r.client.Expire(ctx, k, expiration).Err()
}

func (r *cache) Delete(ctx context.Context, key string) error {
	k := r.getKey(key)
	return r.client.Del(ctx, k).Err()
}

func (r *cache) RPush(ctx context.Context, key string, value ...any) error {
	k := r.getKey(key)
	return r.client.RPush(ctx, k, value...).Err()
}

func (r *cache) LPush(ctx context.Context, key string, value ...any) error {
	k := r.getKey(key)
	return r.client.LPush(ctx, k, value...).Err()
}

func (r *cache) LLen(ctx context.Context, key string) (int64, error) {
	k := r.getKey(key)
	return r.client.LLen(ctx, k).Result()
}

func (r *cache) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	k := r.getKey(key)
	return r.client.LRange(ctx, k, start, stop).Result()
}

func (r *cache) Flush(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}

func (r *cache) redisStringCmd(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *cache) getKey(k string) string {
	return fmt.Sprintf("%s:%s", r.prefix, k)
}

func hostPort(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func prefixKey(key string) string {
	return strings.ToUpper(key)
}
