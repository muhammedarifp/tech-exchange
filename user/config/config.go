package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	APP_PORT              string
	DB_USER               string
	DB_NAME               string
	DB_PASSWORD           string
	DB_PORT               string
	JWT_SECRET            string
	EMAIL                 string
	EMAIL_PASSWORD        string
	REDIS_OTP             int
	REDIS_EMAIL           int
	REDIS_USER            int
	AWS_REGION            string
	AWS_ACCESS_KEYID      string
	AWS_SECRET_ACCESS_KEY string
	BUCKET_NAME           string
}

var cfg Config

func LoadConfig() Config {
	viper.AddConfigPath("")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	viper.Unmarshal(&cfg)

	return cfg
}

func GetConfig() Config {
	return cfg
}
