package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DbUrl                string        `mapstructure:"TEST_DB_URL"`
	DbDriver             string        `mapstructure:"DB_DRIVER"`
	Environment          string        `mapstructure:"ENVIRONMENT"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	PasetoKey            string        `mapstructure:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
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
