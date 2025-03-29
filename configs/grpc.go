package configs

import "fmt"

var (
	errGRPCConfigRequired = fmt.Errorf("grpc config error")
)

type GRPCConfig struct {
	Port     int  `toml:"port"`
	Required bool `toml:"required"`
}

func (c GRPCConfig) Validate() error {
	if !c.Required {
		return nil
	}
	if c.Port == 0 {
		return fmt.Errorf("port is empty: %w", errGRPCConfigRequired)
	}
	return nil
}
