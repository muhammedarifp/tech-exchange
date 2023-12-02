package di

// import (
// 	"github.com/google/wire"
// 	"github.com/muhammedarifp/tech-exchange/notification/api"
// 	"github.com/muhammedarifp/tech-exchange/notification/api/handlers"
// 	"github.com/muhammedarifp/tech-exchange/notification/config"
// 	"github.com/muhammedarifp/tech-exchange/notification/db"
// 	"github.com/muhammedarifp/tech-exchange/notification/repository"
// 	"github.com/muhammedarifp/tech-exchange/notification/usecase"
// )

// func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
// 	wire.Build(
// 		db.ConnectDatabase,

// 		// Repo
// 		repository.NewNotificationRepo,

// 		//Usecase
// 		usecase.NewNotificationUseCase,

// 		// handlers
// 		handlers.NewNotificationHandler,

// 		// api
// 		api.NewServeHTTP,
// 	)
// 	return &api.ServerHTTP{}, nil
// }
