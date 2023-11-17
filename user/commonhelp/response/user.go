package response

import (
	"time"
)

type UserValue struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Is_verified bool      `json:"is_verified"`
	Is_premium  bool      `json:"is_premium"`
	Is_banned   bool      `json:"is_banned"`
}
