package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Rasikrr/learning_platform_core/configs/appenv"
	"github.com/Rasikrr/learning_platform_core/enum"
	"github.com/Rasikrr/learning_platform_core/interfaces"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	devFileConfig  = "./configs/config.toml"
	prodFileConfig = "./configs/config_prod.toml"
	envFile        = ".env"
)

type Config struct {
	Name        string           `toml:"name"`
	Environment enum.Environment `toml:"-"`
	HTTP        *HTTPConfig      `toml:"http"`
	Postgres    *PostgresConfig  `toml:"postgres"`
	Redis       *RedisConfig     `toml:"redis"`
	GRPC        *GRPCConfig      `toml:"grpc"`
	NATS        *NATSConfig      `toml:"nats"`
	Variables   *Variables       `toml:"env"`
	Env         *Variables       `toml:"-"`
}

func initConfig(name string) Config {
	return Config{
		Name:      name,
		Variables: NewVariablesInstance(),
	}
}

func Parse(appName string) (Config, error) {
	// 0) create config instance
	config := initConfig(appName)

	// 1) set .env variables
	if err := config.SetEnv(); err != nil {
		return Config{}, fmt.Errorf("error in set env: %w", err)
	}
	env := config.Env.Get(appenv.AppEnv).GetString()

	envEnum, err := enum.EnvironmentString(env)
	if err != nil {
		return Config{}, fmt.Errorf("invalid environment: %s", env)
	}
	log.Printf("Initializing config for env: %s\n", envEnum.String())

	// 3) set app env
	config.SetAppEnv(envEnum)

	// 4) load config from file
	if err := config.loadFromFile(); err != nil {
		return Config{}, err
	}
	// 5) validate env and variables
	if err := config.Env.Validate(); err != nil {
		return Config{}, err
	}
	// 6) validate config
	if err := config.Validate(); err != nil {
		return Config{}, err
	}
	return config, nil
}

func getFileName(env enum.Environment) string {
	switch env {
	case enum.EnvironmentDev:
		return devFileConfig
	case enum.EnvironmentProd:
		return prodFileConfig
	default:
		return ""
	}
}

func (c *Config) SetAppEnv(env enum.Environment) {
	c.Environment = env
}

func (c *Config) SetEnv() error {
	err := godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	envs, err := godotenv.Read(envFile)
	if err != nil {
		return fmt.Errorf("failed to read .env file: %w", err)
	}
	c.Env = c.Variables
	c.Env.Range()
	if err := c.Env.Collect(envs); err != nil {
		return err
	}
	return nil
}

func (c *Config) loadFromFile() error {
	fileName := getFileName(c.Environment)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("config file not found: %s", fileName)
	}
	if _, err := toml.DecodeFile(fileName, c); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}
	return nil
}

func (c *Config) Validate() error {
	for _, v := range []interfaces.Validatable{
		c.HTTP,
		c.Postgres,
		c.Redis,
		c.GRPC,
		c.NATS,
	} {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) GetEnvironment() enum.Environment {
	return c.Environment
}
