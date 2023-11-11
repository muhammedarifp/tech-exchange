package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type AdminUsecase interface {
	AdminLogin(admin requests.AdminRequest) (response.AdminValue, error)
}
