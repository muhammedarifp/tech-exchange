package interfaces

import "context"

type UserPaymentRepo interface {
	FetchAllPlans(ctx context.Context)
	CreateSubscription(ctx context.Context)
	CancelSubscription(ctx context.Context)
	ChangePlan(ctx context.Context)
}
