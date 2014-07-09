package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/http"
	"net/url"
	"strings"
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
func (act SmsAccount) sendSms(destUrl string, msg Post, resp *Response) error {
	// send post request to twilio
	return common.SendPostRequest(destUrl, msg, act, resp, 201)
}

// Sends a post request to Twilio to send a sms request.
func (act SmsAccount) Send(p Post) (Response, error) {
	var r Response
	err := act.sendSms(fmt.Sprintf(postUrl, act.AccountSid), p, &r)
	return r, err
}
