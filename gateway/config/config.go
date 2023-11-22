package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	USER_SERVICE    string
	CONTENT_SERVICE string
}

var cfg Config

func InitConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("dir fetch error")
	}

	fmt.Println("Current working dir is : ", dir)

	exDir := filepath.Dir(dir)
	fmt.Println("Current exdir is : ", exDir)

	viper.SetConfigName(".env")
	viper.AddConfigPath("")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}
}

func GetConfig() *Config {
	return &cfg
}
