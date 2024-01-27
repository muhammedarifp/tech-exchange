package db

import (
	"fmt"

	"github.com/muhammedarifp/tech-exchange/payments/config"
	"github.com/muhammedarifp/tech-exchange/payments/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(c config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", c.DB_HOST, c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	if err := db.AutoMigrate(
		domain.RazorpayAccount{},
		domain.Plans{},
		domain.Subscription{},
		domain.Payment{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

func Temp() *gorm.DB {
	dsn := "user=arifu password=arifu dbname=users port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil
	}

	return db
}
