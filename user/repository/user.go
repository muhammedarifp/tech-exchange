package repository

import (
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: db}
}

func (d *userDatabase) UserSignup(user requests.UserSignupReq) (response.UserValue, error) {
	insertQury := `INSERT INTO users(username,email,password) VALUES ($1,$2,$3)
				RETURNING id,username,email,password,is_active`
	userVal := new(response.UserValue)
	err := d.DB.Raw(insertQury, user.Name, user.Email, user.Password).Scan(&userVal).Error
	if err != nil {
		return response.UserValue{}, err
	} else {
		return *userVal, err
	}
}
