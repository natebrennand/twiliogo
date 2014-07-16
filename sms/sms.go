package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/act"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Account wraps the act Account struct to embed the AccountSid & Token.
type Account struct {
	act.Account
}

// holds url values used in queries
var sms = struct {
	Post, Get, List string
}{
	Post: "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",    // takes an AccountSid
	Get:  "https://api.twilio.com/2010-04-01/Accounts/%s/Messages/%s.json", // takes an AccountSid & MessageSdi
	List: "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",    // takes an AccountSid
}

var validateSmsSid = regexp.MustCompile(`^(SM|MM)[0-9a-z]{32}$`).MatchString

// Message represents an instance of a SMS resource
type Message struct {
	common.ResponseCore
	Body        string           `json:"body"`
	DateSent    common.JSONTime  `json:"date_sent"`
	NumSegments int              `json:"num_segments,string"`
	NumMedia    int              `json:"num_media,string"`
	Price       common.JSONFloat `json:"price"`
}

// MessageList represents an list of a SMS resources
type MessageList struct {
	common.ListResponseCore
	Messages *[]Message `json:"messages"`
}

// Get a message given that message's sid.
func (act Account) Get(sid string) (Message, error) {
	var m Message
	if !validateSmsSid(sid) {
		return m, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(sms.Get, act.AccountSid, sid), act, &m)
	return m, err
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
	err := common.SendPostRequest(fmt.Sprintf(sms.Post, act.AccountSid), p, act, &m)
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
	err := common.SendGetRequest(fmt.Sprintf(sms.List, act.AccountSid)+f.getQueryString(), act, &ml)
	return ml, err
}
