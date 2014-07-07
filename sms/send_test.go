package sms

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testSendUrl = "/sms/send"
)

func TestSendSms(t *testing.T) {
	// start a server to recieve post request
	serverRequested := false
	testPostServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		serverRequested = true
		fmt.Fprint(resp, testSmsResponseFixture)
	}))
	defer testPostServer.Close()

	var r Response
	jReader := bytes.NewBuffer([]byte(testSmsPostFixture))
	err := sendSms(testPostServer.URL, jReader, &r)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	if serverRequested != true {
		t.Error("Server never recieved a request.")
	}
	if r.AccountSid != testSmsResponseFixtureAccountSid {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}
