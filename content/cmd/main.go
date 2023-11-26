package main

import (
	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/di"
)

func main() {
	cfg, _ := config.LoadConfig()
	server, _ := di.InitWire(cfg)
	server.Start()
}
