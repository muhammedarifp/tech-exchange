package api

import (
	"github.com/labstack/echo/v4"
	handlers "github.com/muhammedarifp/content/api/handlers/user"
)

type ServerHTTP struct {
	echo *echo.Echo
}

func NewServeHTTP(userHandler *handlers.ContentUserHandler) *ServerHTTP {
	e := echo.New()
	e.POST("/create-post", userHandler.CreateNewPost)
	return &ServerHTTP{echo: e}
}

func (s *ServerHTTP) Start() {
	s.echo.Start(":8001")
}
