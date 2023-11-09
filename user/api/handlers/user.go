package handlers

import (
	"net/http"

	services "github.com/muhammedarifp/user/usecases/interfaces"
)

type UserHandler struct {
	userUserCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUserCase: usecase,
	}
}

func (h *UserHandler) UserSignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
