package db

import (
	"github.com/muhammedarifp/tech-exchange/payments/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "user=arifu password=arifu dbname=payments port=5432 sslmode=disable TimeZone=Asia/Taipei"
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
