package requests

type UserLoginReq struct {
	Email    string `validate:"required,email,omitempty"`
	Password string `validate:"required,min=6"`
}

type UserSignupReq struct {
	Name     string `validate:"required,omitempty"`
	Email    string `validate:"email,reqired"`
	Password string `validate:"required,min=6"`
}
