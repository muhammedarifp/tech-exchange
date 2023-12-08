package main

import (
	commonhelp "github.com/muhammedarifp/tech-exchange/payments/commonHelp"
	"github.com/muhammedarifp/tech-exchange/payments/config"
	"github.com/muhammedarifp/tech-exchange/payments/di"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	commonhelp.FailCriticalOnErr(cfgErr, "configuration load error")

	server, err := di.InitWire(*cfg)
	commonhelp.FailCriticalOnErr(err, "init wire error")

	server.Start()
}
