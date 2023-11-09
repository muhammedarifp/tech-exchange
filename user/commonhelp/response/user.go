package response

import "time"

type UserValue struct {
	ID        uint      `json:"id" gorm:"unique;not null"`
	UserName  string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_time"`
}

type UserLoginResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	UserValue
}

type UserSignupResp struct {
	Status  bool
	Message string
	UserValue
}
