package response

import (
	"time"
)

type UserValue struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Is_verified bool      `json:"is_verified"`
	Is_premium  bool      `json:"is_premium"`
	Is_banned   bool      `json:"is_banned"`
	Is_active   bool      `json:"is_active"`
}

type UserProfileValue struct {
	ID         uint   `json:"-"`
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	ProfileImg string `json:"profile_img"`
	Bio        string `json:"bio"`
	City       string `json:"city"`
	Github     string `json:"github"`
	Linkedin   string `json:"linkedin"`
}
