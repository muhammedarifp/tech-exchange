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
	qury := `SELECT id,username,email,created_at,password,is_active FROM users WHERE email = $1`

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
	qury := `SELECT * FROM users WHERE id = $1`
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
		RETURNING *
	)
	INSERT INTO profiles (user_id, name)
	VALUES ((SELECT id FROM inserted_user), (SELECT username FROM inserted_user))
	RETURNING (SELECT id FROM inserted_user),(SELECT username FROM inserted_user),(SELECT email FROM inserted_user),(SELECT id FROM inserted_user),(SELECT is_verified FROM inserted_user), (SELECT is_premium FROM inserted_user),(SELECT is_banned FROM inserted_user),(SELECT is_active FROM inserted_user),(SELECT created_at FROM inserted_user)
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

func (d *userDatabase) FetchUserProfileUsingID(userid string) (response.UserProfileValue, error) {
	var profileVal response.UserProfileValue
	qury := `SELECT id,user_id,name,profile_img,bio,city,github,linkedin FROM profiles WHERE user_id = $1`
	if err := d.DB.Raw(qury, userid).Scan(&profileVal).Error; err != nil {
		return response.UserProfileValue{}, err
	}

	return profileVal, nil
}

func (d *userDatabase) UpdateUserProfile(profile requests.UserPofileUpdate, userid string) (response.UserProfileValue, error) {
	qury := `UPDATE profiles
			SET name = $1, bio = $2, city = $3, github = $4, linkedin = $5 
			WHERE user_id = $6
			RETURNING id,user_id,name,profile_img,bio,city,github,linkedin
			`
	var userProfile response.UserProfileValue
	if err := d.DB.Raw(qury, profile.Name, profile.Bio, profile.City, profile.Github, profile.Linkedin, userid).Scan(&userProfile).Error; err != nil {
		return response.UserProfileValue{}, err
	}

	return userProfile, nil
}

func (d *userDatabase) UpdateUserEmail(account response.UserValue, userid string) (response.UserValue, error) {
	qury := `UPDATE users SET email = $1 WHERE id = $2 RETURNING id`
	var userVal response.UserValue
	if err := d.DB.Raw(qury, account.Email, userid).Scan(&userVal).Error; err != nil {
		return response.UserValue{}, err
	}

	return userVal, nil
}

func (d *userDatabase) DeleteUserAccount(userid string) (response.UserValue, error) {
	qury := `UPDATE users SET is_active = false WHERE id = $1`
	var userVal response.UserValue
	if err := d.DB.Raw(qury, userid).Scan(&userVal).Error; err != nil {
		return response.UserValue{}, err
	}

	return userVal, nil
}

func (d *userDatabase) UploadProfileImage(imageurl string, userid string) (response.UserProfileValue, error) {
	qury := `UPDATE profiles SET profile_img = $1 WHERE user_id = $2 RETURNING *`
	var profileVal response.UserProfileValue
	if err := d.DB.Raw(qury, imageurl, userid).Scan(&profileVal).Error; err != nil {
		return profileVal, err
	}

	return profileVal, nil
}
