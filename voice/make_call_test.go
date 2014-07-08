package voice

import (
	"bytes"
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

func TestValidatePostSuccess(t *testing.T) {
	p := Post{From: testNumber1, To: testNumber2, Url: "http://twimlbin.com/558a498f"}
	if nil != validatePost(p) {
		t.Error("Validation of valid voice post failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, ApplicationSid: "AP7e1e264cc0fd7143f3ef378e86bf3184"}
	if nil != validatePost(p) {
		t.Error("Validation of valid voice post failed.")
	}
}

func TestValidatePostFailure(t *testing.T) {
	p := Post{}
	if nil == validatePost(p) {
		t.Error("Validation of voice post missing To & From failed.")
	}

	p = Post{From: testNumber1}
	if nil == validatePost(p) {
		t.Error("Validation of voice post missing From failed.")
	}

	p = Post{From: testNumber1, To: testNumber2}
	if nil == validatePost(p) {
		t.Error("Validation of voice post missing Url & ApplicationSid failed.")
	}
}

func startMockHttpServer(requests *int) *httptest.Server {
	// start a server to recieve post request
	testServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		*requests += 1
		if strings.Contains(r.URL.Path, validEndpoint) {
			fmt.Fprint(resp, testResponseFixture)
		} else if strings.Contains(r.URL.Path, errorEndpoint) {
			resp.WriteHeader(400)
		} else if strings.Contains(r.URL.Path, badJsonEndpoint) {
			fmt.Fprint(resp, testResponseFixture[0:20])
		}
	}))
	return testServer
}

func TestSendSuccess(t *testing.T) {
	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHttpServer(&numRequests)
	defer testPostServer.Close()

	var r Response
	jReader := bytes.NewBuffer([]byte(testPostFixture))
	err := makeCall(testPostServer.URL+validEndpoint, jReader, &r)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	if numRequests != 1 {
		t.Error("Server never recieved a request.")
	}
	if r.AccountSid != testResponseFixtureAccountSid {
		t.Error("Unmarshal failed to properly parse the response.")
	}
}

func TestSendFailure(t *testing.T) {
	// start a server to recieve post request
	numRequests := 0
	testPostServer := startMockHttpServer(&numRequests)
	defer testPostServer.Close()

	var r Response
	jReader := bytes.NewBuffer([]byte(testPostFixture))
	err := makeCall(testPostServer.URL+errorEndpoint, jReader, &r)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 1 {
		t.Error("server never recieved a request.")
	}

	jReader = bytes.NewBuffer([]byte(testPostFixture))
	err = makeCall(testPostServer.URL+badJsonEndpoint, jReader, &r)
	if err == nil {
		t.Errorf("post should've failed with 400")
	}
	if numRequests != 2 {
		t.Error("server never recieved a request.")
	}
}
