package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedarifp/tech-exchange/payments/usecase/interfaces"
)

type AdminPaymentHandler struct {
	usecase interfaces.AdminPaymentUsecase
}

func NewAdminPaymentHandler(usecase interfaces.AdminPaymentUsecase) *AdminPaymentHandler {
	return &AdminPaymentHandler{
		usecase: usecase,
	}
}

func (a *AdminPaymentHandler) AddPlan(c *gin.Context)    {}
func (a *AdminPaymentHandler) RemovePlan(c *gin.Context) {}
