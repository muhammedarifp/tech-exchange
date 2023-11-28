package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/muhammedarifp/user/commonhelp/response"
)

func (u *UserHandler) RequestOtp(w http.ResponseWriter, r *http.Request) {
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

func (u *UserHandler) VerifyOtp(w http.ResponseWriter, r *http.Request) {
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
