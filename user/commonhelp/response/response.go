package response

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}

type LoginResponse struct {
	StatusCode int         `json:"stastuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
	Token      interface{} `json:"token"`
}

type ReqOtpResp struct {
	StatusCode int         `json:"stastuscode"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}

type VerifyOtpResp struct {
	StatusCode int         `json:"stastuscode"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}
