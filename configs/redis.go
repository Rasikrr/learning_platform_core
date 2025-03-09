package configs

import (
	"fmt"
	"time"
)

var (
	errRedisConfigRequired = fmt.Errorf("redis config error")
)

type RedisConfig struct {
	Required    bool          `toml:"required"`
	PoolSize    int           `toml:"pool_size"`
	MinIdle     int           `toml:"min_idle_conns"`
	MaxIdle     int           `toml:"max_idle_conns"`
	ReadTimeout time.Duration `toml:"read_timeout"`
}

func (c *RedisConfig) Validate() error {
	if !c.Required {
		return nil
	}
	if c.PoolSize == 0 {
		return fmt.Errorf("pool_size is empty: %w", errRedisConfigRequired)
	}
	if c.MinIdle == 0 {
		return fmt.Errorf("min_idle_conns is empty: %w", errRedisConfigRequired)
	}
	if c.MaxIdle == 0 {
		return fmt.Errorf("max_idle_conns is empty: %w", errRedisConfigRequired)
	}
	if c.ReadTimeout == 0 {
		return fmt.Errorf("read_timeout is empty: %w", errRedisConfigRequired)
	}
	return nil
}
