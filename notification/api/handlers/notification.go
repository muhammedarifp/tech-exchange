package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/muhammedarifp/tech-exchange/notification/docs"
	"github.com/muhammedarifp/tech-exchange/notification/usecase/interfaces"
)

type NotificationsHandler struct {
	usecase interfaces.NotificationUsecase
}

func NewNotificationHandler(usecase interfaces.NotificationUsecase) *NotificationsHandler {
	return &NotificationsHandler{
		usecase: usecase,
	}
}

// @Summary Get a list of notifications
// @Description Retrieves a list of notifications.
// @Tags notifications
// @Produce json
// @Success 200 {array} domain.Notifications
// @Router /notifications [get]
func (u *NotificationsHandler) GetallNotifications(ctx *gin.Context) {
	ctx.String(200, "Iam Okkkkkkkkkkkkkkk!")
}

// Store notification on database
func (u *NotificationsHandler) StoreNotificationsOnDB() {

}
