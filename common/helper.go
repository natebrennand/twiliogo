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

func FormNewPostRequest(url string, msg TwilioPost, t TwilioAccount, expectedResponse int) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, msg.GetReader())
	req.SetBasicAuth(t.GetSid(), t.GetToken())
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := t.GetClient()
	resp, err := client.Do(req)
	if err != nil {
		return resp, errors.New(fmt.Sprintf("Error sending req => %s", err.Error()))
	}
	if resp.StatusCode != expectedResponse {
		return resp, NewTwilioError(*resp)
	}
	return resp, err
}
