package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/natebrennand/twilio-go/common"
)

type SmsResponseJson struct {
	common.ResponseCore
	JsonNumSegments string `json:"num_segments"`
	JsonNumMedia    string `json:"num_media"`
	JsonDateCreated string `json:"date_created"`
	JsonDateSent    string `json:"date_sent"`
	JsonPrice       string `json:"price"`
}

type Response struct {
	common.ResponseCore
	NumSegments int
	NumMedia    int
	Price       float64
	DateCreated time.Time
	DateSent    time.Time
}

// Unmarshals a twilio sms response into a Response struct.
func Unmarshal(data []byte, msg *Response) error {
	var msgJson SmsResponseJson
	err := json.Unmarshal(data, &msgJson)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s", err.Error()))
	}

	var (
		numSegments, numMedia int
		price                 float64
		dateCreated, dateSent time.Time
	)

	numSegments, err = strconv.Atoi(msgJson.JsonNumSegments)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while converting num_segments to an integer => %s", err.Error()))
	}

	numMedia, err = strconv.Atoi(msgJson.JsonNumMedia)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while converting num_media to an integer => %s", err.Error()))
	}

	if msgJson.JsonPrice != "" {
		price, err = strconv.ParseFloat(msgJson.JsonPrice, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while parsing price => %s", err.Error()))
		}
	}

	dateCreated, err = time.Parse(twilioTimeFormat, msgJson.JsonDateCreated)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while parsing date_created => %s", err.Error()))
	}

	if msgJson.JsonDateSent != "" { // date sent is not always instantiated
		dateSent, err = time.Parse(twilioTimeFormat, msgJson.JsonDateSent)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while parsing date_sent => %s", err.Error()))
		}
	}

	*msg = Response{
		ResponseCore: msgJson.ResponseCore,
		NumSegments:  numSegments,
		NumMedia:     numMedia,
		Price:        price,
		DateCreated:  dateCreated,
		DateSent:     dateSent,
	}
	return nil
}
