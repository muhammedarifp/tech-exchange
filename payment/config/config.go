package config

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	APP_PORT              string
	RAZORPAY_KEY          string
	RAZORPAY_SEC          string
	JWT_SECRET            string
	PAYMENT_ACCOUNT_QUEUE string
}

var cfg Config

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	var emptyConfig Config
	if err := viper.ReadInConfig(); err != nil {
		return &emptyConfig, errors.Wrap(err, "config file read error found")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return &emptyConfig, errors.Wrap(err, "config file unmarshel error found")
	}

	fmt.Println(cfg)

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
