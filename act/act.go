package act

import (
	"net/http"
)

// Account represents a Twilio account
type Account struct {
	AccountSid string
	Token      string
	Client     http.Client
}

// GetSid teturns the account sid
func (act Account) GetSid() string {
	return act.AccountSid
}

// GetToken returns the secret token for an account
func (act Account) GetToken() string {
	return act.Token
}

// GetClient returns a shared http client
func (act Account) GetClient() http.Client {
	return act.Client
}
