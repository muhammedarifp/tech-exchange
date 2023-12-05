package repository

import (
	"sync"

	commonhelp "github.com/muhammedarifp/tech-exchange/notification/commonHelp"
	"github.com/muhammedarifp/tech-exchange/notification/domain"
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

func (d *notificationsDB) StoreNotificationsOnDB(notification commonhelp.NotificationResp) bool {
	query := `INSERT INTO notifications(user_id,title,body,is_importent,liked_user) 
			VALUES ($1,$2,$3,$4,$5)
			RETURNING *`

	var n domain.Notifications
	err := d.DB.Raw(query, notification.UserID, notification.Title, notification.Body, notification.IsImportent, notification.LikedUserID).Scan(&n).Error
	if err != nil {
		return false
	}

	return true
}
