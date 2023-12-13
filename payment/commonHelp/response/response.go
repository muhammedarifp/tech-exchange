package response

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
	SubscriptionID string `json:"subscription_id"`
	CostemerID     string `json:"customer_id"`
	PlanID         string `json:"plan_id"`
	Status         string `json:"status"`
	StartingDate   string `json:"starting_date"`
	NextDate       string `json:"next_billing_date"`
}
