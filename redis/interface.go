package redis

import (
	"errors"
	"github.com/Rasikrr/learning_platform_core/interfaces"
	"golang.org/x/net/context"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Cache interface {
	interfaces.Closer
	Get(ctx context.Context, key string) (any, error)
	Set(ctx context.Context, key string, value any) error
	Exists(ctx context.Context, key string) (bool, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	SetWithExpiration(ctx context.Context, key string, value any, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Flush(ctx context.Context) error
	RPush(ctx context.Context, key string, value ...any) error
	LPush(ctx context.Context, key string, value ...any) error
	LLen(ctx context.Context, key string) (int64, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
}
