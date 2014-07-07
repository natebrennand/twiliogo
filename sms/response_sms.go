package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type responseCore struct {
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
	responseCore
	JsonNumSegments string `json:"num_segments"`
	JsonNumMedia    string `json:"num_media"`
	JsonDateCreated string `json:"date_created"`
	JsonDateSent    string `json:"date_sent"`
	JsonPrice       string `json:"price"`
}

type Response struct {
	responseCore
	NumSegments int
	NumMedia    int
	Price       float64
	DateCreated time.Time
	DateSent    time.Time
}

func Unmarshal(data []byte, msg *Response) error {
	var msgJson SmsResponseJson
	err := json.Unmarshal(data, &msgJson)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s", err.Error()))
	}

	// copy in sms core to struct
	msg.responseCore = msgJson.responseCore

	msg.NumSegments, err = strconv.Atoi(msgJson.JsonNumSegments)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while converting num_segments to an integer => %s", err.Error()))
	}

	msg.NumMedia, err = strconv.Atoi(msgJson.JsonNumMedia)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while converting num_media to an integer => %s", err.Error()))
	}

	if msgJson.JsonPrice != "" {
		msg.Price, err = strconv.ParseFloat(msgJson.JsonPrice, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while parsing price => %s", err.Error()))
		}
	}

	msg.DateCreated, err = time.Parse(twilioTimeFormat, msgJson.JsonDateCreated)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while parsing date_created => %s", err.Error()))
	}

	if msgJson.JsonDateSent != "" { // date sent is not always instantiated
		msg.DateSent, err = time.Parse(twilioTimeFormat, msgJson.JsonDateSent)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while parsing date_sent => %s", err.Error()))
		}
	}
	return nil
}
