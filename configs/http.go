package configs

import "fmt"

var (
	errHTTPConfigRequired = fmt.Errorf("http config error")
)

type HTTPConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Required bool   `toml:"required"`
}

func (c *HTTPConfig) Validate() error {
	if !c.Required {
		return nil
	}
	if c.Host == "" {
		return fmt.Errorf("host is empty: %w", errHTTPConfigRequired)
	}
	if c.Port == "" {
		return fmt.Errorf("port is empty: %w", errHTTPConfigRequired)
	}
	return nil
}
