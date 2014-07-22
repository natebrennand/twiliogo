package common

import (
	"strconv"
	"time"
)

// These time formats are used throughout the Twilio API.
const (
	TwilioTimeFormat = time.RFC1123Z
	GMTTimeLayout    = "2006-01-02" // YYYY-MM-DD
)

// JSONFloat is a wrapper around float for json encoding
type JSONFloat float64

// UnmarshalJSON unmarshals json JSONFloat objects
func (j *JSONFloat) UnmarshalJSON(b []byte) error {
	t, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return nil
	}
	*j = JSONFloat(t)
	return err
}

// JSONTime is a wrapper for parsing time into a time.Time object at the point of
// unmarshaling
type JSONTime struct {
	time.Time
}

// UnmarshalJSON unmarshals the json string into a time.Time object.
func (j *JSONTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		*j = JSONTime{Time: time.Time{}}
		return nil
	}
	t, err := time.Parse(TwilioTimeFormat, s[1:len(s)-1])
	*j = JSONTime{Time: t}
	return err
}

// ResponseCore is a set of attributes found in many Twilio resource responses.
type ResponseCore struct {
	AccountSid   string   `json:"account_sid"`
	APIVersion   string   `json:"api_version"`
	Direction    string   `json:"direction"`
	ErrorCode    string   `json:"error_code"`
	ErrorMessage string   `json:"error_message"`
	From         string   `json:"from"`
	Sid          string   `json:"sid"`
	Status       string   `json:"status"`
	To           string   `json:"to"`
	URI          string   `json:"uri"`
	DateCreated  JSONTime `json:"date_created"`
	DateUpdated  JSONTime `json:"date_updated"`
}

// ListResponseCore is a set of attributes found all Twilio list resources.
type ListResponseCore struct {
	Start           int    `json:"start"`
	Total           int    `json:"total"`
	NumPages        int    `json:"num_pages"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
	End             int    `json:"end"`
	URI             string `json:"uri"`
	FirstPageURI    string `json:"first_page_uri"`
	LastPageURI     string `json:"last_page_uri"`
	NextPageURI     string `json:"next_page_uri"`
	PreviousPageURI string `json:"previous_page_uri"`
}

// TODO: rename
type ResponseCore2 struct {
	AccountSid  string   `json:"account_sid"`
	Sid         string   `json:"sid"`
	URI         string   `json:"uri"`
	DateCreated JSONTime `json:"date_created"`
	DateUpdated JSONTime `json:"date_updated"`
}
