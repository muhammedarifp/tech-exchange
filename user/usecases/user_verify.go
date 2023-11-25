package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/db"
)

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
