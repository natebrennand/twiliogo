package voice

import (
	// "github.com/natebrennand/twiliogo/common"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func makeTestCallbackReq() *http.Request {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(testCallbackFixtureFormString))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

func makeTestCallbackFailure() *http.Request {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(testCallbackFixtureFormFailureString))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

func TestParseCallbackSuccess(t *testing.T) {
	var cb Callback
	req := makeTestCallbackReq()
	err := cb.Parse(req)
	if err != nil {
		t.Errorf("parseCallback failed with => %s", err.Error())
	}

	if testCallbackFixture.AccountSid != cb.AccountSid ||
		cb.Location != nil {
		t.Error("parseCallback failed to read in AccountSid")
	}
}

func TestCallbackHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	req := makeTestCallbackReq()
	cbChan := make(chan Callback)
	handler := CallbackHandler(cbChan)
	go func() {
		cb := <-cbChan
		if cb.AccountSid != testCallbackFixture.AccountSid {
			t.Error("CallbackHandler failed to parse properly")
		}
	}()

	handler(resp, req)
	if resp.Code != http.StatusOK {
		t.Error("CallbackHandler failed to write OK status")
	}
}
