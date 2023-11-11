package requests

type AdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
