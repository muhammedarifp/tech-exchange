package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
	DB_PORT     string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("erroooor !!!!!!")
	}

	var cfg Config
	viper.Unmarshal(&cfg)

	return cfg
}
