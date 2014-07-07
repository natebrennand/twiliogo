package sms

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testSendUrl = "/sms/send"
)

func TestSendSms(t *testing.T) {
	testPostServer := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		fmt.Fprint(resp, testSmsResponseFixture)
	}))
	defer testPostServer.Close()

	var r Response
	jReader := bytes.NewBuffer([]byte(testSmsPostFixture))
	err := sendSms(testPostServer.URL, jReader, &r)
	if err != nil {
		t.Errorf("Error while sending post request => %s", err.Error())
	}
	fmt.Printf("%#v\n", r)
}
