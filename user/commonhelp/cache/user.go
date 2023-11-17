package cache

type UserTemp struct {
	UniqueID string `json:"unique"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
