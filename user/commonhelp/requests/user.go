package requests

type UserLoginReq struct {
	Email    string `json:"email" validate:"required,email,omitempty"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserSignupReq struct {
	Name     string `json:"name,omitempty" validate:"required,omitempty"`
	Email    string `json:"email,omitempty" validate:"email,required"`
	Password string `json:"password,omitempty" validate:"min=6"`
}

type UserEmailVerificationReq struct {
	Otp string `json:"otp,omitempty" validate:"required,min=5"`
}

type UserPofileUpdate struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	City     string `json:"city"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
}
