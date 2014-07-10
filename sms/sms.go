package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SmsAccount struct {
	AccountSid string
	Token      string
	Client     http.Client
}

func (act SmsAccount) GetSid() string {
	return act.AccountSid
}
func (act SmsAccount) GetToken() string {
	return act.Token
}
func (act SmsAccount) GetClient() http.Client {
	return act.Client
}

// Represents the data used in creating an outbound sms message.
// "From" & "To" are required attributes.
// Either a Body or a MediaURL must also be provided.
// "StatusCallback" and "ApplicationSid" are both optional.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#post for more details.
type Post struct {
	From           string
	To             string
	Body           string
	MediaURL       string
	StatusCallback string
	ApplicationSid string
}

func (p Post) GetReader() io.Reader {
	v := url.Values{}
	v.Set("To", p.To)
	v.Set("From", p.From)
	if p.Body != "" {
		v.Set("Body", p.Body)
	}
	if p.MediaURL != "" {
		v.Set("MediaURL", p.MediaURL)
	}
	if p.StatusCallback != "" {
		v.Set("StatusCallback", p.StatusCallback)
	}
	if p.ApplicationSid != "" {
		v.Set("ApplicationSid", p.ApplicationSid)
	}
	return strings.NewReader(v.Encode())
}

// Validates the Voice Post to ensure validity.
func (p Post) Validate() error {
	if p.From == "" || p.To == "" {
		return errors.New(`Both "From" and "To" must be set in Post.`)
	}
	if p.Body == "" && p.MediaURL == "" {
		return errors.New(`Either "Body" or "MediaURL" must be set.`)
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) sendSms(destURL string, msg Post, resp *Message) error {
	// send post request to twilio
	return common.SendPostRequest(destURL, msg, act, resp, 201)
}

// Sends a post request to Twilio to send a sms request.
func (act SmsAccount) Send(p Post) (Message, error) {
	var m Message
	err := act.sendSms(fmt.Sprintf(postURL, act.AccountSid), p, &m)
	return m, err
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) getSms(destURL string, resp *Message) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp, 200)
}

func (act SmsAccount) Get(sid string) (Message, error) {
	var m Message
	if !validateSmsSid(sid) {
		return m, errors.New("Invalid sid")
	}
	err := act.getSms(fmt.Sprintf(getURL, act.AccountSid, string(sid)), &m)
	return m, err
}

// Used to filter list SMS results
type Filter struct {
	To       string
	From     string
	DateSent *time.Time
}

func (f Filter) GetQueryString() string {
	v := url.Values{}
	if f.To != "" {
		v.Set("To", f.To)
	}
	if f.From != "" {
		v.Set("From", f.From)
	}
	if f.DateSent != nil {
		v.Set("DateSent", f.DateSent.Format(common.GMTTimeLayout))
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

func (f Filter) Validate() error {
	return nil
}

func (act SmsAccount) getList(destURL string, f Filter, resp *MessageList) error {
	return common.SendGetRequest(destURL+f.GetQueryString(), act, resp, 200)
}

func (act SmsAccount) List(f Filter) (MessageList, error) {
	var ml MessageList
	err := act.getList(fmt.Sprintf(listURL, act.AccountSid), f, &ml)
	return ml, err
}
