package api

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/muhammedarifp/content/api/handlers/user"
	"github.com/muhammedarifp/content/api/middleware"
	_ "github.com/muhammedarifp/content/cmd/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerHTTP struct {
	echo *echo.Echo
}

func NewServeHTTP(userHandler *handlers.ContentUserHandler) *ServerHTTP {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())

	userAuth := e.Group("/api/v1/contents")
	userAuth.Use(middleware.AuthMiddleWare)

	// @Swagger
	userAuth.POST("/create-post", userHandler.CreateNewPost)
	userAuth.POST("/add-comment", userHandler.CreateComment)

	return &ServerHTTP{echo: e}
}

func (s *ServerHTTP) Start() {
	s.echo.Start(":8001")
}
