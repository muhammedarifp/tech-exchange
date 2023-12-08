package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
}

var (
	cfg Config
)

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	var emptyConfig Config
	if err := viper.ReadInConfig(); err != nil {
		return &emptyConfig, errors.Wrap(err, "config file read error found")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return &emptyConfig, errors.Wrap(err, "config file unmarshel error found")
	}

	return &cfg, nil
}
