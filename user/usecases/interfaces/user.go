package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type UserUseCase interface {
	UserSignup(user requests.UserSignupReq) (response.UserValue, error)
}
