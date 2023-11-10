package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammedarifp/user/api/handlers"
	"github.com/muhammedarifp/user/api/middleware"
)

type ServerHTTP struct {
	engine *mux.Router
}

func NewServerHTTP(userHandler *handlers.UserHandler) *ServerHTTP {
	engine := mux.NewRouter()

	engine.Use(middleware.LoggingMiddleware)

	// User Routes
	userRouter := engine.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/signup", userHandler.UserSignupHandler).Methods("POST")
	userRouter.HandleFunc("/login", userHandler.UserLoginHandler).Methods("POST")
	userRouter.HandleFunc("/isuserexist")

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
