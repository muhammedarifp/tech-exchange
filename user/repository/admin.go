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

func (d *adminDatabase) BanUser(userid string) (response.UserValue, error) {
	qury := `UPDATE users SET is_banned = true WHERE id = $1 RETURNING id,username,email,is_verified,is_banned`
	userVal := response.UserValue{}
	if err := d.DB.Raw(qury, userid).Scan(&userVal).Error; err != nil {
		return response.UserValue{}, err
	}

	if userVal.ID == 0 {
		return response.UserValue{}, errors.New("user not found")
	}

	return userVal, nil
}

func (d *adminDatabase) GetallUsers(page int) ([]response.UserValue, error) {
	limit := 10
	offset := (page - 1) * limit
	qury := `SELECT * FROM users OFFSET $1 LIMIT $2`
	var users []response.UserValue
	if err := d.DB.Raw(qury, offset, limit).Scan(&users).Error; err != nil {
		return []response.UserValue{}, err
	}

	return users, nil
}
