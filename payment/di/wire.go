package di

// import (
// 	"github.com/google/wire"
// 	"github.com/muhammedarifp/tech-exchange/payments/api"
// 	"github.com/muhammedarifp/tech-exchange/payments/api/handlers"
// 	"github.com/muhammedarifp/tech-exchange/payments/config"
// 	"github.com/muhammedarifp/tech-exchange/payments/db"
// 	"github.com/muhammedarifp/tech-exchange/payments/repository"
// 	"github.com/muhammedarifp/tech-exchange/payments/usecase"
// )

// func InitWire(cfg config.Config) (*api.ServeHTTP, error) {
// 	wire.Build(
// 		// db
// 		db.ConnectDatabase,

// 		// repo
// 		repository.NewAdminPaymentDb,
// 		repository.NewUserPaymentRepo,

// 		// usecase
// 		usecase.NewAdminPaymentsUsecase,
// 		usecase.NewUserPaymentUsecase,

// 		// handlers
// 		handlers.NewAdminPaymentHandler,
// 		handlers.NewUserPaymentHandler,

// 		// final
// 		api.NewServeHTTP,
// 	)

// 	return &api.ServeHTTP{}, nil
// }
