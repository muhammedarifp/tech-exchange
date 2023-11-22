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

//
// ----------------------------------------------
//

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

func (u *UserHandler) UserRequestOtpHandler(w http.ResponseWriter, r *http.Request) {
	qury := r.URL.Query()
	unique := qury["unique"]
	fmt.Println(unique)
	// Get the token from the request header.
	w.Header().Set("Content-Type", "application/json")
	// Check if the token is empty.
	if len(unique) == 0 {
		// Return a 400 Bad Request error with the message "token is invalid".
		w.WriteHeader(400)
		json_resp, _ := json.Marshal(response.ReqOtpResp{
			StatusCode: 400,
			Status:     false,
			Message:    "your unique activation token is empty",
			Error:      "token is invalid",
		})
		w.Write(json_resp)
		return
	}

	if _, err := u.userUserCase.UserEmailVerificationSend(unique[0]); err != nil {
		w.WriteHeader(400)
		json_resp, _ := json.Marshal(response.ReqOtpResp{
			StatusCode: 400,
			Status:     false,
			Message:    "your unique activation token is invalid",
			Error:      err.Error(),
		})
		w.Write(json_resp)
		return
	}

	// Write a success response to the client.
	json_resp, _ := json.Marshal(response.ReqOtpResp{
		StatusCode: 200,
		Status:     true,
		Message:    "OTP sent to your email successfully! Your OTP is valid for 10 minutes. If you do not receive your OTP within 2 minutes, please try again",
		Error:      nil,
	})
	w.Write(json_resp)
}

func (u *UserHandler) VerifyUserOtpHandler(w http.ResponseWriter, r *http.Request) {
	// Get the token from the request header.
	qurys := r.URL.Query()
	unique := qurys.Get("unique")
	otp := qurys.Get("otp")
	w.Header().Set("Content-Type", "application/json")

	// Decode the request body.
	if unique == "" || otp == "" {
		json_resp, _ := json.Marshal(response.VerifyOtpResp{
			StatusCode: 400,
			Status:     false,
			Message:    "unique token or otp is empty",
			Error:      "SOMETHING_WENT_WRONG",
			Data:       nil,
		})
		w.Write(json_resp)
		return
	}

	// Validate the OTP.
	userVal, err := u.userUserCase.UserEmailVerify(unique, otp)
	if err != nil {
		json_resp, _ := json.Marshal(response.VerifyOtpResp{
			StatusCode: 400,
			Status:     false,
			Message:    err.Error(),
			Error:      "SOMETHING_WENT_WRONG",
			Data:       nil,
		})
		w.Write(json_resp)
		return
	}

	json_resp, _ := json.Marshal(response.VerifyOtpResp{
		StatusCode: 200,
		Status:     true,
		Message:    "OTP verification successful!",
		Error:      nil,
		Data:       userVal,
	})
	w.Write(json_resp)
}

// FetchUserProfileUsingID(userid string) (response.UserProfileValue, error)                              // get - done
// UpdateUserProfile(profile response.UserProfileValue, userid string) (response.UserProfileValue, error) // put - done
// UpdateUserEmail(account response.UserValue, userid string) (response.UserValue, error)                 // put - progressing ..
// DeleteUserAccount(userid string) (response.UserValue, error) 										  // delete - next

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
		Message:    "invalid auth token provided",
		Data:       profileVal,
		Errors:     nil,
	})
	if marshelErr != nil {
		log.Println(marshelErr.Error())
	}
	w.WriteHeader(200)
	w.Write(marshelResp)
}

func (u *UserHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	// Fetch Auth Token
	token := w.Header().Get("Token")
	if token == "" {
		// Invalid token provided
		// Write this soon !!
		return
	}
	userid, tokenErr := helperfuncs.GetUserIdFromJwt(token)
	if tokenErr != nil {
		// handle this error
		// This error is failed to fetch user id from token
		return
	}

	// Case : 2
	//
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		// handle body reading error
		return
	}

	var profile requests.UserPofileUpdate
	if json.Unmarshal(body, profile) != nil {
		// error
		// handler or respond that error
	}

	profileVal, usecaseErr := u.userUserCase.UpdateUserProfile(profile, userid)
	if usecaseErr != nil {
		// Handler this usecase error
		//  return this error message only
		return
	}

	// Success response
	fmt.Println(profileVal)
	// Finally updated user profile
}

// func (u *UserHandler) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
// 	u.userUserCase.UpdateUserEmail()
// }

func (u *UserHandler) DeleteUserAccount(w http.ResponseWriter, r *http.Request) {
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
