package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammedarifp/user/api/handlers"
	"github.com/muhammedarifp/user/api/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type ServerHTTP struct {
	engine *mux.Router
}

func NewServerHTTP(userHandler *handlers.UserHandler, adminHandler *handlers.AdminHandler) *ServerHTTP {
	engine := mux.NewRouter()

	engine.Use(middleware.LoggingMiddleware)

	engine.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("/home/arifu/Desktop/Tech Exchange/user/cmd/docs/swagger.json")))).Handler(httpSwagger.WrapHandler)

	// User Routes
	userRouter := engine.PathPrefix("/user").Subrouter()
	userAuthRouter := engine.PathPrefix("/user").Subrouter()
	adminRouter := engine.PathPrefix("/admin").Subrouter()

	// User handlers
	userRouter.HandleFunc("/signup", userHandler.UserSignupHandler).Methods("POST")
	userRouter.HandleFunc("/login", userHandler.UserLoginHandler).Methods("POST")

	// Admin handlers
	adminRouter.HandleFunc("/login", adminHandler.AdminLoginHandler).Methods("POST")

	//
	userAuthRouter.Use(middleware.AuthUserMiddleware)
	userAuthRouter.HandleFunc("/verify-email", userHandler.VerifyEmailHandler).Methods("POST")

	return &ServerHTTP{engine: engine}
}

func (r *ServerHTTP) Start(port string) error {
	fmt.Println("Server Starting ...!")
	if err := http.ListenAndServe(port, r.engine); err != nil {
		return err
	} else {
		return nil
	}
}
