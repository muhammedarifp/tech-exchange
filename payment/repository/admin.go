package repository

import (
	"context"

	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	"gorm.io/gorm"
)

type adminPaymentDb struct {
	db *gorm.DB
}

func NewAdminPaymentDb(db *gorm.DB) interfaces.AdminPaymentRepo {
	return &adminPaymentDb{
		db: db,
	}
}

func (d *adminPaymentDb) AddPlan(ctx context.Context)    {}
func (d *adminPaymentDb) RemovePlan(ctx context.Context) {}
