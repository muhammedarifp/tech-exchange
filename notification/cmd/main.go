package main

import (
	"log"

	"github.com/muhammedarifp/tech-exchange/notification/config"
	"github.com/muhammedarifp/tech-exchange/notification/db"
	"github.com/muhammedarifp/tech-exchange/notification/di"
	"github.com/muhammedarifp/tech-exchange/notification/rabbitmq"
)

func init() {
	rabbitmq.NewRabbitmqConnection()
	go func() {

	}()
}

// @title Nofifications
// @description The Notification Service API allows you to manage and retrieve notifications. It provides endpoints for creating, retrieving, and managing notifications for users.
// @version 1.0
func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		return
	}

	server, diErr := di.InitializeAPI(*cfg)
	if diErr != nil {
		log.Fatalf("di error found : %v", diErr)
	}

	server.Start()

	db.ConnectDatabase(*cfg)
}
