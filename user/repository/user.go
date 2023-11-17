package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/muhammedarifp/user/commonhelp/cache"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/db"
	interfaces "github.com/muhammedarifp/user/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: db}
}

func (d *userDatabase) UserSignup(user cache.UserTemp) (cache.UserTemp, error) {
	rdb_user := db.CreateRedisConnection(2)
	rdb_email := db.CreateRedisConnection(3)
	ctx := context.Background()

	if res, _ := rdb_email.Exists(ctx, user.Email).Result(); res == 1 {
		uniqueid, _ := rdb_email.Get(ctx, user.Email).Result()
		dataByte, _ := rdb_user.Get(ctx, uniqueid).Result()
		data := cache.UserTemp{}
		json.Unmarshal([]byte(dataByte), &data)
		return data, nil
	}

	user_marshel, _ := json.Marshal(user)
	if err := rdb_user.Set(context.Background(), user.UniqueID, user_marshel, time.Hour*1).Err(); err != nil {
		return cache.UserTemp{}, err
	}

	if err := rdb_email.Set(context.Background(), user.Email, user.UniqueID, time.Hour*1).Err(); err != nil {
		return cache.UserTemp{}, err
	}

	return user, nil
}

func (d *userDatabase) UserLogin(user requests.UserLoginReq) (response.UserValue, error) {
	cfg := config.GetConfig()
	qury := `SELECT id,username,email,created_at,password FROM users WHERE email = $1`

	rdb := db.CreateRedisConnection(cfg.REDIS_EMAIL)
	rdbStat, _ := rdb.Exists(context.Background(), user.Email).Result()
	if rdbStat == 1 {
		return response.UserValue{}, fmt.Errorf("account activation required: please check your email to verify and activate your account")
	}

	userVal := response.UserValue{}

	if err := d.DB.Raw(qury, user.Email).Scan(&userVal).Error; err != nil {
		fmt.Println(err.Error())
		return response.UserValue{}, err
	}

	return userVal, nil
}

func (d *userDatabase) GetUserDetaUsingID(userid string) (response.UserValue, error) {
	qury := `SELECT id,username,created_at,email,is_verified FROM users WHERE id = $1`
	userData := response.UserValue{}
	err := d.DB.Raw(qury, userid).Scan(&userData).Error
	if err != nil {
		return userData, err
	}
	return userData, nil
}

func (d *userDatabase) StoreOtpAndUniqueID(userid, otp string) error {
	cfg := config.GetConfig()
	rdb := db.CreateRedisConnection(cfg.REDIS_OTP)
	status := rdb.Set(context.Background(), userid, otp, time.Minute*5)
	if status.Err() != nil {
		return status.Err()
	} else {
		return nil
	}
}

func (d *userDatabase) CreateNewUser(user cache.UserTemp) (response.UserValue, error) {
	newQury := `
	WITH inserted_user AS (
		INSERT INTO users (username, email, password, is_verified)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, is_verified, is_banned, username, email, is_premium
	)
	INSERT INTO profiles (user_id, name)
	VALUES ((SELECT id FROM inserted_user), (SELECT username FROM inserted_user))
	`
	var userVal response.UserValue
	if err := d.DB.Raw(newQury, user.Username, user.Email, user.Password, true).Scan(&userVal).Error; err != nil {
		return response.UserValue{}, err
	}

	rdb_email := db.CreateRedisConnection(config.GetConfig().REDIS_EMAIL)
	if err := rdb_email.Del(context.Background(), user.Email).Err(); err != nil {
		log.Println(err.Error())
	}

	return userVal, nil
}

func (d *userDatabase) EmailSearchOnDatabase(email string) (int, error) {
	qury := `SELECT COUNT(*) FROM users WHERE email = $1`
	var count int
	if err := d.DB.Raw(qury, email).Scan(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
