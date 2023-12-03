package usecases

import (
	"errors"
	"strings"

	"github.com/aidarkhanov/nanoid"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/muhammedarifp/user/commonhelp/cache"
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

func (u *userUseCase) UserSignup(user requests.UserSignupReq) (cache.UserTemp, error) {
	// Replace user password with its hash
	user.Password = helperfuncs.PaaswordToHash(user.Password)

	// Check if the email already exists in the database
	count, searchErr := u.userRepo.EmailSearchOnDatabase(user.Email)
	if searchErr != nil {
		return cache.UserTemp{}, searchErr
	}

	// If the email count is greater than 1, it means the email already exists
	if count >= 1 {
		return cache.UserTemp{}, errors.New("this email already exists")
	}

	// Generate a unique identifier (uniqueid) for creating a unique username
	uniqueID, uniqueErr := nanoid.Generate("AaBcDdEeFfGgHhIiJjKkLlMmNnOoPpQqSsTtUuVvWwXxYyZz1234567890", 4)
	if uniqueErr != nil {
		return cache.UserTemp{}, errors.New("failed to create username")
	}

	// Create a new username using the first part of the user's name and the unique identifier
	newUsername := strings.Split(user.Name, " ")[0] + uniqueID
	user.Name = newUsername

	// Create a new UserTemp object with uniqueID, newUsername, user's email, and hashed password
	newCache := cache.UserTemp{
		UniqueID: uuid.New().String(),
		Username: newUsername,
		Email:    user.Email,
		Password: user.Password,
	}

	// Call User Repository to persist the new user data
	res, ferr := u.userRepo.UserSignup(newCache)

	if ferr != nil {
		return cache.UserTemp{}, ferr
	} else {
		return res, nil
	}
}

func (u *userUseCase) UserLogin(user requests.UserLoginReq) (response.UserValue, int, error) {
	// Call the User Repository to retrieve user details based on the login request
	userVal, err := u.userRepo.UserLogin(user)

	// Check for errors during user retrieval
	var responseEmpty response.UserValue
	if err != nil {
		return responseEmpty, 400, err
	}

	if !userVal.Is_active {
		return responseEmpty, 400, errors.New("account creation failed. your account has been permanently deactivated. if you believe this is an error, please contact support")
	}

	// Check if the entered email and password match the retrieved user's credentials
	if userVal.Email == user.Email && helperfuncs.CompareHashPassAndEnteredPass(userVal.Password, user.Password) {
		// If credentials match, return user details, HTTP status code 200, and no error
		return userVal, 200, nil
	} else {
		// If credentials do not match, return an empty user value, HTTP status code 401, and an authentication error
		return response.UserValue{}, 401, errors.New("username or password is invalid")
	}
}

func (u *userUseCase) UpdateUserEmail(account response.UserValue, userid string) (response.UserValue, error) {
	var userValueEmpty response.UserValue
	if !govalidator.IsEmail(account.Email) {
		return userValueEmpty, errors.New("invalid email provided")
	}

	userVal, repoErr := u.userRepo.UpdateUserEmail(account, userid)
	if repoErr != nil {
		return userValueEmpty, repoErr
	}

	return userVal, nil
}

func (u *userUseCase) DeleteUserAccount(userid string) (response.UserValue, error) {
	var userValueEmpty response.UserValue
	if userid == "" {
		return userValueEmpty, errors.New("UserID cannot be empty")
	}
	userVal, repoErr := u.userRepo.DeleteUserAccount(userid)
	if repoErr != nil {
		return userValueEmpty, repoErr
	}

	return userVal, nil
}

func (u *userUseCase) FetchUserAccount(userid string) (response.UserValue, error) {
	val, err := u.userRepo.GetUserDetaUsingID(userid)
	if err != nil {
		return response.UserValue{}, err
	}

	return val, nil
}
