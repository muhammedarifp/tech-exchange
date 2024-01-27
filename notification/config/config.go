package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	APP_PORT    string
	DB_NAME     string
	DB_HOST     string
	JWT_SECRET  string
}

var (
	cfg Config
)

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper file read error : %v", err)
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("viper unmarsher error : %v", err)
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
