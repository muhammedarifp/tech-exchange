package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/tech-exchange/notification/api/handlers"
	"github.com/muhammedarifp/tech-exchange/notification/config"
	_ "github.com/muhammedarifp/tech-exchange/notification/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

// Package docs provides documentation for your API.
// @title Nofifications
// @description The Notification Service API allows you to manage and retrieve notifications. It provides endpoints for creating, retrieving, and managing notifications for users.
// @version 1.0
func NewServeHTTP(notificationHandler *handlers.NotificationsHandler) *ServerHTTP {
	app := gin.New()

	app.Use(gin.Logger())

	app.GET("/notifications/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Create a coustom route group
	notification := app.Group("/api/v1/users/notification")
	{
		notification.GET("/getall", notificationHandler.GetallNotifications)
	}

	return &ServerHTTP{
		engine: app,
	}
}

func (s *ServerHTTP) Start() {
	cfg := config.GetConfig()
	s.engine.Run(cfg.APP_PORT)
}
