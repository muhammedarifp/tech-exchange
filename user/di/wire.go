package di

// import (
// 	"github.com/google/wire"
// 	"github.com/muhammedarifp/user/api"
// 	"github.com/muhammedarifp/user/api/handlers"
// 	"github.com/muhammedarifp/user/config"
// 	"github.com/muhammedarifp/user/db"
// 	"github.com/muhammedarifp/user/repository"
// 	"github.com/muhammedarifp/user/usecases"
// )

// func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
// 	wire.Build(
// 		db.ConnectDatabase,

// 		repository.NewUserRepository,
// 		repository.NewAdminRepository,

// 		usecases.NewUserUseCase,
// 		usecases.NewAdminUsecase,

// 		handlers.NewUserHandler,
// 		handlers.NewAdminHandler,

// 		api.NewServerHTTP,
// 	)

// 	return &api.ServerHTTP{}, nil
// }
