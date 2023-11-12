package repository

import (
	"fmt"

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
	userVal := response.UserValue{}
	err := d.DB.Raw(insertQury, user.Name, user.Email, user.Password).Scan(&userVal).Error
	if err != nil {
		return response.UserValue{}, err
	} else {
		fmt.Println(userVal)
		return userVal, nil
	}
}

func (d *userDatabase) UserLogin(user requests.UserLoginReq) (response.UserValue, error) {
	qury := `SELECT id,username,email,created_at,password FROM users WHERE email = $1`
	userVal := response.UserValue{}

	if err := d.DB.Raw(qury, user.Email).Scan(&userVal).Error; err != nil {
		fmt.Println(err.Error())
		return response.UserValue{}, err
	}

	return userVal, nil
}

func (d *userDatabase) GetUserDetaUsingID(userid string) (response.UserValue, error) {
	qury := `SELECT id,username,created_at,email FROM users WHERE id = $1`
	userData := response.UserValue{}
	err := d.DB.Raw(qury, userid).Scan(&userData).Error
	if err != nil {
		return userData, err
	}
	return userData, nil
}
