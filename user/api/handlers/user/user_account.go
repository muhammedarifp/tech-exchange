package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/response"
)

func (u *UserHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	w.Header().Set("Content-Type", "application/json")
	if token == "" {
		// Case is token is empty or null
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "invalid auth token provided",
			Data:       nil,
			Errors:     "Invalid auth token provided",
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}
	userid, tokenErr := helperfuncs.GetUserIdFromJwt(token)
	if tokenErr != nil {
		// case is userid fetch error from user auth token
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "invalid auth token provided",
			Data:       nil,
			Errors:     "Invalid auth token provided",
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}
	userVal, usecaseErr := u.userUserCase.DeleteUserAccount(userid)
	if usecaseErr != nil {
		// handle this case
		// This is a usecase layer error
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "internal server error",
			Data:       nil,
			Errors:     usecaseErr.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	// Success response
	marshelResp, marshelErr := json.Marshal(response.Response{
		StatusCode: 200,
		Message:    "user account deactivated",
		Data:       userVal,
		Errors:     nil,
	})
	if marshelErr != nil {
		log.Println(marshelErr.Error())
	}
	w.WriteHeader(200)
	w.Write(marshelResp)
}

//	func (u *UserHandler) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
//		u.userUserCase.UpdateUserEmail()
//	}
func (u *UserHandler) ViewAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userid, _ := helperfuncs.GetUserIdFromJwt(r.Header.Get("Token"))
	if userid == "" {
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     errors.New("invalid userid provided"),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	accountVal, usecaseErr := u.userUserCase.FetchUserAccount(userid)
	if usecaseErr != nil {
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     usecaseErr.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	marshelResp, marshelErr := json.Marshal(response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       accountVal,
		Errors:     nil,
	})
	if marshelErr != nil {
		log.Println(marshelErr.Error())
	}
	w.WriteHeader(200)
	w.Write(marshelResp)

}
