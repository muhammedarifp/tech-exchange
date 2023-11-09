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
		fmt.Println(err.Error())
	}
	server.Start()
}
