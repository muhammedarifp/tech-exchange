package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
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

// @Summary Get user by ID
// @Description Get a user by its ID
// @ID get-user-by-id
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) UserSignupHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create new UserSignupReq object
	// For using managing user entered data
	userReq := requests.UserSignupReq{}

	// Unmarshel user enterd data into json
	// str -> json using json package
	if err := json.Unmarshal(body, &userReq); err != nil {
		fmt.Println(err.Error())
	}

	// Create validator new instance
	// For using validate user entered data using throw validator package
	validate := validator.New()
	w.Header().Set("Content-Type", "application/json")

	// if not validate user entered data return error response
	// response like response.Response model
	if err := validate.Struct(&userReq); err != nil {
		resp := response.Response{
			StatusCode: 422,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		}

		// Struct -> Json
		val, err := json.Marshal(resp)
		if err != nil {
			panic(err.Error())
		}

		// Setup header. content type is json
		// Finally write data
		w.Write(val)
		return
	}

	// throw any error from usecases
	// respond error response
	userval, usecaseErr := h.userUserCase.UserSignup(userReq)
	if usecaseErr != nil {
		jsonVal, err := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     usecaseErr.Error(),
		})

		if err != nil {
			panic(err.Error())
		}

		w.Write(jsonVal)
		return
	}

	// Success response from server
	if usecaseErr == nil {
		jsonVal, marshelErr := json.Marshal(response.Response{
			StatusCode: 201,
			Message:    "user signup successfully",
			Data:       userval,
			Errors:     nil,
		})

		if marshelErr != nil {
			panic(marshelErr.Error())
		}

		w.Write(jsonVal)
	}
}

// User Login function
func (h *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	userEnterVal := requests.UserLoginReq{}
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		fmt.Println("1. ", bodyErr.Error())
	}

	// Setup header
	// Our response is json
	w.Header().Set("Content-Type", "application/json")

	// Unmarshel user entered data
	// Using json package
	if err := json.Unmarshal(body, &userEnterVal); err != nil {
		fmt.Println("2. ", err.Error())
	}

	// Struct validation using validator package
	// For using simplycity
	if err := validator.New().Struct(userEnterVal); err != nil {
		jsonVal, jsonErr := json.Marshal(response.Response{
			StatusCode: 422,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})

		if jsonErr != nil {
			panic(jsonErr.Error())
		}

		w.Write(jsonVal)
		return
	}

	userVal, err := h.userUserCase.UserLogin(userEnterVal)
	if err != nil {
		jsonVal, jsonErr := json.Marshal(response.Response{
			StatusCode: 422,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})

		if jsonErr != nil {
			panic(jsonErr.Error())
		}

		w.Write(jsonVal)
		return
	}

	if userEnterVal.Email == userVal.Email && helperfuncs.CompareHashPassAndEnteredPass(userVal.Password, userEnterVal.Password) {
		token := helperfuncs.CreateJwtToken(userVal.ID)
		jsonVal, jsonErr := json.Marshal(response.LoginResponse{
			StatusCode: 200,
			Message:    "Login success",
			Data:       userVal,
			Errors:     nil,
			Token:      token,
		})
		if jsonErr != nil {
			panic("myraaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
		w.Write(jsonVal)
	} else {
		jsonVal, jsonErr := json.Marshal(response.LoginResponse{
			StatusCode: 401,
			Message:    "Login failure",
			Data:       nil,
			Errors:     "email or password is incorrect",
			Token:      nil,
		})
		if jsonErr != nil {
			panic("myraaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		}
		w.Write(jsonVal)
	}

	helperfuncs.CreateJwtToken(userVal.ID)
}

// Email verification handler

func (u *UserHandler) VerifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Token")

	// Case 1
	// token is inavlid case
	if token == "" {
		jsonVal, _ := json.Marshal(response.Response{
			StatusCode: 400,
			Message:    "token is invalid",
			Data:       nil,
			Errors:     "noken is invalid",
		})

		w.Write(jsonVal)
		return
	}

	// Case 2
	//
	if status, _ := u.userUserCase.UserEmailVerify(token); !status {
		w.Write([]byte("Somthing wrong"))
	} else {
		w.Write([]byte("Okkk !!"))
	}
}
