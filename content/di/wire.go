package di

// import (
// 	"github.com/google/wire"
// 	"github.com/muhammedarifp/content/api"
// 	adminhandlers "github.com/muhammedarifp/content/api/handlers/admin"
// 	userhandlers "github.com/muhammedarifp/content/api/handlers/user"
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
// 		repository.NewAdminContentRepository,

// 		// usecase
// 		usecases.NewContentUserUsecase,
// 		usecases.NewAdminContentUsecase,

// 		// handlers
// 		userhandlers.NewContentUserHandler,
// 		adminhandlers.NewAdminContentHandler,

// 		// server
// 		api.NewServeHTTP,
// 	)

// 	return &api.ServerHTTP{}, nil
// }
