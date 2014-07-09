package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

func (t testAccount) GetSid() string {
	return t.AccountSid
}
func (t testAccount) GetToken() string {
	return t.Token
}
func (t testAccount) GetClient() http.Client {
	return t.Client
}

func (act testAccount) sendSms(destUrl string, msg testPost, resp *testMessage) error {
	// send post request to twilio
	return SendPostRequest(destUrl, msg, act, resp, 201)
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

func (r *testMessage) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
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
)

type testMessageWithDate struct {
	ResponseCore
	NumSegments int       `json:"num_segments,string"`
	NumMedia    int       `json:"num_media,string"`
	Price       JsonPrice `json:"price"`
	DateCreated JsonTime  `json:"date_created"`
	DateSent    JsonTime  `json:"date_sent"`
	DateUpdated JsonTime  `json:"date_updated"`
}
