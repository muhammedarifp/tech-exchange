package db

import (
	"fmt"

	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db_public *gorm.DB
	Err       error
)

func ConnectDatabase(c config.Config) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", c.DB_HOST, c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT)
	db_public, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		return nil, Err
	}

	if err := db_public.AutoMigrate(
		&domain.Users{},
		&domain.Profiles{},
		&domain.Badges{},
		&domain.UserBadges{},
		&domain.Skills{},
		&domain.Followers{},
	); err != nil {
		return nil, err
	}

	return db_public, nil
}

func GetDatabase() *gorm.DB {
	return db_public
}
