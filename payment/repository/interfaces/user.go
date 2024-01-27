package interfaces

import (
	"context"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
)

type UserPaymentRepo interface {
	FetchAllPlans(ctx context.Context, page int) ([]response.Plans, error)
	CreateSubscription(ctx context.Context, subsc map[string]interface{}) (response.Subscription, error)
	CancelSubscription(ctx context.Context, subid string) (response.Subscription, error)
	ChangePlan(ctx context.Context)
	CreateRazorpayAccount(ctx context.Context, userid uint, account map[string]interface{})
	//FetchRazorpayAccount(userid uint) (response.Account, error)
}

//
