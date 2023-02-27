package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DbUrl    string `mapstructure:"TEST_DB_URL"`
	DbDriver string `mapstructure:"DB_DRIVER"`
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.Unmarshal(&config)
	return
}
