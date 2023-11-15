package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type UserUseCase interface {
	UserSignup(user requests.UserSignupReq) (response.UserValue, error)
	UserLogin(user requests.UserLoginReq) (response.UserValue, int, error)
	UserEmailVerificationSend(token string) (bool, error)
	UserEmailVerify(otp requests.UserEmailVerificationReq, token string) (response.UserValue, error)
}
