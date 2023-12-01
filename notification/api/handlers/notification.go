package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/muhammedarifp/tech-exchange/notification/cmd/docs"
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
// @Accept json
// @Produce json
// @Param page body int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of items to return per page (default is 10)"
// @Param sortBy query string false "Sort field for results"
// @Param sortOrder query string false "Sort order for results (asc or desc)"
// @Success 200 {array} domain.Notifications
// @Router /notifications [get]
func (u *NotificationsHandler) GetallNotifications(ctx *gin.Context) {
	ctx.String(200, "Iam Okkkkkkkkkkkkkkk!")
}

// Store notification on database
func (u *NotificationsHandler) StoreNotificationsOnDB() {

}
