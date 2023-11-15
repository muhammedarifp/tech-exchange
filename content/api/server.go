package api

import "github.com/labstack/echo/v4"

type ServerHTTP struct {
	echo *echo.Echo
}

func NewServeHTTP() *ServerHTTP {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"hello": "hello world"})
	})
	e.Start(":8001")
	return &ServerHTTP{
		echo: e,
	}
}
