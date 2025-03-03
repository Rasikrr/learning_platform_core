package configs

type GRPCConfig struct {
	Port     int  `toml:"port"`
	Required bool `toml:"required"`
}
