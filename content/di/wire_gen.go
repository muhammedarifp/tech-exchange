// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/muhammedarifp/content/api"
	handlers2 "github.com/muhammedarifp/content/api/handlers/admin"
	"github.com/muhammedarifp/content/api/handlers/user"
	"github.com/muhammedarifp/content/config"
	"github.com/muhammedarifp/content/db"
	"github.com/muhammedarifp/content/repository"
	"github.com/muhammedarifp/content/usecases"
)

// Injectors from wire.go:

func InitWire(cfg config.Config) (*api.ServerHTTP, error) {
	client, err := db.InitDbConnection()
	if err != nil {
		return nil, err
	}
	contentUserRepository := repository.NewContentUserRepo(client)
	contentUserUsecase := usecases.NewContentUserUsecase(contentUserRepository)
	contentUserHandler := handlers.NewContentUserHandler(contentUserUsecase)
	adminContentRepo := repository.NewAdminContentRepository(client)
	adminContentUseCase := usecases.NewAdminContentUsecase(adminContentRepo)
	adminContentHandler := handlers2.NewAdminContentHandler(adminContentUseCase)
	serverHTTP := api.NewServeHTTP(contentUserHandler, adminContentHandler)
	return serverHTTP, nil
}
