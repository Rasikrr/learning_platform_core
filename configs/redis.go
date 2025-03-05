package configs

import "time"

type RedisConfig struct {
	Required    bool          `toml:"required"`
	PoolSize    int           `toml:"pool_size"`
	MinIdle     int           `toml:"min_idle_conns"`
	MaxIdle     int           `toml:"max_idle_conns"`
	ReadTimeout time.Duration `toml:"read_timeout"`
}
