package interfaces

import (
	"mime/multipart"

	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type UserUseCase interface {
	UserSignup(user requests.UserSignupReq) (cache.UserTemp, error)
	UserLogin(user requests.UserLoginReq) (response.UserValue, int, error)

	// user verification
	UserEmailVerificationSend(token string) (bool, error)
	UserEmailVerify(unique, otp string) (response.UserValue, error)

	// user profile
	FetchUserProfileUsingID(userid string) (response.UserProfileValue, error)                                                  // get
	UpdateUserProfile(profile requests.UserPofileUpdate, userid string) (response.UserProfileValue, error)                     // put
	UploadNewProfilePhoto(photo multipart.File, header multipart.FileHeader, userid string) (response.UserProfileValue, error) // post

	// user account
	UpdateUserEmail(account response.UserValue, userid string) (response.UserValue, error) // put
	DeleteUserAccount(userid string) (response.UserValue, error)                           // delete
	FetchUserAccount(userid string) (response.UserValue, error)                            // get
}
