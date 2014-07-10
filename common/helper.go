package common

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TwilioAccount interface {
	GetSid() string
	GetToken() string
	GetClient() http.Client
}

type TwilioResponse interface {
	Build(*http.Response) error
}

type TwilioPost interface {
	GetReader() io.Reader
	Validate() error
}

func SendPostRequest(url string, msg TwilioPost, t TwilioAccount, resp TwilioResponse, expectedResponse int) error {
	if nil != msg.Validate() {
		return errors.New(fmt.Sprintf("Error validating sms post => %s.\n", msg.Validate().Error()))
	}

	req, err := http.NewRequest("POST", url, msg.GetReader())
	req.SetBasicAuth(t.GetSid(), t.GetToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := t.GetClient()
	httpResp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Error sending req => %s", err.Error()))
	}
	if httpResp.StatusCode != expectedResponse {
		return NewTwilioError(*httpResp)
	}

	// build response
	return resp.Build(httpResp)
}

func SendGetRequest(url string, t TwilioAccount, resp TwilioResponse, expectedResponse int) error {
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(t.GetSid(), t.GetToken())
	req.Header.Add("Accept", "application/json")

	client := t.GetClient()
	httpResp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Error sending req => %s", err.Error()))
	}
	if httpResp.StatusCode != expectedResponse {
		return NewTwilioError(*httpResp)
	}

	// build response
	return resp.Build(httpResp)
}
