package di

// import (
// 	"github.com/google/wire"
// 	"github.com/muhammedarifp/content/api"
// 	handlers "github.com/muhammedarifp/content/api/handlers/user"
// 	"github.com/muhammedarifp/content/config"
// 	"github.com/muhammedarifp/content/db"
// 	"github.com/muhammedarifp/content/repository"
// 	"github.com/muhammedarifp/content/usecases"
// )

// func InitWire(cfg config.Config) (*api.ServerHTTP, error) {
// 	wire.Build(
// 		//db
// 		db.InitDbConnection,

// 		// repositry
// 		repository.NewContentUserRepo,

// 		// usecase
// 		usecases.NewContentUserUsecase,

// 		// handlers
// 		handlers.NewContentUserHandler,

// 		// server
// 		api.NewServeHTTP,
// 	)

// 	return &api.ServerHTTP{}, nil
// }
