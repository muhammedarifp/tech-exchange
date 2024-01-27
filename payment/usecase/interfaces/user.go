package interfaces

import (
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
)

type UserPaymentUsecase interface {
	FetchAllPlans(page int) ([]response.Plans, error)
	CreateSubscription(userid uint, request request.CreateSubscriptionReq) (response.Subscription, error)
	CancelSubscription(subid string) (response.Subscription, error)
	ChangePlan()
	VerifyPayment(payload, signature string) (bool, error)
}
