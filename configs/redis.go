package configs

type RedisConfig struct {
	Required bool   `toml:"required"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}
