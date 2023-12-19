package interfaces

import (
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
)

type AdminPaymentUsecase interface {
	AddPlan(plan request.Plans) (response.Plans, error)
	RemovePlan(planid string) (response.Plans, error)
}
