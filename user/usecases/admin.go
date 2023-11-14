package usecases

import (
	"errors"

	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	service "github.com/muhammedarifp/user/usecases/interfaces"

	_ "github.com/muhammedarifp/user/cmd/docs"
)

type adminUsecase struct {
	AdminRepo interfaces.AdminRepository
}

func NewAdminUsecase(repo interfaces.AdminRepository) service.AdminUsecase {
	return &adminUsecase{
		AdminRepo: repo,
	}
}

func (u *adminUsecase) AdminLogin(admin requests.AdminRequest) (response.AdminValue, string, error) {
	adminVal, err := u.AdminRepo.AdminLogin(admin)
	if err != nil {
		return adminVal, "", err
	}

	if !adminVal.Is_admin {
		return response.AdminValue{}, "", errors.New("permission denaid")
	}

	if adminVal.Email == admin.Email && helperfuncs.CompareHashPassAndEnteredPass(adminVal.Password, admin.Password) {
		return response.AdminValue{}, "", errors.New("incorrect username or password")
	}

	token := helperfuncs.CreateJwtToken(adminVal.ID, true)

	return adminVal, token, nil
}
