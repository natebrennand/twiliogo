package common

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// implements TwilioAccount
type testAccount struct {
	AccountSid string
	Token      string
	Client     http.Client
}

func (act testAccount) GetSid() string {
	return act.AccountSid
}
func (act testAccount) GetToken() string {
	return act.Token
}
func (act testAccount) GetClient() http.Client {
	return act.Client
}

func (act testAccount) sendSms(destURL string, msg testPost, resp *testMessage) error {
	// send post request to twilio
	return SendPostRequest(destURL, msg, act, resp, 201)
}

func (act testAccount) getSms(destURL string, resp *testMessage) error {
	// send post request to twilio
	return SendGetRequest(destURL, act, resp, 200)
}

type testPost struct {
	From  string
	To    string
	Body  string
	valid error
}

func (p testPost) GetReader() io.Reader {
	v := url.Values{}
	v.Set("To", p.To)
	v.Set("From", p.From)
	if p.Body != "" {
		v.Set("Body", p.Body)
	}
	return strings.NewReader(v.Encode())
}

func (p testPost) Validate() error {
	return p.valid
}

type testMessage struct {
	Foo string
}

var (
	testSmsResponseFixtureAccountSid = "AC5ef8732a3c49700934481addd5ce1659"
	testSmsResponseFixtureString     = `
	{
		"account_sid": "AC5ef8732a3c49700934481addd5ce1659",
		"num_segments": "1",
		"num_media": "1",
		"date_created": "Wed, 18 Aug 2010 20:01:40 +0000",
		"date_sent": null,
		"date_updated": "Wed, 18 Aug 2010 20:01:40 +0000"
	}`
	testMessageFixture       = testMessage{Foo: "Bar"}
	testMessageFixtureString = `{"Foo":"Bar"}`
	testPostFixture          = testPost{
		From:  "A",
		To:    "B",
		Body:  "Yo",
		valid: nil,
	}
	testPostFixtureInvalid = testPost{
		From:  "A",
		To:    "B",
		Body:  "Yo",
		valid: errors.New(""),
	}
	testStandardRequestFormString = `AccountSid=AC381707b751dbe4c74b15c5697ba67afd&From=+14248004123&To=+13605847116`
	testStandardRequest           = StandardRequest{
		AccountSid: "AC381707b751dbe4c74b15c5697ba67afd",
		From:       "+14248004123",
		To:         "+13605847116",
	}
	testStandardRequestFormStringWithCity = `AccountSid=AC381707b751dbe4c74b15c5697ba67afd&From=+14248004123&To=+13605847116&FromCity=SF`
	testStandardRequestWithCity           = StandardRequest{
		AccountSid: "AC381707b751dbe4c74b15c5697ba67afd",
		From:       "+14248004123",
		To:         "+13605847116",
		Location: &Location{
			FromCity: "SF",
		},
	}
)

type testMessageWithDate struct {
	ResponseCore
	NumSegments int       `json:"num_segments,string"`
	NumMedia    int       `json:"num_media,string"`
	Price       JSONPrice `json:"price"`
	DateCreated JSONTime  `json:"date_created"`
	DateSent    JSONTime  `json:"date_sent"`
	DateUpdated JSONTime  `json:"date_updated"`
}
