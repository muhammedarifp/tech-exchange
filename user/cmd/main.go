package main

import (
	"fmt"

	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/di"
)

func main() {
	cfg := config.LoadConfig()
	server, err := di.InitializeAPI(cfg)

	if err != nil {
		fmt.Println("1. ", err.Error())
	}
	if err := server.Start(cfg.APP_PORT); err != nil {
		fmt.Println("2. ", err.Error())
	}
}
