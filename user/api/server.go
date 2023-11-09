package api

import (
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

	engine.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	engine.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello User"))
	})

	return &ServerHTTP{engine: engine}
}

func (r *ServerHTTP) Start() error {
	if err := http.ListenAndServe(":8080", r.engine); err != nil {
		return err
	} else {
		return nil
	}
}
