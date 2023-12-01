package usecase

import (
	interfaces "github.com/muhammedarifp/tech-exchange/notification/repository/interface"
	usercaseInterfaces "github.com/muhammedarifp/tech-exchange/notification/usecase/interfaces"
)

type notificationUsecase struct {
	repo interfaces.NotificationRepo
}

func NewNotificationUseCase(repo interfaces.NotificationRepo) usercaseInterfaces.NotificationUsecase {
	return &notificationUsecase{
		repo: repo,
	}
}

// func (u *notificationUsecase) GetallNotifications() {
// 	fmt.Println("Okkkk")
// }
