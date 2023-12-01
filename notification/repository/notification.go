package repository

import (
	interfaces "github.com/muhammedarifp/tech-exchange/notification/repository/interface"
	"gorm.io/gorm"
)

type notificationsDB struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) interfaces.NotificationRepo {
	return &notificationsDB{DB: db}
}

// func (d *notificationsDB) GetallNotifications(userid string) {
// 	fmt.Println("Ok")
// }
