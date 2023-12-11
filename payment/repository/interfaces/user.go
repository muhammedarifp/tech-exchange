package interfaces

import "context"

type UserPaymentRepo interface {
	FetchAllPlans(ctx context.Context)
	CreateSubscription(ctx context.Context, subsc map[string]interface{})
	CancelSubscription(ctx context.Context)
	ChangePlan(ctx context.Context)
	CreateRazorpayAccount(ctx context.Context, userid uint, account map[string]interface{})
}

//
