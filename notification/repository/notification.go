package repository

import (
	"log"
	"sync"

	commonhelp "github.com/muhammedarifp/tech-exchange/notification/commonHelp"
	interfaces "github.com/muhammedarifp/tech-exchange/notification/repository/interface"
	"gorm.io/gorm"
)

var (
	mu sync.Mutex
)

type notificationsDB struct {
	DB *gorm.DB
}

func NewNotificationRepo(db *gorm.DB) interfaces.NotificationRepo {
	return &notificationsDB{DB: db}
}

func (d *notificationsDB) StoreNotificationsOnDB(notification commonhelp.NotificationResp) {
	mu.Lock()
	defer mu.Unlock()
	query := ``
	if err := d.DB.Raw(query).Error; err != nil {
		log.Fatalf("Errororoo")
	}

	if notification.IsImportent {
		// send mail
	}
}
