package sms

import (
	"net/http"
	"strings"
	"testing"
)

func TestParseCallbackSuccess(t *testing.T) {
	var cb Callback
	req, _ := http.NewRequest("POST", "/", strings.NewReader(testSmsCallbackFixtureFormString))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	err := parseCallback(req, &cb)
	if err != nil {
		t.Errorf("parseCallback failed with => %s", err.Error())
	}

	if testSmsCallbackFixture.AccountSid != cb.AccountSid {
		t.Error("parseCallback failed to read in AccountSid")
	}
}

func TestParseCallbackFailure(t *testing.T) {
	var cb Callback
	req, _ := http.NewRequest("POST", "/", strings.NewReader(testSmsCallbackFixtureFormStringFailure))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	err := parseCallback(req, &cb)
	if err == nil {
		t.Errorf("parseCallback should've failed with a NumMedia error")
	}
}
