package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type UserUseCase interface {
	UserSignup(user requests.UserSignupReq) (cache.UserTemp, error)
	UserLogin(user requests.UserLoginReq) (response.UserValue, int, error)
	UserEmailVerificationSend(token string) (bool, error)
	UserEmailVerify(unique, otp string) (response.UserValue, error)
}
