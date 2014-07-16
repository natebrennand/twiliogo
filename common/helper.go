package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type TwilioAccount interface {
	GetSid() string
	GetToken() string
	GetClient() http.Client
}

type TwilioResponse interface{} // Empty interface

type TwilioPost interface {
	GetReader() io.Reader
	Validate() error
}

func buildResp(resp *TwilioResponse, httpResp *http.Response) error {
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

func SendPostRequest(url string, msg TwilioPost, t TwilioAccount, resp TwilioResponse) error {
	if nil != msg.Validate() {
		return fmt.Errorf("Error validating post => %s.\n", msg.Validate().Error())
	}

	req, err := http.NewRequest("POST", url, msg.GetReader())
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

func SendGetRequest(url string, t TwilioAccount, resp TwilioResponse) error {
	req, err := http.NewRequest("GET", url, nil)
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

func SendDeleteRequest(url string, t TwilioAccount) error {
	req, err := http.NewRequest("DELETE", url, nil)
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
