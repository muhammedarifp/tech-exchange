package repository

import (
	"context"

	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	"gorm.io/gorm"
)

type userPaymentDb struct {
	db *gorm.DB
}

func NewUserPaymentRepo(db *gorm.DB) interfaces.UserPaymentRepo {
	return &userPaymentDb{
		db: db,
	}
}

func (d *userPaymentDb) FetchAllPlans(ctx context.Context)      {}
func (d *userPaymentDb) CreateSubscription(ctx context.Context) {}
func (d *userPaymentDb) CancelSubscription(ctx context.Context) {}
func (d *userPaymentDb) ChangePlan(ctx context.Context)         {}
