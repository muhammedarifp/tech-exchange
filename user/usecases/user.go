package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aidarkhanov/nanoid"
	"github.com/google/uuid"
	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/db"
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
	if count > 1 {
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
	if err != nil {
		return response.UserValue{}, 400, err
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

func (u *userUseCase) UserEmailVerificationSend(unique string) (bool, error) {
	// Get the application configuration
	cfg := config.GetConfig()

	// Create a Redis connection using the specified configuration
	rdb := db.CreateRedisConnection(cfg.REDIS_USER)

	// Retrieve the cached user data associated with the provided unique identifier from Redis
	res, _ := rdb.Get(context.Background(), unique).Result()
	cacheVal := cache.UserTemp{}

	// Unmarshal the JSON data retrieved from Redis into the cache.UserTemp struct
	marshelErr := json.Unmarshal([]byte(res), &cacheVal)
	if marshelErr != nil {
		return false, marshelErr
	}

	// Generate a random OTP (One-Time Password)
	otp := helperfuncs.RandomOtpGenarator()

	// Store the generated OTP and unique ID in the user repository
	if err := u.userRepo.StoreOtpAndUniqueID(cacheVal.UniqueID, otp); err != nil {
		return false, err
	}

	// Asynchronously send the verification email using a goroutine
	go func() {
		helperfuncs.SendVerificationMail(cacheVal.Email, otp, cacheVal.UniqueID)
	}()

	// Return true to indicate that the email verification process has been initiated successfully
	return true, nil
}

// UserEmailVerify verifies the user based on the provided OTP
func (u *userUseCase) UserEmailVerify(unique, otp string) (response.UserValue, error) {
	// Get the application configuration
	cfg := config.GetConfig()

	// Create a Redis connection for OTP storage
	rdb := db.CreateRedisConnection(cfg.REDIS_OTP)

	// Retrieve the stored OTP associated with the provided unique identifier from Redis
	rdbOTP, err := rdb.Get(context.Background(), unique).Result()
	if err != nil {
		// Log the error (consider using a proper logging mechanism instead of fmt.Println)
		fmt.Println("Error retrieving OTP from Redis:", err.Error())
	}

	// Compare the provided OTP with the stored OTP
	if rdbOTP == otp {
		// If the OTP is valid, proceed with user verification

		// Create a Redis connection for user data storage
		rdbUser := db.CreateRedisConnection(cfg.REDIS_USER)

		// Retrieve the stored user data associated with the provided unique identifier from Redis
		cacheStr, rdbUserErr := rdbUser.Get(context.Background(), unique).Result()
		if rdbUserErr != nil {
			// Log the error (consider using a proper logging mechanism)
			fmt.Println("Error retrieving user data from Redis:", rdbUserErr.Error())
		}

		// Unmarshal the JSON data retrieved from Redis into the cache.UserTemp struct
		var cacheVal cache.UserTemp
		if err := json.Unmarshal([]byte(cacheStr), &cacheVal); err != nil {
			// Log the error (consider using a proper logging mechanism)
			fmt.Println("Error unmarshalling user data:", err.Error())
		}

		// Call the user repository to create a new user with the retrieved data
		userVal, repoErr := u.userRepo.CreateNewUser(cacheVal)
		if repoErr != nil {
			// Return an empty user value and the repository error if user creation fails
			return response.UserValue{}, repoErr
		}

		// Return the created user value and nil error, indicating successful user creation
		return userVal, nil
	} else {
		// Return an empty user value and an error indicating an invalid OTP
		return response.UserValue{}, errors.New("invalid OTP provided")
	}
}
