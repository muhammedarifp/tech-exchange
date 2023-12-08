package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	USER_SERVICE         string
	CONTENT_SERVICE      string
	NOTIFICATION_SERVICE string
	PAYMENT_SERVICE      string
	USER_COMMON          string
	USER_ADMIN_COMMON    string
	CONTENT_COMMON       string
	CONTENT_ADMIN_COMMON string
	NOTIFICATION_COMMON  string
	PAYMENT_COMMON       string
	PAYMENT_ADMIN_COMMON string
}

var cfg Config

func InitConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
