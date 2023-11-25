package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

func (u *UserHandler) FetchUserProfileUsingIDHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")
	w.Header().Set("Content-Type", "application/json")

	// token variable iclude empty value
	if token == "" {
		// Invalid token provided
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "1invalid auth token provided",
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

	// Fetch userid from token
	userid, tokenFetchErr := helperfuncs.GetUserIdFromJwt(token)
	if tokenFetchErr != nil {
		// userid is not available in this token
		// Server Respond error
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "2invalid auth token provided",
			Data:       nil,
			Errors:     tokenFetchErr.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	//
	profileVal, profileErr := u.userUserCase.FetchUserProfileUsingID(userid)
	if profileErr != nil {
		// profile value fetching error
		// Server respond error
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "3invalid auth token provided",
			Data:       nil,
			Errors:     profileErr.Error(),
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
		StatusCode: 400,
		Message:    "Success",
		Data:       profileVal,
		Errors:     nil,
	})
	if marshelErr != nil {
		log.Println(marshelErr.Error())
	}
	w.WriteHeader(200)
	w.Write(marshelResp)
}

// Update user profile
func (u *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	// Fetch Auth Token
	token := r.Header.Get("Token")

	userid, tokenErr := helperfuncs.GetUserIdFromJwt(token)
	if tokenErr != nil {
		// This error is failed to fetch user id from token
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "invalid auth token provided",
			Data:       nil,
			Errors:     tokenErr.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	// Case : 2
	//
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		// handle body reading error
		fmt.Println(bodyErr.Error())
		return
	}

	var profile requests.UserPofileUpdate
	if err := json.Unmarshal(body, &profile); err != nil {
		// error
		// handler or respond that error
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "unmarshel error",
			Data:       nil,
			Errors:     err.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	profileVal, usecaseErr := u.userUserCase.UpdateUserProfile(profile, userid)
	if usecaseErr != nil {
		// Handler this usecase error
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "invalid auth token provided",
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
		Message:    "success",
		Data:       profileVal,
		Errors:     nil,
	})
	if marshelErr != nil {
		log.Println(marshelErr.Error())
	}
	w.WriteHeader(200)
	w.Write(marshelResp)
}

// Upload new profile image
func (u *UserHandler) UploadNewProfileImage(w http.ResponseWriter, r *http.Request) {
	userid, _ := helperfuncs.GetUserIdFromJwt(r.Header.Get("Token"))
	maxsize := 5 * 1024 * 1024
	if err := r.ParseMultipartForm(int64(maxsize)); err != nil { // max 10 mb ;
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	image, header, err := r.FormFile("image")
	if err != nil {
		marshelResp, marshelErr := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		if marshelErr != nil {
			log.Println(marshelErr.Error())
		}
		w.WriteHeader(400)
		w.Write(marshelResp)
		return
	}

	profileVal, usecaseErr := u.userUserCase.UploadNewProfilePhoto(image, *header, userid)
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
	}

	marshelResp, marshelErr := json.Marshal(response.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       profileVal,
		Errors:     nil,
	})
	if marshelErr != nil {
		log.Println(marshelErr.Error())
	}
	w.WriteHeader(200)
	w.Write(marshelResp)
}
