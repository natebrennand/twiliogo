package sms

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// helper fn that creates a valid callback request
func makeTestCallbackReq() *http.Request {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(testSmsCallbackFixtureFormString))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

// helper fn that creates an invalid callback request
func makeTestCallbackReqFailure() *http.Request {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(testSmsCallbackFixtureFormStringFailure))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

// Tests that a correct Callback request is properly parsed into a callback struct.
func TestParseCallbackSuccess(t *testing.T) {
	var cb Callback
	req := makeTestCallbackReq()
	err := cb.Parse(req)
	if err != nil {
		t.Errorf("parseCallback failed with => %s", err.Error())
	}

	if testSmsCallbackFixture.AccountSid != cb.AccountSid ||
		cb.Location != nil ||
		len(cb.MediaList) != 0 {
		t.Error("parseCallback failed to read in AccountSid")
	}
}

// Tests that an incorrect Callback request is detected and properly returns an error.
func TestParseCallbackFailure(t *testing.T) {
	var cb Callback
	req := makeTestCallbackReqFailure()

	err := cb.Parse(req)
	if err == nil {
		t.Errorf("parseCallback should've failed with a NumMedia error")
	}
}

// Test that the produced http Handler sends the parsed Callback into the provided channel.
func TestCallbackHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	req := makeTestCallbackReq()
	cbChan := make(chan Callback)
	handler := CallbackHandler(cbChan)
	go func() {
		cb := <-cbChan
		if cb.AccountSid != testSmsCallbackFixture.AccountSid {
			t.Error("CallbackHandler failed to parse the Callback properly")
		}
	}()

	handler(resp, req)
	if resp.Code != http.StatusOK {
		t.Error("CallbackHandler failed to write successful status")
	}
}

// Test that the produced http Handler does not send the incorrect Callback request into the provided channel.
func TestCallbackHandlerOnFailure(t *testing.T) {
	resp := httptest.NewRecorder()
	req := makeTestCallbackReqFailure()
	cbChan := make(chan Callback)
	handler := CallbackHandler(cbChan)

	handler(resp, req)
	if resp.Code != 400 {
		t.Errorf("CallbackHandler failed to write failure status, wrote %d", resp.Code)
	}
}
