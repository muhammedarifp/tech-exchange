package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
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

// Fetch all plans
func (d *userPaymentDb) FetchAllPlans(ctx context.Context) {}

// Create subscription
func (d *userPaymentDb) CreateSubscription(ctx context.Context, subsc map[string]interface{}) (response.Subscription, error) {
	// qury := `INSERT INTO subscriptions (subscription_id,customer_id,plan_id,status,starting_date,next_billing_date) VALUES ($1,$2,$3,$4,$5,$6)`
	// d.db.Raw(qury, "", "", "", "active", time.Now(), time.Now().AddDate(0, 0, 30))
	fmt.Println(subsc)
	return response.Subscription{}, nil
}

// Cancel
func (d *userPaymentDb) CancelSubscription(ctx context.Context) {}

// Change Plan
func (d *userPaymentDb) ChangePlan(ctx context.Context) {}

func (d *userPaymentDb) CreateRazorpayAccount(ctxrazorpay_id context.Context, userid uint, account map[string]interface{}) {
	qury := `INSERT INTO razorpay_accounts(user_id,razorpay_id,email) VALUES ($1,$2,$3)`
	var a interface{}
	err := d.db.Raw(qury, userid, account["id"].(string), account["email"].(string)).Scan(&a).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
}

func (d *userPaymentDb) FetchRazorpayAccount(userid uint) (response.Account, error) {
	qury := `SELECT * FROM razorpay_accounts WHERE user_id = ?`
	accountData := response.Account{}
	if err := d.db.Raw(qury, userid).Scan(&accountData).Error; err != nil {
		return response.Account{}, err
	}

	return accountData, nil
}

// func (d *userPaymentDb) CreatePaymentAccount(userid, costemerid string, msgs msgs.PaymentAccount) bool {
// 	qury := `SELECT * FROM razorpay_accounts (user_id,razorpay_id,email) VALUES ($1,$2,$3,$4)`
// 	if err := d.db.Raw(qury, userid, costemerid, msgs.Email).Error; err != nil {
// 		return false
// 	}

// 	return true
// }
