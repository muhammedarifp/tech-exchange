package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/muhammedarifp/user/cmd/docs"
	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	services "github.com/muhammedarifp/user/usecases/interfaces"
	_ "github.com/swaggo/http-swagger"
)

type UserHandler struct {
	userUserCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUserCase: usecase,
	}
}

// @Summary User Signup
// @Description Register a new user
// @ID user-signup
// @Accept  json
// @Produce  json
// @Param req body UserSignupReq true "User signup request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/signup [post]
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
			Message:    "Temporary storage successful. Use the OTP provided to verify your account within 2 hours",
			Data:       userval,
			Errors:     nil,
		})

		if marshelErr != nil {
			panic(marshelErr.Error())
		}

		w.Write(jsonVal)
	}
}

func (h *UserHandler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body.
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println("Error reading request body:", bodyErr)
		http.Error(w, "Can't read request body", http.StatusBadRequest)
	}

	// Set the response header to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Unmarshal the request body into the user struct.
	var userEnterVal requests.UserLoginReq
	if err := json.Unmarshal(body, &userEnterVal); err != nil {
		log.Println("Error unmarshalling request body:", err)
		json_val, json_err := json.Marshal(response.Response{
			StatusCode: 400,
			Errors:     "Can't bind request body",
			Data:       nil,
			Message:    "Can't bind",
		})

		if json_err != nil {
			log.Fatal(json_err.Error())
		}

		w.Write(json_val)
		return
	}

	// Validate the user struct.
	if err := validator.New().Struct(userEnterVal); err != nil {
		jsonVal, jsonErr := json.Marshal(response.Response{
			StatusCode: 422,
			Message:    "Can't bind request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		if jsonErr != nil {
			log.Fatalln("Error marshalling JSON response:", jsonErr)
		}
		w.Write(jsonVal)
		return
	}

	// Get the user from the database.
	userVal, status_code, err := h.userUserCase.UserLogin(userEnterVal)
	if err != nil {
		w.WriteHeader(status_code)
		jsonVal, jsonErr := json.Marshal(response.Response{
			StatusCode: 422,
			Message:    "Something went wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		if jsonErr != nil {
			log.Fatalln("Error marshalling JSON response:", jsonErr)
		}
		w.Write(jsonVal)
		return
	}

	// Generate the JWT token.
	token := helperfuncs.CreateJwtToken(userVal.ID, false)

	// Marshal the response into JSON.
	jsonVal, jsonErr := json.Marshal(response.LoginResponse{
		StatusCode: 200,
		Message:    "Login success",
		Data:       userVal,
		Errors:     nil,
		Token:      token,
	})
	if jsonErr != nil {
		log.Println("Error marshalling JSON response:", jsonErr)
	}

	// Write the response to the client.
	w.Write(jsonVal)
}
