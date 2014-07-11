package sms

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const (
	validEndpoint   = "/valid"
	listEndpoint    = "/list"
	errorEndpoint   = "/error"
	badJSONEndpoint = "/badJSON"
)

func TestValidateSmsPostSuccess(t *testing.T) {
	p := Post{From: testNumber1, To: testNumber2, Body: "test"}
	if nil != p.Validate() {
		t.Error("Validation of valid sms post failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, MediaURL: "https://www.twilio.com/"}
	if nil != p.Validate() {
		t.Error("Validation of valid sms post failed.")
	}
}

func TestValidateSmsPostFailure(t *testing.T) {
	p := Post{}
	if nil == p.Validate() {
		t.Error("Validation of sms post missing To & From failed.")
	}

	p = Post{From: testNumber1}
	if nil == p.Validate() {
		t.Error("Validation of sms post missing From failed.")
	}

	p = Post{From: testNumber1, To: testNumber2}
	if nil == p.Validate() {
		t.Error("Validation of sms post missing Body & MediaURL failed.")
	}
}

func startMockHTTPServer(requests *int) *httptest.Server {
	// start a server to recieve post request
	testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		*requests++
		if strings.Contains(r.URL.Path, validEndpoint) {
			if strings.Contains(r.URL.Path, listEndpoint) {
				resp.WriteHeader(http.StatusOK)
				fmt.Fprint(resp, testSmsListFixtureString)
			} else {
				resp.WriteHeader(http.StatusCreated)
				fmt.Fprint(resp, testSmsResponseFixtureString)
			}
		} else if strings.Contains(r.URL.Path, errorEndpoint) {
			resp.WriteHeader(http.StatusBadRequest)
		} else if strings.Contains(r.URL.Path, badJSONEndpoint) {
			fmt.Fprint(resp, testSmsResponseFixtureString[0:20])
		}
	}))
	return testServer
}

func TestSendSmsSuccess(t *testing.T) {
	act := Account{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHTTPServer(&numRequests)
	defer testPostServer.Close()

	var m Message
	err := act.sendSms(testPostServer.URL+validEndpoint, testSmsPostFixture, &m)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	if numRequests != 1 {
		t.Error("Server never recieved a request.")
	}
	if m.AccountSid != testSmsResponseFixtureAccountSid {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}

func TestSendSmsFailure(t *testing.T) {
	act := Account{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHTTPServer(&numRequests)
	defer testPostServer.Close()

	var m Message
	err := act.sendSms(testPostServer.URL+errorEndpoint, testSmsPostFixture, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 1 {
		t.Error("server never recieved a request.")
	}

	err = act.sendSms(testPostServer.URL+badJSONEndpoint, testSmsPostFixture, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}
}

func TestListFilterReader(t *testing.T) {
	// empty filter
	f := Filter{}
	if "" != f.GetQueryString() {
		t.Error("url encoding of filter should be empty")
	}

	// w/ To
	f = Filter{To: "12345"}
	if f.GetQueryString() != "?To=12345" {
		t.Errorf("url encoding of filter should include To&From key:value pairs, found => %s", f.GetQueryString())
	}

	// w/ To & From
	f = Filter{To: "12345", From: "678"}
	if !strings.Contains(f.GetQueryString(), "To=12345") || !strings.Contains(f.GetQueryString(), "From=678") {
		t.Errorf("url encoding of filter should include To&From key:value pairs, found => %s", f.GetQueryString())
	}

	// w/ Date
	tm := time.Date(2010, time.August, 16, 3, 45, 01, 0, &time.Location{})
	f = Filter{DateSent: &tm}
	if "?DateSent=2010-08-16" != f.GetQueryString() {
		t.Errorf("url encoding of filter should encode dates in GMT, found => %s", f.GetQueryString())
	}

}

func TestGetSmsListSuccess(t *testing.T) {
	act := Account{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testServer := startMockHTTPServer(&numRequests)
	defer testServer.Close()

	var ml MessageList
	err := act.getList(testServer.URL+validEndpoint+listEndpoint, Filter{}, &ml)
	if err != nil {
		t.Errorf("Error while sending get request => %s", err.Error())
	}
	if numRequests != 1 {
		t.Error("Server never recieved a request.")
	}
	if ml.Total != testSmsListFixture.Total {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}

func TestSendSmsFailure2(t *testing.T) {
	act := Account{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testServer := startMockHTTPServer(&numRequests)
	defer testServer.Close()

	var ml MessageList
	err := act.getList(testServer.URL+errorEndpoint, Filter{}, &ml)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 1 {
		t.Error("server never recieved a request.")
	}

	err = act.getList(testServer.URL+badJSONEndpoint, Filter{}, &ml)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}
}
