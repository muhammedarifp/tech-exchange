package requests

import "time"

type UserValue struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       User        `json:"data"`
	Errors     interface{} `json:"error"`
}

type User struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Is_verified bool      `json:"is_verified"`
	Is_premium  bool      `json:"is_premium"`
	Is_banned   bool      `json:"is_banned"`
	Is_active   bool      `json:"is_active"`
}
