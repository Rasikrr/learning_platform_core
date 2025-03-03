package configs

type HTTPConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Required bool   `toml:"required"`
}
