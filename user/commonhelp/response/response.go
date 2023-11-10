package response

type Response struct {
	StatusCode int         `json:"stastuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}

type LoginResponse struct {
	StatusCode int         `json:"stastuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
	Token      string      `json:"token"`
}
