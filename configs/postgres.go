package configs

import (
	"time"
)

type PostgresConfig struct {
	Required            bool          `toml:"required"`
	MaxConns            int           `toml:"max_conns"`
	MinConns            int           `toml:"min_conns"`
	MaxIdleConnIdleTime time.Duration `toml:"max_idle_conn_time"`
}
