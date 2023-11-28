package api

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/muhammedarifp/content/api/handlers/user"
	"github.com/muhammedarifp/content/api/middleware"
)

type ServerHTTP struct {
	echo *echo.Echo
}

func NewServeHTTP(userHandler *handlers.ContentUserHandler) *ServerHTTP {
	e := echo.New()

	userAuth := e.Group("/api/v1/contents")
	userAuth.Use(middleware.AuthMiddleWare)

	userAuth.POST("/create-post", userHandler.CreateNewPost)
	userAuth.POST("/add-comment", userHandler.CreateComment)

	return &ServerHTTP{echo: e}
}

func (s *ServerHTTP) Start() {
	s.echo.Start(":8001")
}
