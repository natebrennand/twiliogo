package applications

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// holds url values used in queries
var applications = struct {
	Get, Update, List, Create string
}{
	Get:    "/2010-04-01/Accounts/%s/Applications/%s.json", // takes an AccountSid & ApplicationSid
	Update: "/2010-04-01/Accounts/%s/Applications/%s.json", // takes an AccountSid & ApplicationSid
	List:   "/2010-04-01/Accounts/%s/Applications.json",    // takes an AccountSid
	Create: "/2010-04-01/Accounts/%s/Applications.json",    // takes an AccountSid
}

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

var validateApplicationSid = regexp.MustCompile(`^AP[0-9a-z]{32}$`).MatchString

// Resource represents an Application resource
//
// https://www.twilio.com/docs/api/rest/applications
type Resource struct {
	Sid                   string          `json:"sid"`
	DateCreated           common.JSONTime `json:"date_created"`
	DateUpdated           common.JSONTime `json:"date_updated"`
	FriendlyName          string          `json:"friendly_name"`
	AccountSid            string          `json:"account_sid"`
	APIVersion            string          `json:"api_version"`
	VoiceURL              string          `json:"voice_url"`
	VoiceMethod           string          `json:"voice_method"`
	VoiceFallbackURL      string          `json:"voice_fallback_url"`
	VoiceFallbackMethod   string          `json:"voice_fallback_method"`
	VoiceCallerIDLookup   bool            `json:"voice_caller_id_lookup"`
	SmsURL                string          `json:"sms_url"`
	SmsMethod             string          `json:"sms_method"`
	SmsFallbackURL        string          `json:"sms_fallback_url"`
	SmsFallbackMethod     string          `json:"sms_fallback_method"`
	SmsStatusCallback     string          `json:"sms_status_callback"`
	MessageStatusCallback string          `json:"message_status_callback"`
	URI                   string          `json:"uri"`
}

// Get returns all information about an Application.
func (act Account) Get(sid string) (Resource, error) {
	var r Resource
	if !validateApplicationSid(sid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(applications.Get, act.AccountSid, sid), act, &r)
	return r, err
}

// Modification is used to change the state of an application.
type Modification struct {
	FriendlyName          string
	APIVersion            string
	VoiceURL              string
	VoiceMethod           string
	VoiceFallbackURL      string
	VoiceFallbackMethod   string
	VoiceCallerIDLookup   *bool
	SmsURL                string
	SmsMethod             string
	SmsFallbackURL        string
	SmsFallbackMethod     string
	SmsStatusCallback     string
	MessageStatusCallback string
}

// Validate is implemented for the common.twilioPost interface.
func (m Modification) Validate() error {
	if len(m.FriendlyName) > 64 {
		return errors.New("Modification: Invalid FriendlyName, must be <= 64 characters")
	}
	return nil
}

// GetReader is implemented for the common.twilioPost interface.
func (m Modification) GetReader() io.Reader {
	v := url.Values{}
	if m.FriendlyName != "" {
		v.Add("FriendlyName", m.FriendlyName)
	}
	if m.APIVersion != "" {
		//TODO: validate it to the available options
		v.Add("ApiVersion", m.APIVersion)
	}
	if m.VoiceURL != "" {
		v.Add("VoiceUrl", m.VoiceURL)
	}
	if m.VoiceMethod != "" {
		v.Add("VoiceMethod", m.VoiceMethod)
	}
	if m.VoiceFallbackURL != "" {
		v.Add("VoiceFallbackUrl", m.VoiceFallbackURL)
	}
	if m.VoiceFallbackMethod != "" {
		v.Add("VoiceFallbackMethod", m.VoiceFallbackMethod)
	}
	if m.VoiceCallerIDLookup != nil {
		v.Add("VoiceCallerIdLookup", strconv.FormatBool(*m.VoiceCallerIDLookup))
	}
	if m.SmsURL != "" {
		v.Add("SmsUrl", m.SmsURL)
	}
	if m.SmsMethod != "" {
		v.Add("SmsMethod", m.SmsMethod)
	}
	if m.SmsFallbackURL != "" {
		v.Add("SmsFallbackURL", m.SmsFallbackURL)
	}
	if m.SmsFallbackMethod != "" {
		v.Add("SmsFallbackMethod", m.SmsFallbackMethod)
	}
	if m.SmsStatusCallback != "" {
		v.Add("SmsStatusCallback", m.SmsStatusCallback)
	}
	if m.MessageStatusCallback != "" {
		v.Add("MessageStatusCallback", m.MessageStatusCallback)
	}
	return strings.NewReader(v.Encode())
}

// Modify is used with a Modification struc to alter the settings of an Application.
func (act Account) Modify(sid string, m Modification) (Resource, error) {
	var r Resource
	if !validateApplicationSid(sid) {
		return r, errors.New("Invalid application sid")
	}
	if m.Validate() != nil {
		return r, m.Validate()
	}
	err := common.SendPostRequest(fmt.Sprintf(applications.Update, act.AccountSid, sid), m, act, &r)
	return r, err
}

// ResourceList is used to contain the list of applications associated with this account.
type ResourceList struct {
	common.ListResponseCore
	Applications *[]Resource `json:"applications"`
	act          *Account
}

// ListFilter is used to limit the number of applications returned in a List() query.
type ListFilter struct {
	FriendlyName string
}

func (f ListFilter) getQueryString() string {
	v := url.Values{}
	if f.FriendlyName != "" {
		v.Add("FriendlyName", f.FriendlyName)
		return "?" + v.Encode()
	}
	return ""
}

// List returns the list of all applications the fulfill the filter provided.
func (act Account) List(f ListFilter) (ResourceList, error) {
	var rl ResourceList
	err := common.SendGetRequest(fmt.Sprintf(applications.List, act.AccountSid)+f.getQueryString(), act, &rl)
	rl.act = &act
	return rl, err
}

// Next sets the ResourceList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (rl *ResourceList) next() error {
	if rl.Page == rl.NumPages-1 {
		return errors.New("no more new pages")
	}
	return common.SendGetRequest(rl.NextPageURI, *rl.act, rl)
}

// NewResource represents a new application to be ssociated with a Twilio account.
type NewResource struct {
	FriendlyName          string
	AccountSid            string
	APIVersion            string
	VoiceURL              string
	VoiceMethod           string
	VoiceFallbackURL      string
	VoiceFallbackMethod   string
	VoiceCallerIDLookup   bool
	SmsURL                string
	SmsMethod             string
	SmsFallbackURL        string
	SmsFallbackMethod     string
	SmsStatusCallback     string
	MessageStatusCallback string
}

// GetReader is implemented for the common.twilioPost interface.
func (r NewResource) GetReader() io.Reader {
	v := url.Values{}
	if r.FriendlyName != "" {
		v.Add("FriendlyName", r.FriendlyName)
	}
	if r.APIVersion != "" {
		v.Add("ApiVersion", r.APIVersion)
	}
	if r.VoiceURL != "" {
		v.Add("VoiceUrl", r.VoiceURL)
	}
	if r.VoiceMethod != "" {
		v.Add("VoiceMethod", r.VoiceMethod)
	}
	if r.VoiceFallbackURL != "" {
		v.Add("VoiceFallbackUrl", r.VoiceFallbackURL)
	}
	if r.VoiceFallbackMethod != "" {
		v.Add("VoiceFallbackMethod", r.VoiceFallbackMethod)
	}
	if r.SmsURL != "" {
		v.Add("SmsUrl", r.SmsURL)
	}
	if r.SmsMethod != "" {
		v.Add("SmsMethod", r.SmsMethod)
	}
	if r.SmsFallbackURL != "" {
		v.Add("SmsFallbackUrl", r.SmsFallbackURL)
	}
	if r.SmsFallbackMethod != "" {
		v.Add("SmsFallbackMethod", r.SmsFallbackMethod)
	}
	if r.SmsStatusCallback != "" {
		v.Add("SmsStatusCallback", r.SmsStatusCallback)
	}
	if r.MessageStatusCallback != "" {
		v.Add("MessageStatusCallback", r.MessageStatusCallback)
	}
	v.Add("VoiceCallerIdLookup", strconv.FormatBool(r.VoiceCallerIDLookup))
	return strings.NewReader(v.Encode())
}

// Validate is implemented for the common.twilioPost interface.
func (r NewResource) Validate() error {
	if r.FriendlyName == "" {
		return errors.New("NewResource: FriendlyName must be set when creating a new Application")
	}
	// TODO: validate all optional fields, https://www.twilio.com/docs/api/rest/applications#list-post-optional-parameters
	return nil
}

// Create send a request to create a new Twilio Application.
func (act Account) Create(nr NewResource) (Resource, error) {
	var r Resource
	if nr.Validate() != nil {
		return r, nr.Validate()
	}
	err := common.SendPostRequest(fmt.Sprintf(applications.Create, act.AccountSid), nr, act, &r)
	return r, err
}
