package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	adminhandlers "github.com/muhammedarifp/user/api/handlers/admin"
	userhandlers "github.com/muhammedarifp/user/api/handlers/user"
	"github.com/muhammedarifp/user/api/middleware"
	//_ "github.com/muhammedarifp/user/cmd/docs"
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
func NewServerHTTP(userHandler *userhandlers.UserHandler, adminHandler *adminhandlers.AdminHandler) *ServerHTTP {
	// Create a new mux router.
	engine := mux.NewRouter()

	// Add a logging middleware to all routes.
	engine.Use(middleware.LoggingMiddleware)

	// Serve the Swagger UI documentation.
	// engine.PathPrefix("/api/users/swagger/").Handler(httpSwagger.WrapHandler)
	// engine.Handle("/api/users/swagger.json", http.FileServer(http.Dir("docs")))

	// Serve Swagger UI
	engine.PathPrefix("/api/users/swagger/*any").Handler(http.StripPrefix("/api/users/swagger/", http.FileServer(http.Dir("docs"))))
	engine.Handle("/swagger.json", http.FileServer(http.Dir("docs")))

	// Create a subrouter for the user routes.
	userRouter := engine.PathPrefix("/api/users").Subrouter()

	// Create a subrouter for the user authentication routes.
	userAuthRouter := engine.PathPrefix("/api/users").Subrouter()

	// Create a subrouter for the admin routes.
	adminRouter := engine.PathPrefix("/api/admins").Subrouter()

	// Add the user handlers.
	userRouter.HandleFunc("/signup", userHandler.UserSignupHandler).Methods("POST")
	userRouter.HandleFunc("/login", userHandler.UserLoginHandler).Methods("POST")
	userRouter.HandleFunc("/otp/send", userHandler.UserRequestOtpHandler).Methods("POST")
	userRouter.HandleFunc("/otp/verify", userHandler.VerifyUserOtpHandler).Methods("POST")

	// Add the admin handler.
	adminRouter.HandleFunc("/login", adminHandler.AdminLoginHandler).Methods("POST")
	adminRouter.HandleFunc("/users/ban/{userid}", adminHandler.AdminBanUserHandler)

	// Add a middleware to the user authentication routes to check if the user is authenticated.
	userAuthRouter.Use(middleware.AuthUserMiddleware)

	// Add the user authentication handlers.
	userAuthRouter.HandleFunc("/view-profile", userHandler.FetchUserProfileUsingIDHandler).Methods("GET") //
	userAuthRouter.HandleFunc("/update-profile", userHandler.UpdateUserProfile).Methods("PUT")
	userAuthRouter.HandleFunc("/delete-acc", userHandler.DeleteUserAccount).Methods("DELETE") // working
	userAuthRouter.HandleFunc("/upload-profileimg", userHandler.UploadNewProfileImage).Methods("POST")

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
