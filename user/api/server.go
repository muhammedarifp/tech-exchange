package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammedarifp/user/api/handlers"
	"github.com/muhammedarifp/user/api/middleware"
	_ "github.com/muhammedarifp/user/cmd/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ServerHTTP struct {
	engine *mux.Router
}

// @Summary Create a new user
// @Description Create a new user with the specified details
// @Tags users
// @Accept json
// @Produce json
// @Param name formData string true "User's name"
// @Param email formData string true "User's email"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/users/create [post]
func NewServerHTTP(userHandler *handlers.UserHandler, adminHandler *handlers.AdminHandler) *ServerHTTP {
	// Create a new mux router.
	engine := mux.NewRouter()

	// Add a logging middleware to all routes.
	engine.Use(middleware.LoggingMiddleware)

	// Serve the Swagger UI documentation.
	engine.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Create a subrouter for the user routes.
	userRouter := engine.PathPrefix("/api/users").Subrouter()

	// Create a subrouter for the user authentication routes.
	userAuthRouter := engine.PathPrefix("/api/users").Subrouter()

	// Create a subrouter for the admin routes.
	adminRouter := engine.PathPrefix("/api/admins").Subrouter()

	// Add the user handlers.
	userRouter.HandleFunc("/signup", userHandler.UserSignupHandler).Methods("POST")
	userRouter.HandleFunc("/login", userHandler.UserLoginHandler).Methods("POST")

	// Add the admin handler.
	adminRouter.HandleFunc("/login", adminHandler.AdminLoginHandler).Methods("POST")

	// Add a middleware to the user authentication routes to check if the user is authenticated.
	userAuthRouter.Use(middleware.AuthUserMiddleware)

	// Add the user authentication handlers.
	userAuthRouter.HandleFunc("/otp/send", userHandler.UserRequestOtpHandler).Methods("POST")
	userAuthRouter.HandleFunc("/otp/verify", userHandler.VerifyUserOtpHandler).Methods("POST")

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
