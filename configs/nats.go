package configs

import (
	"errors"
)

type NATSConfig struct {
	Required bool   `toml:"required"`
	DSN      string `toml:"dsn"`
	Queue    string `toml:"queue"`
}

func (c *NATSConfig) Validate() error {
	if !c.Required {
		return nil
	}
	if c.DSN == "" {
		return errors.New("nats config error: DSN is empty")
	}
	return nil
}
