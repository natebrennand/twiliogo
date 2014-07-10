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

func (s SmsAccount) GetSid() string {
	return s.AccountSid
}
func (s SmsAccount) GetToken() string {
	return s.Token
}
func (s SmsAccount) GetClient() http.Client {
	return s.Client
}

// Represents the data used in creating an outbound sms message.
// "From" & "To" are required attributes.
// Either a Body or a MediaUrl must also be provided.
// "StatusCallback" and "ApplicationSid" are both optional.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#post for more details.
type Post struct {
	From           string
	To             string
	Body           string
	MediaUrl       string
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
	if p.MediaUrl != "" {
		v.Set("MediaUrl", p.MediaUrl)
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
	if p.Body == "" && p.MediaUrl == "" {
		return errors.New(`Either "Body" or "MediaUrl" must be set.`)
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) sendSms(destUrl string, msg Post, resp *Message) error {
	// send post request to twilio
	return common.SendPostRequest(destUrl, msg, act, resp, 201)
}

// Sends a post request to Twilio to send a sms request.
func (act SmsAccount) Send(p Post) (Message, error) {
	var m Message
	err := act.sendSms(fmt.Sprintf(postUrl, act.AccountSid), p, &m)
	return m, err
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) getSms(destUrl string, resp *Message) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp, 200)
}

func (act SmsAccount) Get(sid string) (Message, error) {
	if true != validateSmsSid(sid) {
		return Message{}, errors.New("Invalid sid")
	}
	var m Message
	err := act.getSms(fmt.Sprintf(getUrl, act.AccountSid, string(sid)), &m)
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

func (act SmsAccount) getList(destUrl string, f Filter, resp *MessageList) error {
	return common.SendGetRequest(destUrl+f.GetQueryString(), act, resp, 200)
}

func (act SmsAccount) List(f Filter) (MessageList, error) {
	var ml MessageList
	err := act.getList(fmt.Sprintf(listUrl, act.AccountSid), f, &ml)
	return ml, err
}
