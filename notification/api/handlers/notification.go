package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/content/commonHelp/response"
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
	token := ctx.GetHeader("Token")
	if token == "" {
		ctx.JSON(400, response.Response{
			StatusCode: 404,
			Message:    "Token not found. Please login first to get token and then try again.",
			Data:       nil,
			Errors:     "Token not found",
		})
		return
	}

	resp, err := u.usecase.GetAllNotifications(token)
	if err != nil {
		ctx.JSON(400, response.Response{
			StatusCode: 400,
			Message:    "Something went wrong",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(200, response.Response{
		StatusCode: 200,
		Message:    "Success",
		Data:       resp,
		Errors:     nil,
	})
}

func (u *NotificationsHandler) StoreNotificationOnDatabase() {
	u.usecase.StoreNotificationsOnDB()
}
