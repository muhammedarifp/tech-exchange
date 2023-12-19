package msgs

type PaymentAccount struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

type Customer struct {
	Contact   interface{} `json:"contact"`
	CreatedAt float64     `json:"created_at"`
	Email     string      `json:"email"`
	Entity    string      `json:"entity"`
	Gstin     interface{} `json:"gstin"`
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Notes     interface{} `json:"notes"`
}

// Collection represents the overall structure of the JSON data
type Collection struct {
	Count  int        `json:"count"`
	Entity string     `json:"entity"`
	Items  []Customer `json:"items"`
}
