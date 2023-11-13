package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type UserRepository interface {
	UserSignup(user requests.UserSignupReq) (response.UserValue, error)
	UserLogin(user requests.UserLoginReq) (response.UserValue, error)
	GetUserDetaUsingID(userid string) (response.UserValue, error)
	VerifyUserAccount(userid, otp string) (response.UserValue, error)
	StoreOtpAndUniqueID(unique, otp string) error
}
