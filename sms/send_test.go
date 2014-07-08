package sms

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	validEndpoint   = "/valid"
	errorEndpoint   = "/error"
	badJsonEndpoint = "/badJson"
)

func TestValidateSmsPostSuccess(t *testing.T) {
	p := Post{From: testNumber1, To: testNumber2, Body: "test"}
	if nil != validateSmsPost(p) {
		t.Error("Validation of valid sms post failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, MediaUrl: "https://www.twilio.com/"}
	if nil != validateSmsPost(p) {
		t.Error("Validation of valid sms post failed.")
	}
}

func TestValidateSmsPostFailure(t *testing.T) {
	p := Post{}
	if nil == validateSmsPost(p) {
		t.Error("Validation of sms post missing To & From failed.")
	}

	p = Post{From: testNumber1}
	if nil == validateSmsPost(p) {
		t.Error("Validation of sms post missing From failed.")
	}

	p = Post{From: testNumber1, To: testNumber2}
	if nil == validateSmsPost(p) {
		t.Error("Validation of sms post missing Body & MediaUrl failed.")
	}
}

func startMockHttpServer(requests *int) *httptest.Server {
	// start a server to recieve post request
	testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		*requests += 1
		if strings.Contains(r.URL.Path, validEndpoint) {
			resp.WriteHeader(201)
			fmt.Fprint(resp, testSmsResponseFixtureString)
		} else if strings.Contains(r.URL.Path, errorEndpoint) {
			resp.WriteHeader(400)
		} else if strings.Contains(r.URL.Path, badJsonEndpoint) {
			fmt.Fprint(resp, testSmsResponseFixtureString[0:20])
		}
	}))
	return testServer
}

func TestSendSmsSuccess(t *testing.T) {
	act := SmsAccount{"act", "token"}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHttpServer(&numRequests)
	defer testPostServer.Close()

	var r Response
	err := act.sendSms(testPostServer.URL+validEndpoint, testSmsPostFixture, &r)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	if numRequests != 1 {
		t.Error("Server never recieved a request.")
	}
	if r.AccountSid != testSmsResponseFixtureAccountSid {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}

func TestSendSmsFailure(t *testing.T) {
	act := SmsAccount{"act", "token"}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHttpServer(&numRequests)
	defer testPostServer.Close()

	var r Response
	err := act.sendSms(testPostServer.URL+errorEndpoint, testSmsPostFixture, &r)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 1 {
		t.Error("server never recieved a request.")
	}

	err = act.sendSms(testPostServer.URL+badJsonEndpoint, testSmsPostFixture, &r)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}
}
