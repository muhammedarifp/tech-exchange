package interfaces

import commonhelp "github.com/muhammedarifp/tech-exchange/notification/commonHelp"

type NotificationUsecase interface {
	StoreNotificationsOnDB()
	GetAllNotifications(token string) ([]commonhelp.NotificationResp, error)
}
