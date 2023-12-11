package repository

import (
	"context"
	"fmt"
	"log"

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

func (d *userPaymentDb) FetchAllPlans(ctx context.Context) {}
func (d *userPaymentDb) CreateSubscription(ctx context.Context, subsc map[string]interface{}) {
	fmt.Println(subsc)
}
func (d *userPaymentDb) CancelSubscription(ctx context.Context) {}
func (d *userPaymentDb) ChangePlan(ctx context.Context)         {}

func (d *userPaymentDb) CreateRazorpayAccount(ctxrazorpay_id context.Context, userid uint, account map[string]interface{}) {
	qury := `INSERT INTO razorpay_accounts(user_id,razorpay_id,phone,email) VALUES ($1,$2,$3,$4)`
	var a interface{}
	err := d.db.Raw(qury, userid, account["id"].(string), account["contact"].(string), account["email"].(string)).Scan(&a).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
}
