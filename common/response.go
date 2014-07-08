package common

type ResponseCore struct {
	AccountSid   string `json:"account_sid"`
	ApiVersion   string `json:"api_version"`
	Body         string `json:"body"`
	Direction    string `json:"direction"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	From         string `json:"from"`
	Sid          string `json:"sid"`
	Status       string `json:"status"`
	To           string `json:"to"`
	Uri          string `json:"uri"`
}
