package api

import (
	"github.com/labstack/echo/v4"
	adminhandlers "github.com/muhammedarifp/content/api/handlers/admin"
	userhandlers "github.com/muhammedarifp/content/api/handlers/user"
	"github.com/muhammedarifp/content/api/middleware"
	_ "github.com/muhammedarifp/content/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerHTTP struct {
	echo *echo.Echo
}

func NewServeHTTP(userHandler *userhandlers.ContentUserHandler, adminHandler *adminhandlers.AdminContentHandler) *ServerHTTP {
	e := echo.New()

	e.GET("/contents/swagger/*", echoSwagger.EchoWrapHandler())

	userAuth := e.Group("/api/v1/contents")
	adminAuth := e.Group("/api/v1/contents/admin")
	userAuth.Use(middleware.AuthMiddleWare)

	userAuth.POST("/create-post", userHandler.CreateNewPost)
	userAuth.POST("/add-comment", userHandler.CreateComment)
	userAuth.PATCH("/like", userHandler.LikePost)
	userAuth.PUT("/update", userHandler.UpdateContent)
	userAuth.DELETE("/delete", userHandler.DeleteContent)
	userAuth.GET("/getown", userHandler.GetUserContents)
	userAuth.GET("/getall", userHandler.GetallPosts)
	userAuth.GET("/getone", userHandler.GetOnePost)
	userAuth.GET("/get-recomented", userHandler.FetchRecomentedContents)
	userAuth.POST("/follow-tag", userHandler.FollowTag)
	userAuth.GET("/getall-tags", userHandler.Getalltags)

	adminAuth.GET("/getall", adminHandler.GetallPosts)
	adminAuth.DELETE("/delete", adminHandler.DeleteContent)
	adminAuth.POST("/create-tag", adminHandler.AddNewTag)

	return &ServerHTTP{echo: e}
}

func (s *ServerHTTP) Start() {
	s.echo.Start(":8001")
}
