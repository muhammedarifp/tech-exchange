package response

import "time"

type AdminValue struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Is_admin  bool      `json:"is_admin"`
	Password  string    `json:"-"`
}
