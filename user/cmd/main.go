package main

import (
	"fmt"

	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/di"

	_ "github.com/muhammedarifp/user/cmd/docs"
)

// @title Tech Exchange
// @version 1.0
// @description This is a content sharing platform like medium or dev.to
// @host localhost:8000
// @BasePath /api/v1
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
