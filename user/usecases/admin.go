package usecases

import (
	"errors"

	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	service "github.com/muhammedarifp/user/usecases/interfaces"
)

type adminUsecase struct {
	AdminRepo interfaces.AdminRepository
}

func NewAdminUsecase(repo interfaces.AdminRepository) service.AdminUsecase {
	return &adminUsecase{
		AdminRepo: repo,
	}
}

func (u *adminUsecase) AdminLogin(admin requests.AdminRequest) (response.AdminValue, error) {
	adminVal, err := u.AdminRepo.AdminLogin(admin)
	if err != nil {
		return adminVal, err
	}

	if adminVal.Email == admin.Email && helperfuncs.CompareHashPassAndEnteredPass(adminVal.Password, admin.Password) {
		return response.AdminValue{}, errors.New("incorrect username or password")
	}

	return adminVal, nil
}
