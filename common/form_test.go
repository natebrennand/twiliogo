package common

import (
	"net/http"
	"strings"
	"testing"
)

// helper fn that creates a valid callback request
func makeTestCallbackReq(fixture string) *http.Request {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(fixture))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

// Tests that a correct Callback request is properly parsed into a callback struct.
func TestParseStandardRequest(t *testing.T) {
	req := makeTestCallbackReq(testStandardRequestFormString)
	stdReq := ParseStandardRequest(req)

	if testStandardRequest.AccountSid != stdReq.AccountSid || stdReq.Location != nil {
		t.Error("ParseStandardRequest failed to set AccountSid")
	}
}

func TestParseStandardRequestWithLocation(t *testing.T) {
	req := makeTestCallbackReq(testStandardRequestFormStringWithCity)
	stdReq := ParseStandardRequest(req)

	if testStandardRequestWithCity.AccountSid != stdReq.AccountSid || stdReq.Location == nil {
		t.Error("ParseStandardRequest failed to set AccountSid")
	}
	if stdReq.Location.FromCity != testStandardRequestWithCity.Location.FromCity {
		t.Error("ParseStandardRequest failed to set Location's FromCity")
	}
}
