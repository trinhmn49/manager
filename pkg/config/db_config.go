package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DbConfig struct {
	UserName string `env:"DB_USER_NAME"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	DbName   string `env:"DB_NAME"`
}

func LoadDbConfig() (*DbConfig, error) {
	envType := getEnvType()
	v := viper.New()
	v.AutomaticEnv() // read from OS env


	if envType == EnvTypeLocal { // if local, inject env vars from local .env file
		if err := godotenv.Load(".env"); err != nil {
			return nil, fmt.Errorf("failed to load local env file: %s", err)
		}
	}

	// parse config from env vars
	cfg := new(DbConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %s", err)
	}

	return cfg, nil
}
