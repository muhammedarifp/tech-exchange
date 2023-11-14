package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/usecases/interfaces"
)

type AdminHandler struct {
	AdminUsecase interfaces.AdminUsecase
}

func NewAdminHandler(usecase interfaces.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		AdminUsecase: usecase,
	}
}

func (h *AdminHandler) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Read the request body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())

		// Write the error response to the client.
		http.Error(w, "Can't read request body", http.StatusBadRequest)
		return
	}

	// Unmarshal the request body into the admin struct.
	var adminEnterVal requests.AdminRequest
	if err := json.Unmarshal(body, &adminEnterVal); err != nil {
		log.Println(err.Error())

		// Write the error response to the client.
		jsonResp, err := json.Marshal(response.LoginResponse{
			StatusCode: 404,
			Message:    "Can't bind",
			Data:       nil,
			Errors:     err.Error(),
			Token:      nil,
		})
		if err != nil {
			log.Println(err.Error())
		}
		w.Write(jsonResp)
		return
	}

	// Get the admin from the database and generate the JWT token.
	adminVal, token, usercaseErr := h.AdminUsecase.AdminLogin(adminEnterVal)

	if usercaseErr != nil {
		log.Println(usercaseErr.Error())

		// Write the error response to the client.
		jsonResp, err := json.Marshal(response.LoginResponse{
			StatusCode: 404,
			Message:    "Login failed",
			Data:       nil,
			Errors:     usercaseErr.Error(),
			Token:      nil,
		})
		if err != nil {
			log.Println(err.Error())
		}

		w.Write(jsonResp)
		return
	}

	// Marshal the response into JSON.
	jsonResp, err := json.Marshal(response.LoginResponse{
		StatusCode: 200,
		Message:    "Login success",
		Data:       adminVal,
		Errors:     nil,
		Token:      token,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Write the response to the client.
	w.Write(jsonResp)
}
