package db

import (
	"fmt"

	"github.com/muhammedarifp/user/config"
	"github.com/muhammedarifp/user/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&domain.Users{},
		&domain.Profiles{},
		&domain.Badges{},
		&domain.UserBadges{},
		&domain.Skills{},
		&domain.Followers{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
