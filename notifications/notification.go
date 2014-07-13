package notifications

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

const (
	getURL    = "https://api.twilio.com/2010-04-01/Accounts/%s/Notifications/%s.json" // takes an AccountSid & NotifcationSid
	deleteURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Notifications/%s.json" // takes an AccountSid & NotifcationSid
	listURL   = "https://api.twilio.com/2010-04-01/Accounts/%s/Notifications.json"    // takes an AccountSid
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

var validateNotificationSid = regexp.MustCompile(`^NO[0-9a-z]{32}$`).MatchString

// https://www.twilio.com/docs/api/rest/notification#instance-properties
type Resource struct {
	Sid              string          `json:"sid"`
	AccountSid       string          `json:"account_sid"`
	DateCreated      common.JSONTime `json:"date_created"`
	DateUpdated      common.JSONTime `json:"date_updated"`
	CallSid          string          `json:"call_sid"`
	APIVersion       string          `json:"api_version"`
	Log              int64           `json:"log,string"`
	ErrorCode        string          `json:"error_code"`
	MoreInfo         string          `json:"more_info"`
	MessageText      string          `json:"message_text"`
	RequestURL       string          `json:"request_url"`
	RequestMethod    string          `json:"request_method"`
	RequestVariables string          `json:"request_variables"`
	ResponseHeaders  string          `json:"response_headers"`
	ResponseBody     string          `json:"response_body"`
	URI              string          `json:"uri"`
}

func (act Account) Get(sid string) (Resource, error) {
	var r Resource
	if !validateNotificationSid(sid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, sid), act, &r)
	return r, err
}

func (act Account) Delete(sid string) error {
	if !validateNotificationSid(sid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(getURL, act.AccountSid, sid), act)
}

// Is similar to a full resource but lacks
//	RequestVariables, ResponseHeaders, &ResponseBody
type ReducedResource struct {
	Sid           string          `json:"sid"`
	AccountSid    string          `json:"account_sid"`
	DateCreated   common.JSONTime `json:"date_created"`
	DateUpdated   common.JSONTime `json:"date_updated"`
	CallSid       string          `json:"call_sid"`
	APIVersion    string          `json:"api_version"`
	Log           int             `json:"log,string"`
	ErrorCode     string          `json:"error_code"`
	MoreInfo      string          `json:"more_info"`
	MessageText   string          `json:"message_text"`
	RequestURL    string          `json:"request_url"`
	RequestMethod string          `json:"request_method"`
	URI           string          `json:"uri"`
}

type ResourceList struct {
	common.ListResponseCore
	Notifications *[]ReducedResource `json:"notifications"`
}

// Used to filter notification lists. You may set OnMessageDate or a combination of
// BeforeMessageDate and AfterMessageDate.
//
// https://www.twilio.com/docs/api/rest/notification#list-get-filters
type Filter struct {
	Log               *int64
	OnMessageDate     *time.Time
	BeforeMessageDate *time.Time
	AfterMessageDate  *time.Time
}

func (f Filter) validate() error {
	if f.OnMessageDate != nil && (f.BeforeMessageDate != nil || f.AfterMessageDate != nil) {
		return errors.New(`Only "On" or a combination of "Before" and "After" can be set`)
	}
	return nil
}

func (f Filter) GetQueryString() string {
	v := url.Values{}
	if f.Log != nil {
		v.Set("To", strconv.FormatInt(*f.Log, 10))
	}
	// Only allow "On" or a combination of "After" & "Before"
	if f.OnMessageDate != nil {
		v.Set("MessageDate", f.OnMessageDate.Format(common.GMTTimeLayout))
	} else {
		if f.AfterMessageDate != nil {
			v.Set("MessageDate>", f.AfterMessageDate.Format(common.GMTTimeLayout))
		}
		if f.BeforeMessageDate != nil {
			v.Set("MessageDate<", f.BeforeMessageDate.Format(common.GMTTimeLayout))
		}
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

func (act Account) List(f Filter) (ResourceList, error) {
	var rl ResourceList
	err := f.validate()
	if err != nil {
		return rl, err
	}
	err = common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid)+f.GetQueryString(), act, &rl)
	return rl, err
}
