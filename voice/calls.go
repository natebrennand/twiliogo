package voice

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	validateCallSid      = regexp.MustCompile(`^CA[0-9a-z]{32}$`).MatchString
	validateRecordingSid = regexp.MustCompile(`^RE[0-9a-z]{32}$`).MatchString
)

const (
	postURL   = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json"
	updateURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls/%s.json"
	getURL    = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls/%s.json"
	listURL   = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json"
)

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

// Call represents the data used in creating an outbound voice message.
// "From" & "To" are required attributes.
// Either a ApplicationSid or a URL must also be provided.
// Visit https://www.twilio.com/docs/api/rest/making-calls#post-parameters for more details and
// explanation of other optional parameters.
type Call struct {
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

// GetReader implements the common.twilioPost interface
func (p Call) GetReader() io.Reader {
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

// Validate implements the common.twilioPost interface
func (p Call) Validate() error {
	if p.From == "" || p.To == "" {
		return errors.New("Both \"From\" and \"To\" must be set in Call.")
	}
	if p.ApplicationSid == "" && p.URL == "" {
		return errors.New("Either \"ApplicationSid\" or \"URL\" must be set.")
	}
	if p.SendDigits != "" {
		match, err := regexp.MatchString(`^[0-9#\*w]+$`, p.SendDigits)
		if match != true || err != nil {
			return errors.New("Call's SendDigits can only contain digits, #, * or w")
		}
	}
	return nil
}

// Call creates a new call with Twilio.
func (act Account) Call(p Call) (Resource, error) {
	var r Resource
	err := common.SendPostRequest(fmt.Sprintf(postURL, act.AccountSid), p, act, &r)
	return r, err
}

// Resource represents a call record.
type Resource struct {
	common.ResponseCore
	Price          common.JSONFloat `json:"price"`
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      common.JSONTime `json:"start_time"`
	EndTime        common.JSONTime `json:"end_time"`
	Duration       int64           `json:"duration,string"`
	AnsweredBy     string          `json:"answered_by"`
	ForwardedFrom  string          `json:"fowarded_from"`
	CallerName     string          `json:"caller_name"`
}

// Get returns a call resource record.
func (act Account) Get(sid string) (Resource, error) {
	var m Resource
	if !validateCallSid(sid) {
		return m, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, sid), act, &m)
	return m, err
}

// Update is used to modify a call with Twiml.
type Update struct {
	URL                  string `json:"url"`
	Method               string `json:"method"`
	Status               string `json:"status"`
	FallbackURL          string `json:"fallback_url"`
	FallbackMethod       string `json:"fallback_method"`
	StatusCallback       string `json:"status_callback"`
	StatusCallbackMethod string `json:"status_callback_method"`
}

// GetReader implements the common.twilioPost interface
func (p Update) GetReader() io.Reader {
	vals := url.Values{}
	if p.URL != "" {
		vals.Set("URL", p.URL)
	}
	if p.Status != "" {
		vals.Set("Status", p.Status)
	}
	if p.StatusCallback != "" {
		vals.Set("StatusCallback", p.StatusCallback)
	}
	if p.StatusCallbackMethod != "" {
		vals.Set("StatusCallbackMethod", p.StatusCallbackMethod)
	}
	if p.Method != "" {
		vals.Set("Method", p.Method)
	}
	if p.FallbackURL != "" {
		vals.Set("FallbackURL", p.FallbackURL)
	}
	if p.FallbackMethod != "" {
		vals.Set("FallbackMethod", p.FallbackMethod)
	}

	return strings.NewReader(vals.Encode())
}

// Validate implements the common.twilioPost interface
func (p Update) Validate() error {
	if p.URL == "" && p.Method == "" && p.Status == "" {
		return errors.New("URL or Status or Method must all be set")
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act Account) postUpdate(dest string, msg Update, resp *Resource) error {
	// send post request to twilio
	return common.SendPostRequest(dest, msg, act, resp)
}

// Update sends an update to a Twilio call.
func (act Account) Update(p Update, sid string) (Resource, error) {
	var r Resource
	err := common.SendPostRequest(fmt.Sprintf(updateURL, act.AccountSid, sid), p, act, &r)
	return r, err
}

// CallList represents a list of call records.
type CallList struct {
	common.ListResponseCore
	Calls *[]Resource `json:"calls"`
	act   *Account
}

// ListFilter is used to filter call logs results
type ListFilter struct {
	To            string
	From          string
	Status        string
	StartTime     *time.Time
	ParentCallSid string
	PageSize      int
}

func (f ListFilter) getQueryString() string {
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

// List returns a list of Twilio call records
func (act Account) List(f ListFilter) (CallList, error) {
	var callList CallList
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid)+f.getQueryString(), act, &callList)
	callList.act = &act
	return callList, err
}

// Next sets the MessageList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (cl *CallList) Next() error {
	if cl.Page == cl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(cl.NextPageURI, *cl.act, cl)
}
