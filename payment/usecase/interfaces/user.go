package interfaces

import (
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/request"
	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
)

type UserPaymentUsecase interface {
	FetchAllPlans()
	CreateSubscription(userid uint, request request.CreateSubscriptionReq) (response.Subscription, error)
	CancelSubscription()
	ChangePlan()
	CreatePaymentAccount(userid uint)
}
