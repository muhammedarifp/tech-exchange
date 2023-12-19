package interfaces

import (
	"context"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
)

type AdminPaymentRepo interface {
	AddPlan(ctx context.Context, plan map[string]interface{}) (response.Plans, error)
	RemovePlan(ctx context.Context, planid string) (response.Plans, error)
}
