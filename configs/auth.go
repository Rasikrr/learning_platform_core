package configs

import "time"

type AuthConfig struct {
	Secret               string        `toml:"-"`
	AccessTokenLifeTime  time.Duration `toml:"access_token_lifetime"`
	RefreshTokenLifeTime time.Duration `toml:"refresh_token_lifetime"`
}
