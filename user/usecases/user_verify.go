package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/helperfuncs"
	"github.com/muhammedarifp/user/commonhelp/msgs"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/db"
	"github.com/muhammedarifp/user/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
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
			return response.UserValue{}, repoErr
		}

		// Return the created user value and nil error, indicating successful user creation
		conn, connErr := rabbitmq.NewRabbitmqConnection()
		if connErr != nil {
			log.Fatal("Failed to establish connection")
		}

		ch, chErr := conn.Channel()
		if chErr != nil {
			log.Fatal("Failed to open channel")
		}

		queue, queErr := ch.QueueDeclare("payment_acc", true, false, false, false, nil)
		if queErr != nil {
			log.Fatal(queErr.Error())
		}

		msg := msgs.PaymentAccount{
			UserID: userVal.ID,

			Email: userVal.Email,
			Name:  userVal.Username,
		}

		msgByte, _ := json.Marshal(msg)

		if err := ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp091.Publishing{
			ContentType: "application/json",
			Body:        msgByte,
		}); err != nil {
			log.Fatal("Failed to publish message")
		}
		return userVal, nil
	} else {
		return response.UserValue{}, errors.New("invalid OTP provided")
	}
}
