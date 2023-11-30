package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	adminhandlers "github.com/muhammedarifp/user/api/handlers/admin"
	userhandlers "github.com/muhammedarifp/user/api/handlers/user"
	"github.com/muhammedarifp/user/api/middleware"

	_ "github.com/muhammedarifp/user/cmd/docs"
)

type ServerHTTP struct {
	engine *mux.Router
}

func NewServerHTTP(userHandler *userhandlers.UserHandler, adminHandler *adminhandlers.AdminHandler) *ServerHTTP {
	// Create a new mux router.
	engine := mux.NewRouter()

	// Add a logging middleware to all routes.
	engine.Use(middleware.LoggingMiddleware)

	// Serve the Swagger UI documentation.
	//engine.PathPrefix("/api/users/swagger/").Handler(httpSwagger.WrapHandler)
	//engine.Handle("/api/users/swagger.json", http.FileServer(http.Dir("docs")))
	// engine.HandleFunc("/api/users/swagger/", httpSwagger.Handler(
	// 	httpSwagger.URL("github.com/muhammedarifp/user/cmd/docs"),
	// ))

	userRouter := engine.PathPrefix("/api/v1/users").Subrouter()
	userAuthRouter := engine.PathPrefix("/api/v1/users").Subrouter()
	adminRouter := engine.PathPrefix("/api/v1/users/admins").Subrouter()

	// User endpoints
	userRouter.HandleFunc("/signup", userHandler.Signup).Methods("POST")
	userRouter.HandleFunc("/login", userHandler.Login).Methods("POST")
	userRouter.HandleFunc("/otp/send", userHandler.RequestOtp).Methods("POST")
	userRouter.HandleFunc("/otp/verify", userHandler.VerifyOtp).Methods("POST")

	// Admin endpoints
	adminRouter.HandleFunc("/login", adminHandler.Login).Methods("POST")
	adminRouter.HandleFunc("/{userid}/ban", adminHandler.BanUser).Methods("POST")
	adminRouter.HandleFunc("/{page}/getallusers", adminHandler.GetAllUsers).Methods("GET")

	// User authentication routes
	userAuthRouter.Use(middleware.AuthMiddleware)

	userAuthRouter.HandleFunc("/profile", userHandler.ViewProfile).Methods("GET")
	userAuthRouter.HandleFunc("/account", userHandler.ViewAccount).Methods("GET")
	userAuthRouter.HandleFunc("/update-profile", userHandler.UpdateProfile).Methods("PUT")
	userAuthRouter.HandleFunc("/delete-account", userHandler.DeleteAccount).Methods("DELETE")
	userAuthRouter.HandleFunc("/upload-profile-image", userHandler.UploadProfileImage).Methods("POST")

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
