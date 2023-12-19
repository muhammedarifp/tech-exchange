package request

// admin create plan
type Plans struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Interval    int     `json:"interval"`
	Period      string  `json:"period"`
	Amount      float64 `json:"amount"`
}

type CreateSubscriptionReq struct {
	PlanID string `json:"plan"`
}
