package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	validEndpoint200 = "/valid200"
	validEndpoint201 = "/valid201"
	errorEndpoint    = "/error"
	badJSONEndpoint  = "/badJSON"
)

func startMockHTTPServer(requests *int) *httptest.Server {
	// start a server to recieve post request
	baseURL = "" // reset the base url of all queries
	testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		*requests++
		if strings.Contains(r.URL.Path, validEndpoint200) {
			resp.WriteHeader(200)
			fmt.Fprint(resp, testMessageFixtureString)
		} else if strings.Contains(r.URL.Path, validEndpoint201) {
			resp.WriteHeader(201)
			fmt.Fprint(resp, testMessageFixtureString)
		} else if strings.Contains(r.URL.Path, errorEndpoint) {
			resp.WriteHeader(400)
		} else if strings.Contains(r.URL.Path, badJSONEndpoint) {
			fmt.Fprint(resp, testSmsResponseFixtureString[0:20])
		}
	}))
	return testServer
}

func TestSendSuccess(t *testing.T) {
	act := testAccount{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHTTPServer(&numRequests)
	defer testPostServer.Close()

	var m testMessage
	err := act.sendSms(testPostServer.URL+validEndpoint201, testPostFixture, &m)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	if numRequests != 1 {
		t.Error("Server never recieved a request.")
	}
	if m.Foo != "Bar" {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}

func TestSendFailure(t *testing.T) {
	act := testAccount{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHTTPServer(&numRequests)
	defer testPostServer.Close()

	var m testMessage
	err := act.sendSms(testPostServer.URL+badJSONEndpoint, testPostFixtureInvalid, &m)
	if err == nil {
		t.Errorf("post should've failed with an invalid post")
	}
	if numRequests != 0 {
		t.Error("server recieved a request.")
	}

	err = act.sendSms(testPostServer.URL+errorEndpoint, testPostFixture, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 1 {
		t.Error("server never recieved a request.")
	}

	err = act.sendSms(testPostServer.URL+badJSONEndpoint, testPostFixture, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}
}

func TestGetSuccess(t *testing.T) {
	act := testAccount{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testGetServer := startMockHTTPServer(&numRequests)
	defer testGetServer.Close()

	var m testMessage
	err := act.getSms(testGetServer.URL+validEndpoint200, &m)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	if numRequests != 1 {
		t.Error("Server never recieved a request.")
	}
	if m.Foo != "Bar" {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}

func TestGetFailure(t *testing.T) {
	act := testAccount{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testGetServer := startMockHTTPServer(&numRequests)
	defer testGetServer.Close()

	var m testMessage
	err := act.getSms(testGetServer.URL+errorEndpoint, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 1 {
		t.Error("server never recieved a request.")
	}

	err = act.getSms(testGetServer.URL+badJSONEndpoint, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}
}
