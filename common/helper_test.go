package common

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

func startMockHttpServer(requests *int) *httptest.Server {
	// start a server to recieve post request
	testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		*requests += 1
		if strings.Contains(r.URL.Path, validEndpoint) {
			resp.WriteHeader(201)
			fmt.Fprint(resp, testMessageFixtureString)
		} else if strings.Contains(r.URL.Path, errorEndpoint) {
			resp.WriteHeader(400)
		} else if strings.Contains(r.URL.Path, badJsonEndpoint) {
			fmt.Fprint(resp, testSmsResponseFixtureString[0:20])
		}
	}))
	return testServer
}

func TestSendPostRequest(t *testing.T) {
}

func TestSendSuccess(t *testing.T) {
	act := testAccount{"act", "token", http.Client{}}

	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHttpServer(&numRequests)
	defer testPostServer.Close()

	var m testMessage
	err := act.sendSms(testPostServer.URL+validEndpoint, testPostFixture, &m)
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
	testPostServer := startMockHttpServer(&numRequests)
	defer testPostServer.Close()

	var m testMessage
	err := act.sendSms(testPostServer.URL+badJsonEndpoint, testPostFixtureInvalid, &m)
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

	err = act.sendSms(testPostServer.URL+badJsonEndpoint, testPostFixture, &m)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}

}
