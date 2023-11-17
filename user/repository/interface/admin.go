package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type AdminRepository interface {
	AdminLogin(admin requests.AdminRequest) (response.AdminValue, error)
	BanUser(userid string) (response.UserValue, error)
}
