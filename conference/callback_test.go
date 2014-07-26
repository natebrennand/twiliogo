package conference

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// helper fn that creates a valid callback request
func makeTestCallbackReq() *http.Request {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(`RecordingUrl=foobar.com/owl`))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req
}

func TestCallbackParse(t *testing.T) {
	expectedCB := Callback{"foobar.com/owl"}
	var cb Callback

	req := makeTestCallbackReq()

	assert.Nil(t, cb.Parse(req))
	assert.Equal(t, expectedCB, cb)
}

// Test that the produced http Handler sends the parsed Callback into the provided channel.
func TestCallbackHandler(t *testing.T) {
	expectedCB := Callback{"foobar.com/owl"}

	resp := httptest.NewRecorder()
	req := makeTestCallbackReq()
	cbChan := make(chan Callback)
	handler := CallbackHandler(cbChan)

	// send HTTP request to handler
	handler(resp, req)
	// HTTP request ends and is THEN sends the callback into the channel
	cb := <-cbChan
	if cb.RecordingURL != expectedCB.RecordingURL {
		t.Error("CallbackHandler failed to parse the Callback properly")
	}
	if resp.Code != http.StatusOK {
		t.Error("CallbackHandler failed to write successful status")
	}
}
