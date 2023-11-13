package usecases

import (
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

func (u *userUseCase) UserLogin(user requests.UserLoginReq) (response.UserValue, error) {
	userVal, err := u.userRepo.UserLogin(user)
	if err != nil {
		fmt.Println(err.Error())
		return response.UserValue{}, err
	}

	fmt.Println(userVal)
	return userVal, nil
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

	otp := helperfuncs.RandomOtpGenarator()

	helperfuncs.SendVerificationMail(userVal.Email, otp, userid)
	if err := u.userRepo.StoreOtpAndUniqueID(userid, otp); err != nil {
		return false, err
	}

	return true, nil
}

func (u *userUseCase) UserEmailVerify(otp requests.UserEmailVerificationReq, token string) (response.UserValue, error) {
	userid, err := helperfuncs.GetUserIdFromJwt(token)
	if err != nil {
		return response.UserValue{}, fmt.Errorf("invalid auth token")
	}

	userVal, repo_err := u.userRepo.VerifyUserAccount(userid, otp.Otp)
	if repo_err != nil {
		return userVal, repo_err
	}

	return userVal, nil
}
