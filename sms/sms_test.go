package sms

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

var (
	validSmsSid = "SM90c6fc909d8504d45ecdb3a3d5b3556e"
)

func TestValidateSmsPostSuccess(t *testing.T) {
	p := Post{From: testNumber1, To: testNumber2, Body: "test"}
	if nil != p.Validate() {
		t.Error("Validation of valid sms post failed.")
	}

	p = Post{From: testNumber1, To: testNumber2, MediaURL: "https://www.twilio.com/"}
	if nil != p.Validate() {
		t.Error("Validation of valid sms post failed.")
	}
}

func TestValidateSmsPostFailure(t *testing.T) {
	p := Post{}
	if nil == p.Validate() {
		t.Error("Validation of sms post missing To & From failed.")
	}

	p = Post{From: testNumber1}
	if nil == p.Validate() {
		t.Error("Validation of sms post missing From failed.")
	}

	p = Post{From: testNumber1, To: testNumber2}
	if nil == p.Validate() {
		t.Error("Validation of sms post missing Body & MediaURL failed.")
	}
}

func TestListFilterReader(t *testing.T) {
	// empty filter
	f := Filter{}
	if "" != f.getQueryString() {
		t.Error("url encoding of filter should be empty")
	}

	// w/ To
	f = Filter{To: "12345"}
	if f.getQueryString() != "?To=12345" {
		t.Errorf("url encoding of filter should include To&From key:value pairs, found => %s", f.getQueryString())
	}

	// w/ To & From
	f = Filter{To: "12345", From: "678"}
	if !strings.Contains(f.getQueryString(), "To=12345") || !strings.Contains(f.getQueryString(), "From=678") {
		t.Errorf("url encoding of filter should include To&From key:value pairs, found => %s", f.getQueryString())
	}

	// w/ Date
	tm := time.Date(2010, time.August, 16, 3, 45, 01, 0, &time.Location{})
	f = Filter{DateSent: &tm}
	if "?DateSent=2010-08-16" != f.getQueryString() {
		t.Errorf("url encoding of filter should encode dates in GMT, found => %s", f.getQueryString())
	}
}

func TestGetSms(t *testing.T) {
	_, err := act.Get("sldkfj")
	assert.Error(t, err)

	// TODO: test HTTP call
}

func TestPostValidate(t *testing.T) {
	b, err := ioutil.ReadAll(testSmsPostFixture.GetReader())
	assert.Nil(t, err)
	form := string(b)

	assert.Contains(t, form, "Body=Jenny+please%3F%21+I+love+you+%3C3")
	assert.Contains(t, form, "To=%2B15558675309")
	assert.Contains(t, form, "From=%2B14158141829")
	assert.Contains(t, form, "MediaURL=http%3A%2F%2Fwww.example.com%2Fhearts.png")
	assert.Contains(t, form, "StatusCallback=foobar.com")
	assert.Contains(t, form, "ApplicationSid=AP5ef8732a3c49700934481addd5ce1659")
}

func TestSendSms(t *testing.T) {
	// TODO: test HTTP call
}

func TestListSms(t *testing.T) {
	// TODO: test HTTP call
}

func TestGetSmsList(t *testing.T) {
	var ml MessageList
	ml.Page = 0
	ml.NumPages = 1
	assert.Error(t, ml.Next())

	// TODO: test HTTP call
}
