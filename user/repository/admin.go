package repository

import (
	"errors"

	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{
		DB: db,
	}
}

func (d *adminDatabase) AdminLogin(admin requests.AdminRequest) (response.AdminValue, error) {
	adminVal := response.AdminValue{}
	qury := `SELECT id,username,email,password,is_admin FROM users WHERE email = $1 LIMIT 1`
	if err := d.DB.Raw(qury, admin.Email).Scan(&adminVal).Error; err != nil {
		return adminVal, err
	}

	if adminVal.ID == 0 {
		return response.AdminValue{}, errors.New("record not found")
	}

	return adminVal, nil
}
