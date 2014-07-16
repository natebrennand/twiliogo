package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func decodeError(err error, body []byte) error {
	return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(body))
}

// Error represents an error returned by the Twilio API
type Error struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

// Error generates a string representation of the Error.
func (e Error) Error() string {
	return fmt.Sprintf("Twilio Error %d => %s, more info @ %s", e.Code, e.Message, e.MoreInfo)
}

// NewTwilioError is the constructer for a Twilio error given an erroneous http response.
func NewTwilioError(resp http.Response) error {
	var twilioErr Error
	var buf bytes.Buffer
	_, err := buf.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf("Twilio error encountered, failure while reading body => %s", err.Error())
	}

	err = json.Unmarshal(buf.Bytes(), &twilioErr)
	if err != nil {
		return fmt.Errorf("Twilio error encountered, failure while parsing => %s", err.Error())
	}

	if twilioErr.Code == 0 {
		return errors.New("Twilio error not found, perhaps you set the expected status code incorrectly?")
	}

	return twilioErr
}
