package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const twilioTimeFormat = time.RFC1123Z

type SmsResponseCore struct {
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

type SmsResponseJson struct {
	SmsResponseCore
	JsonNumSegments string `json:"num_segments"`
	JsonNumMedia    string `json:"num_media"`
	JsonDateCreated string `json:"date_created"`
	JsonDateSent    string `json:"date_sent"`
	JsonPrice       string `json:"price"`
}

type SmsResponse struct {
	SmsResponseCore
	NumSegments int
	NumMedia    int
	Price       float32
	DateCreated time.Time
	DateSent    time.Time
}

func Unmarshal(data []byte, msg *SmsResponse) error {

	var msgJson SmsResponseJson
	err := json.Unmarshal(data, &msgJson)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s", err.Error()))
	}

	// copy in sms core to struct
	msg.SmsResponseCore = msgJson.SmsResponseCore

	msg.NumSegments, err = strconv.Atoi(msgJson.JsonNumSegments)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while converting num_segments to an integer => %s", err.Error()))
	}

	msg.NumMedia, err = strconv.Atoi(msgJson.JsonNumMedia)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while converting num_media to an integer => %s", err.Error()))
	}

	msg.DateCreated, err = time.Parse(twilioTimeFormat, msgJson.JsonDateCreated)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while parsing date_created => %s", err.Error()))
	}

	if msgJson.JsonDateSent != "" {
		msg.DateSent, err = time.Parse(twilioTimeFormat, msgJson.JsonDateSent)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while parsing date_sent => %s", err.Error()))
		}
	}

	return nil
}
