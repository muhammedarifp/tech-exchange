package interfaces

import (
	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
)

type UserRepository interface {
	UserSignup(user cache.UserTemp) (cache.UserTemp, error)
	UserLogin(user requests.UserLoginReq) (response.UserValue, error)
	GetUserDetaUsingID(userid string) (response.UserValue, error)
	CreateNewUser(cache.UserTemp) (response.UserValue, error)
	StoreOtpAndUniqueID(unique, otp string) error
	EmailSearchOnDatabase(email string) (int, error)

	// today
	FetchUserProfileUsingID(userid string) (response.UserProfileValue, error)                              // get
	UpdateUserProfile(profile requests.UserPofileUpdate, userid string) (response.UserProfileValue, error) // put
	UpdateUserEmail(account response.UserValue, userid string) (response.UserValue, error)
	DeleteUserAccount(userid string) (response.UserValue, error)
	// FollowOrUnfollowOthers()
}
