package interfaces

import "context"

type AdminPaymentRepo interface {
	AddPlan(ctx context.Context)
	RemovePlan(ctx context.Context)
}
