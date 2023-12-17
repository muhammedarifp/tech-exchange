package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/muhammedarifp/tech-exchange/payments/commonHelp/response"
	"github.com/muhammedarifp/tech-exchange/payments/repository/interfaces"
	"gorm.io/gorm"
)

type userPaymentDb struct {
	db *gorm.DB
}

var (
	databaseb *gorm.DB
)

func NewUserPaymentRepo(db *gorm.DB) interfaces.UserPaymentRepo {
	databaseb = db
	return &userPaymentDb{
		db: db,
	}
}

// Fetch all plans
func (d *userPaymentDb) FetchAllPlans(ctx context.Context) {}

// Create subscription
func (d *userPaymentDb) CreateSubscription(ctx context.Context, subsc map[string]interface{}) (response.Subscription, error) {
	qury := `INSERT INTO subscriptions (subscription_id,customer_id,plan_id,status,starting_date,next_billing_date) VALUES ($1,$2,$3,$4,$5,$6)`
	var resp response.Subscription
	if err := d.db.Raw(qury, subsc["id"].(string), subsc["customer_id"].(string), subsc["plan_id"].(string), subsc["status"].(string), time.Now(), time.Now().AddDate(0, 0, 30)).Scan(resp).Error; err != nil {
		return response.Subscription{}, err
	}
	return resp, nil
}

// Cancel
func (d *userPaymentDb) CancelSubscription(ctx context.Context, subid string) (response.Subscription, error) {
	qury := `UPDATE subscriptions SET status = 'cancelled' WHERE subscription_id = $1`
	var subData response.Subscription
	if err := d.db.Raw(qury, subid).Scan(&subData).Error; err != nil {
		return response.Subscription{}, err
	}

	return subData, nil
}

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
	useridStr := strconv.Itoa(int(userid))
	qury := `SELECT * FROM razorpay_accounts WHERE user_id = ?`
	accountData := response.Account{}
	if err := d.db.Raw(qury, useridStr).Scan(&accountData).Error; err != nil {
		return response.Account{}, err
	}

	return accountData, nil
}
