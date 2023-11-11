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
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}

	adminEnterVal := requests.AdminRequest{}
	if err := json.Unmarshal(body, &adminEnterVal); err != nil {
		jsonResp, err := json.Marshal(response.Response{
			StatusCode: 404,
			Message:    "cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		if err != nil {
			log.Println(err.Error())
		}
		w.Write(jsonResp)
		return
	}

	adminVal, usercaseErr := h.AdminUsecase.AdminLogin(adminEnterVal)
	if usercaseErr != nil {
		jsonResp, err := json.Marshal(response.Response{
			StatusCode: 404,
			Message:    "login not success",
			Data:       nil,
			Errors:     usercaseErr.Error(),
		})
		if err != nil {
			log.Println(err.Error())
		}
		w.Write(jsonResp)
		return
	}

	jsonResp, err := json.Marshal(response.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       adminVal,
		Errors:     nil,
	})
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(jsonResp)

}
