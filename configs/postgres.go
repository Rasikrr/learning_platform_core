package configs

import (
	"fmt"
	"time"
)

var (
	errPostgresConfigRequired = fmt.Errorf("postgres config error")
)

type PostgresConfig struct {
	Required            bool          `toml:"required"`
	MaxConns            int           `toml:"max_conns"`
	MinConns            int           `toml:"min_conns"`
	MaxIdleConnIdleTime time.Duration `toml:"max_idle_conn_time"`
}

func (c *PostgresConfig) Validate() error {
	if !c.Required {
		return nil
	}
	if c.MaxConns == 0 {
		return fmt.Errorf("max_conns is empty: %w", errPostgresConfigRequired)
	}
	if c.MinConns == 0 {
		return fmt.Errorf("min_conns is empty: %w", errPostgresConfigRequired)
	}
	if c.MaxIdleConnIdleTime == 0 {
		return fmt.Errorf("max_idle_conn_time is empty: %w", errPostgresConfigRequired)
	}
	return nil
}
