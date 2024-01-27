package db

import (
	"fmt"

	"github.com/muhammedarifp/tech-exchange/notification/config"
	"github.com/muhammedarifp/tech-exchange/notification/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	//dsn := "user=arifu password=arifu dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		domain.Notifications{},
	); err != nil {
		return nil, err
	}

	return db, nil

}
