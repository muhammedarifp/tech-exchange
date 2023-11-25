package repository

import (
	"errors"

	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/db"
)

func FetchUserUsingID_public(userid string) (*response.UserValue, error) {
	DB := db.GetDatabase()
	var emptyUserVal response.UserValue
	if userid == "" {
		return &emptyUserVal, errors.New("user value is invalid")
	}

	//
	var userVal response.UserValue
	qury := `SELECT * FROM users WHERE id = $1`
	if err := DB.Raw(qury, userid).Scan(&userVal).Error; err != nil {
		return &emptyUserVal, err
	}

	return &userVal, nil
}
