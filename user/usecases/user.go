package usecases

import (
	"strings"

	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	services "github.com/muhammedarifp/user/usecases/interfaces"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (u *userUseCase) UserSignup(user requests.UserSignupReq) (response.UserValue, error) {
	// User Password Hashing
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		logrus.Info("password hashing error")
	}

	// Replace user pass into hash
	user.Password = string(hash)

	// Create username
	newUsername := strings.Split(user.Name, " ")[0] + "1"
	user.Name = newUsername

	// Call User Repository
	res, ferr := u.userRepo.UserSignup(user)

	if ferr != nil {
		return response.UserValue{}, err
	} else {
		return res, ferr
	}
}
