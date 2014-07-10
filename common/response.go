package common

import (
	// "fmt"
	"strconv"
	"time"
)

const (
	TwilioTimeFormat = time.RFC1123Z
)

type JsonPrice float64

func (j *JsonPrice) UnmarshalJSON(b []byte) error {
	t, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return nil
	}
	*j = JsonPrice(t)
	return err
}

type JsonTime struct {
	time.Time
}

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*j = JsonTime{time.Time{}}
		return nil
	}
	t, err := time.Parse(TwilioTimeFormat, s[1:len(s)-1])
	*j = JsonTime{t}
	return err
}

type ResponseCore struct {
	AccountSid   string   `json:"account_sid"`
	ApiVersion   string   `json:"api_version"`
	Body         string   `json:"body"` //why is this here
	Direction    string   `json:"direction"`
	ErrorCode    string   `json:"error_code"`
	ErrorMessage string   `json:"error_message"`
	From         string   `json:"from"`
	Sid          string   `json:"sid"`
	Status       string   `json:"status"`
	To           string   `json:"to"`
	Uri          string   `json:"uri"`
	DateCreated  JsonTime `json:"date_created"`
	DateSent     JsonTime `json:"date_sent"` //why is this here
	DateUpdated  JsonTime `json:"date_updated"`
}
