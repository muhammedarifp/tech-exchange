package interfaces

import "context"

type AdminPaymentRepo interface {
	AddPlan(ctx context.Context, plan map[string]interface{})
	RemovePlan(ctx context.Context)
}
