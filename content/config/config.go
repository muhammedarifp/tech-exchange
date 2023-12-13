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
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("from config fold 1st err : " + err.Error()) // 1
		return Config{}, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("from config fold 2st err : " + err.Error()) // 2
		return Config{}, err
	}

	fmt.Println(" ========= =============== ========== ", cfg)

	return cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
