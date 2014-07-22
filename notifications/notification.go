package notifications

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// holds url values used in queries
var notifications = struct {
	Get, Delete, List string
}{
	Get:    "/2010-04-01/Accounts/%s/Notifications/%s.json", // takes an AccountSid & NotifcationSid
	Delete: "/2010-04-01/Accounts/%s/Notifications/%s.json", // takes an AccountSid & NotifcationSid
	List:   "/2010-04-01/Accounts/%s/Notifications.json",    // takes an AccountSid
}

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

var validateNotificationSid = regexp.MustCompile(`^NO[0-9a-z]{32}$`).MatchString

// Notification represents an error logged by Twilio.
//
// https://www.twilio.com/docs/api/rest/notification#instance-properties
type Notification struct {
	common.ResponseCore2
	CallSid          string `json:"call_sid"`
	APIVersion       string `json:"api_version"`
	Log              int64  `json:"log,string"`
	ErrorCode        string `json:"error_code"`
	MoreInfo         string `json:"more_info"`
	MessageText      string `json:"message_text"`
	RequestURL       string `json:"request_url"`
	RequestMethod    string `json:"request_method"`
	RequestVariables string `json:"request_variables"`
	ResponseHeaders  string `json:"response_headers"`
	ResponseBody     string `json:"response_body"`
}

// Get retrieves a Twilio error log based off of the sid of the notification.
func (act Account) Get(sid string) (Notification, error) {
	var r Notification
	if !validateNotificationSid(sid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(notifications.Get, act.AccountSid, sid), act, &r)
	return r, err
}

// Delete removes a Twilio error notification based off of the sid.
func (act Account) Delete(sid string) error {
	if !validateNotificationSid(sid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(notifications.Get, act.AccountSid, sid), act)
}

// ReducedNotification is similar to a full Notification but lacks
// RequestVariables, ResponseHeaders, &ResponseBody. It is used in list responses.
type ReducedNotification struct {
	common.ResponseCore2
	CallSid       string `json:"call_sid"`
	APIVersion    string `json:"api_version"`
	Log           int    `json:"log,string"`
	ErrorCode     string `json:"error_code"`
	MoreInfo      string `json:"more_info"`
	MessageText   string `json:"message_text"`
	RequestURL    string `json:"request_url"`
	RequestMethod string `json:"request_method"`
}

// NotificationList represents a list of notifications returne by a query.
type NotificationList struct {
	common.ListResponseCore
	Notifications *[]ReducedNotification `json:"notifications"`
}

// Filter is used to filter notification lists. You may set OnMessageDate or a combination of
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

func (f Filter) getQueryString() string {
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

// List returns a list of recent Twilio error notifications that pass a filter you supply.
func (act Account) List(f Filter) (NotificationList, error) {
	var rl NotificationList
	err := f.validate()
	if err != nil {
		return rl, err
	}
	err = common.SendGetRequest(fmt.Sprintf(notifications.List, act.AccountSid)+f.getQueryString(), act, &rl)
	return rl, err
}
