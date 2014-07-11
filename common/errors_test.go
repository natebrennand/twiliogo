package common

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNewTwilioErrorSuccess(t *testing.T) {
	resp := http.Response{Body: ioutil.NopCloser(strings.NewReader(testErrorfixtureString))}
	err := NewTwilioError(resp)
	if err == nil {
		t.Error("NewTwilioError failed to create a twilio error")
	}
}

func TestNewTwilioErrorfailure(t *testing.T) {
	// test invalid json
	resp := http.Response{Body: ioutil.NopCloser(strings.NewReader("{"))}
	err := NewTwilioError(resp)
	if err == nil {
		t.Error("NewTwilioError failed to create a twilio error")
	}
}

func TestCommonError(t *testing.T) {
	if testErrorfixtureError.Error() != "Twilio Error 400 => Bad req, more info @ https://www.twilio.com/docs/errors/reference" {
		t.Error("Common.Error failed to properly create an error message")
	}
}

func TestDecodeError(t *testing.T) {
	err, resp := errors.New("TEST1"), []byte(`TEST2`)
	newErr := decodeError(err, resp)
	if newErr.Error() != "Error while decoding json => TEST1, recieved msg => TEST2" {
		t.Error("decodeError failed to format correctly")
	}
}

var testErrorfixtureError = Error{
	Code:     400,
	Message:  "Bad req",
	MoreInfo: "https://www.twilio.com/docs/errors/reference",
	Status:   404,
}
var testErrorfixtureString = `{
	"code": 400,
	"message": "Bad req",
	"more_info": "https://www.twilio.com/docs/errors/reference",
	"status": 404
}`
