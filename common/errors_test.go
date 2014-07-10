package common

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNewTwilioError(t *testing.T) {
	resp := http.Response{Body: ioutil.NopCloser(strings.NewReader(testErrorFixtureString))}
	err := NewTwilioError(resp)
	if err == nil {
		t.Error("NewTwilioError failed to create a twilio error")
	}
}

var testErrorFixture Error = Error{
	Code:     400,
	Message:  "Bad req",
	MoreInfo: "https://www.twilio.com/docs/errors/reference",
	Status:   404,
}
var testErrorFixtureString = `{
	"code": 400,
	"message": "Bad req",
	"more_info": "https://www.twilio.com/docs/errors/reference",
	"status": 404
}`
