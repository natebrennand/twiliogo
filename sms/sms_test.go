package sms

import (
	"strings"
	"testing"
	"time"
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
