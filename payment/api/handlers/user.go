package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
)

type UserPaymentHandler struct {
	usecase interfaces.UserPaymentUsecase
}

func NewUserPaymentHandler(usecase interfaces.UserPaymentUsecase) *UserPaymentHandler {
	return &UserPaymentHandler{
		usecase: usecase,
	}
}

func (a *UserPaymentHandler) FetchPlans(c *gin.Context) {}
func (a *UserPaymentHandler) CreateSubscription(c *gin.Context) {
	a.usecase.CreateSubscription()
}
func (a *UserPaymentHandler) CancelSubscription(c *gin.Context) {}
func (a *UserPaymentHandler) ChangePlan(c *gin.Context)         {}
