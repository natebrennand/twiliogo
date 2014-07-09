package common

import (
	"strconv"
	"time"
)

const (
	TwilioTimeFormat = time.RFC1123Z
)

type JsonPrice float64

func (j *JsonPrice) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*j = JsonPrice(0.0)
		return nil
	}
	t, err := strconv.ParseFloat(s, 64)
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
	Body         string   `json:"body"`
	Direction    string   `json:"direction"`
	ErrorCode    string   `json:"error_code"`
	ErrorMessage string   `json:"error_message"`
	From         string   `json:"from"`
	Sid          string   `json:"sid"`
	Status       string   `json:"status"`
	To           string   `json:"to"`
	Uri          string   `json:"uri"`
	DateCreated  JsonTime `json:"date_created"`
	DateSent     JsonTime `json:"date_sent"`
	DateUpdated  JsonTime `json:"date_updated"`
}

type Media struct {
	ContentType string
	Url         string
}

type Location struct {
	FromCity    string
	FromState   string
	FromZip     string
	FromCountry string
	ToCity      string
	ToState     string
	ToZip       string
	ToCountry   string
}

// Standard request parameters for Twiml responses
// https://www.twilio.com/docs/api/twiml/sms/twilio_request#request-parameters
type StandardRequest struct {
	AccountSid string
	From       string
	To         string
	MediaList  []Media
	Location   *Location // Only sent when Twilio can look up the geographic data.
}
