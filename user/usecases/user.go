package usecases

import (
	"errors"
	"fmt"
	"strings"

	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	services "github.com/muhammedarifp/user/usecases/interfaces"
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

	// Replace user pass into hash
	user.Password = helperfuncs.PaaswordToHash(user.Password)

	// Create username
	newUsername := strings.Split(user.Name, " ")[0] + "1"
	user.Name = newUsername

	// Call User Repository
	res, ferr := u.userRepo.UserSignup(user)

	fmt.Println("-->", ferr)

	if ferr != nil {
		return response.UserValue{}, ferr
	} else {
		return res, nil
	}
}

func (u *userUseCase) UserLogin(user requests.UserLoginReq) (response.UserValue, int, error) {
	userVal, err := u.userRepo.UserLogin(user)
	if err != nil {
		return response.UserValue{}, 400, err
	}

	if !userVal.Is_verified {
		return response.UserValue{}, 403, errors.New("your email address has not been verified. please verify your email address before logging in")
	}

	if userVal.Email == user.Email && helperfuncs.CompareHashPassAndEnteredPass(userVal.Password, user.Password) {
		return userVal, 200, nil
	} else {
		return response.UserValue{}, 401, errors.New("username or password is invalid")
	}
}

func (u *userUseCase) UserEmailVerificationSend(token string) (bool, error) {
	userid, err := helperfuncs.GetUserIdFromJwt(token)
	if err != nil {
		fmt.Println("Token not valid")
		return false, err
	}

	userVal, repoErr := u.userRepo.GetUserDetaUsingID(userid)
	if repoErr != nil {
		return false, err
	}

	if userVal.Is_verified {
		return false, errors.New("the user with the email address '" + userVal.Email + "' has already verified their email address")
	}

	otp := helperfuncs.RandomOtpGenarator()

	helperfuncs.SendVerificationMail(userVal.Email, otp, userid)
	if err := u.userRepo.StoreOtpAndUniqueID(userid, otp); err != nil {
		return false, err
	}

	return true, nil
}

func (u *userUseCase) UserEmailVerify(otp requests.UserEmailVerificationReq, token string) (response.UserValue, error) {
	userid, err := helperfuncs.GetUserIdFromJwt(token)

	userVal, _ := u.userRepo.GetUserDetaUsingID(userid)

	fmt.Println(userVal.Is_verified)

	if userVal.Is_verified {
		return response.UserValue{}, errors.New("the user with the email address '" + userVal.Email + "' has already verified their email address")
	}

	if err != nil {
		return response.UserValue{}, fmt.Errorf("invalid auth token")
	}

	userVal, repo_err := u.userRepo.VerifyUserAccount(userid, otp.Otp)
	if repo_err != nil {
		return userVal, repo_err
	}

	return userVal, nil
}
