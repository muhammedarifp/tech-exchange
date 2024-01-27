package interfaces

import commonhelp "github.com/muhammedarifp/tech-exchange/notification/commonHelp"

type NotificationRepo interface {
	StoreNotificationsOnDB(notification commonhelp.NotificationResp) bool
	GetAllNotifications(userid string) ([]commonhelp.NotificationResp, error)
}
