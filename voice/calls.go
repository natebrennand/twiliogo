package voice

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Account struct {
	AccountSid string
	Token      string
	Client     http.Client
}

func (act Account) GetSid() string {
	return act.AccountSid
}
func (act Account) GetToken() string {
	return act.Token
}
func (act Account) GetClient() http.Client {
	return act.Client
}

// Represents the data used in creating an outbound voice message.
// "From" & "To" are required attributes.
// Either a ApplicationSid or a URL must also be provided.
// Visit https://www.twilio.com/docs/api/rest/making-calls#post-parameters for more details and
// explanation of other optional parameters.
type Post struct {
	From                 string
	To                   string
	Body                 string
	URL                  string
	ApplicationSid       string
	StatusCallback       string
	Method               string
	FallbackURL          string
	StatusCallbackMethod string
	SendDigits           string
	IfMachine            string
	TimeOut              *int64
	Record               *bool
}

func (p Post) GetReader() io.Reader {
	vals := url.Values{}
	vals.Set("To", p.To)
	vals.Set("From", p.From)
	if p.URL != "" {
		vals.Set("Url", p.URL)
	}
	if p.ApplicationSid != "" {
		vals.Set("ApplicationSid", p.ApplicationSid)
	}
	if p.StatusCallback != "" {
		vals.Set("StatusCallback", p.StatusCallback)
	}
	if p.Method != "" {
		vals.Set("Method", p.Method)
	}
	if p.FallbackURL != "" {
		vals.Set("FallbackURL", p.FallbackURL)
	}
	if p.StatusCallbackMethod != "" {
		vals.Set("StatusCallbackMethod", p.StatusCallbackMethod)
	}
	if p.SendDigits != "" {
		vals.Set("SendDigits", p.SendDigits)
	}
	if p.IfMachine != "" {
		vals.Set("IfMachine", p.IfMachine)
	}
	if p.TimeOut != nil {
		vals.Set("TimeOut", strconv.FormatInt(*p.TimeOut, 10))
	}
	if p.Record != nil {
		vals.Set("Record", strconv.FormatBool(*p.Record))
	}

	return strings.NewReader(vals.Encode())
}

// Validates a Voice Post.
func (p Post) Validate() error {
	if p.From == "" || p.To == "" {
		return errors.New("Both \"From\" and \"To\" must be set in Post.")
	}
	if p.ApplicationSid == "" && p.URL == "" {
		return errors.New("Either \"ApplicationSid\" or \"URL\" must be set.")
	}
	if p.SendDigits != "" {
		match, err := regexp.MatchString(`^[0-9#\*w]+$`, p.SendDigits)
		if match != true || err != nil {
			return errors.New("Post's SendDigits can only contain digits, #, * or w")
		}
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act Account) makeCall(dest string, msg Post, resp *Call) error {
	// send post request to twilio
	return common.SendPostRequest(dest, msg, act, resp)
}

// Sends a post request to Twilio to create a call.
func (act Account) Call(p Post) (Call, error) {
	var r Call
	err := act.makeCall(fmt.Sprintf(postURL, act.AccountSid), p, &r)
	return r, err
}

// Internal function for sending the post request to twilio.
func (act Account) getCall(destURL string, resp *Call) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp)
}

func (act Account) Get(sid string) (Call, error) {
	var m Call
	if !validateCallSid(sid) {
		return m, errors.New("Invalid sid")
	}

	err := act.getCall(fmt.Sprintf(getURL, act.AccountSid, string(sid)), &m)
	return m, err
}

// Used to filter call logs results
type ListFilter struct {
	To            string
	From          string
	Status        string
	StartTime     *time.Time
	ParentCallSid string
	PageSize      int
}

func (f ListFilter) GetQueryString() string {
	v := url.Values{}
	if f.To != "" {
		v.Set("To", f.To)
	}
	if f.From != "" {
		v.Set("From", f.From)
	}
	if f.Status != "" {
		v.Set("Status", f.Status)
	}
	if f.StartTime != nil {
		v.Set("StartTime", f.StartTime.Format(common.GMTTimeLayout))
	}
	if f.ParentCallSid != "" {
		v.Set("ParentCallSid", f.ParentCallSid)
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

func (act Account) getList(destURL string, f ListFilter, resp *CallList) error {
	return common.SendGetRequest(destURL+f.GetQueryString(), act, resp)
}

func (act Account) List(f ListFilter) (CallList, error) {
	var callList CallList
	err := act.getList(fmt.Sprintf(listURL, act.AccountSid), f, &callList)
	return callList, err
}
