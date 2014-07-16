package voice

import (
	"testing"
)

const (
	validEndpoint   = "/valid"
	errorEndpoint   = "/error"
	badJSONEndpoint = "/badJSON"
)

func TestValidatePostSuccess(t *testing.T) {
	p := Post{From: testNumber1, To: testNumber2, URL: "http://twimlbin.com/558a498f"}
	if nil != p.Validate() {
		t.Error("Validation of valid voice post failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, ApplicationSid: "AP7e1e264cc0fd7143f3ef378e86bf3184"}
	if nil != p.Validate() {
		t.Error("Validation of valid voice post failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, URL: "http://twimlbin.com/558a498f", SendDigits: "1234"}
	if nil != p.Validate() {
		t.Error("Validation of valid voice post failed with SendDigits.")
	}
}

func TestValidatePostFailure(t *testing.T) {
	p := Post{}
	if nil == p.Validate() {
		t.Error("Validation of voice post missing To & From failed.")
	}

	p = Post{From: testNumber1}
	if nil == p.Validate() {
		t.Error("Validation of voice post missing From failed.")
	}

	p = Post{From: testNumber1, To: testNumber2}
	if nil == p.Validate() {
		t.Error("Validation of voice post missing URL & ApplicationSid failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, URL: "http://twimlbin.com/558a498f", SendDigits: "1234a"}
	if nil == p.Validate() {
		t.Error("Validation of invalid voice post failed with bad SendDigits.")
	}
}
