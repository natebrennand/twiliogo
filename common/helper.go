package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Base URL for API requests. You probably only want to change this if you're writing tests using
// a mock http server.
var BaseURL = "https://api.twilio.com"

type twilioAccount interface {
	GetSid() string
	GetToken() string
	GetClient() http.Client
}

type twilioResponse interface{} // Empty interface

type twilioPost interface {
	GetReader() io.Reader
	Validate() error
}

func buildResp(resp *twilioResponse, httpResp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading json from buffer => %s", err.Error())
	}
	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes))
	}
	return nil
}

func successfulResponse(statusCode int) bool {
	if statusCode >= 200 && statusCode < 300 {
		return true
	}
	return false
}

// SendPostRequest sends an authenticated POST request to Twilio with the encoded data.
func SendPostRequest(url string, msg twilioPost, t twilioAccount, resp twilioResponse) error {
	if nil != msg.Validate() {
		return fmt.Errorf("Error validating post => %s.\n", msg.Validate().Error())
	}

	req, err := http.NewRequest("POST", BaseURL+url, msg.GetReader())
	req.SetBasicAuth(t.GetSid(), t.GetToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := t.GetClient()
	httpResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending req => %s", err.Error())
	}
	if !successfulResponse(httpResp.StatusCode) {
		return NewTwilioError(*httpResp)
	}

	return buildResp(&resp, httpResp)
}

// SendGetRequest sends an authenticated GET request to Twilio
func SendGetRequest(url string, t twilioAccount, resp twilioResponse) error {
	req, err := http.NewRequest("GET", BaseURL+url, nil)
	req.SetBasicAuth(t.GetSid(), t.GetToken())
	req.Header.Add("Accept", "application/json")

	client := t.GetClient()
	httpResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending req => %s", err.Error())
	}
	if !successfulResponse(httpResp.StatusCode) {
		return NewTwilioError(*httpResp)
	}

	return buildResp(&resp, httpResp)
}

// SendDeleteRequest sends an authenticated DELETE request to Twilio
func SendDeleteRequest(url string, t twilioAccount) error {
	req, err := http.NewRequest("DELETE", BaseURL+url, nil)
	req.SetBasicAuth(t.GetSid(), t.GetToken())
	req.Header.Add("Accept", "application/json")

	client := t.GetClient()
	httpResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending req => %s", err.Error())
	}
	if !successfulResponse(httpResp.StatusCode) {
		return NewTwilioError(*httpResp)
	}
	return nil
}
