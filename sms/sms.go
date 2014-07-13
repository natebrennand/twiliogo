package sms

import (
	"github.com/natebrennand/twiliogo/act"
	"github.com/natebrennand/twiliogo/common"

	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

// Account wraps the act Account struct to embed the AccountSid & Token.
type Account struct {
	act.Account
}

// Post represents the data used in creating an outbound sms message.
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

// GetReader encodes the Post into an io.Reader for consumption while building a HTTP request to
// Twilio.
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

// Validate the Voice Post to ensure validity.
func (p Post) Validate() error {
	if p.From == "" || p.To == "" {
		return errors.New(`Both "From" and "To" must be set in Post.`)
	}
	if p.Body == "" && p.MediaURL == "" {
		return errors.New(`Either "Body" or "MediaURL" must be set.`)
	}
	return nil
}

// Send a post request to Twilio to send a sms request.
func (act Account) Send(p Post) (Message, error) {
	var m Message
	if nil != p.Validate() {
		return m, p.Validate()
	}
	err := common.SendPostRequest(fmt.Sprintf(postURL, act.AccountSid), p, act, &m)
	return m, err
}

// Get a message given that message's sid.
func (act Account) Get(sid string) (Message, error) {
	var m Message
	if !validateSmsSid(sid) {
		return m, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, sid), act, &m)
	return m, err
}

// Filter is used to filter list SMS results
type Filter struct {
	To       string
	From     string
	DateSent *time.Time
}

func (f Filter) getQueryString() string {
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

// List returns a list of message records given a filter.
func (act Account) List(f Filter) (MessageList, error) {
	var ml MessageList
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid)+f.getQueryString(), act, &ml)
	return ml, err
}
