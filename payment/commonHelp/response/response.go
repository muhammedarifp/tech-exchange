package response

import "time"

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}

type Account struct {
	UserID     uint   `json:"user_id"`
	RazorpayID string `json:"razorpay_id"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

type Subscription struct {
	SubscriptionID  string    `json:"subscription_id"`
	CustomerID      string    `json:"customer_id"`
	PlanID          string    `json:"plan_id"`
	Status          string    `json:"status"`
	StartingDate    time.Time `json:"starting_date"`
	NextBillingDate time.Time `json:"next_billing_date"`
}

// Plans Response
type Plans struct {
	ID          uint    `json:"id"`
	PlanID      string  `json:"plan_id" gorm:"plan_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Interval    int     `json:"interval"`
	Period      string  `json:"period"`
	Amount      float64 `json:"amount"`
	IsActive    bool    `json:"is_active"`
}
