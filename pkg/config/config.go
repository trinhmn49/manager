package config

import (
	"manager/pkg/utils"
	"github.com/spf13/viper"
)

type AppConfig struct {
	RunMode string `mapstructure:"RUN_MODE"`
	Port    int    `mapstructure:"PORT"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort int    `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER_NAME"`
	DBName string `mapstructure:"DB_NAME"`
	DBPwd  string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string, cfg *AppConfig) {
	dirPath := utils.GetDirectoryPath(path)
	fileName, err := utils.GetFileName(path)
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}
}
