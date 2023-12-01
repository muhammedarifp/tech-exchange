package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MONGO_PORT string
	DB_NAME    string
	JWT_SECRET string
}

var (
	cfg Config
)

func LoadConfig() (Config, error) {
	var emptyConfig Config
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("from config fold 1st err : " + err.Error()) // 1
		return emptyConfig, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("from config fold 2st err : " + err.Error()) // 2
		return emptyConfig, err
	}

	return cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
