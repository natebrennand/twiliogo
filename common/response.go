package common

import (
	"strconv"
	"time"
)

const (
	TwilioTimeFormat = time.RFC1123Z
	GMTTimeLayout    = "2006-01-02" // YYYY-MM-DD
)

type JSONFloat float64

func (j *JSONFloat) UnmarshalJSON(b []byte) error {
	t, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return nil
	}
	*j = JSONFloat(t)
	return err
}

type JSONTime struct {
	time.Time
}

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
