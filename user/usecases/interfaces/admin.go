package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type AdminUsecase interface {
	AdminLogin(admin requests.AdminRequest) (response.AdminValue, string, error)
	BanUser(userid string) (response.UserValue, error)
	GetallUsers(page int) ([]response.UserValue, error)
}
