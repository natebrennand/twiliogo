package shortcodes

import (
	"github.com/natebrennand/twiliogo/act"
	"github.com/natebrennand/twiliogo/common"

	"errors"
	"fmt"
	"net/url"
	"regexp"
)

var short = struct {
	Get, List string
}{
	Get:  "https://api.twilio.com/2010-04-01/Accounts/%s/SMS/ShortCodes/%s.json", // takes an AccountSid & ShortcodeSid
	List: "https://api.twilio.com/2010-04-01/Accounts/%s/SMS/ShortCodes.json",    // takes an AccountSid
}

// Account wraps the act Account struct to embed the AccountSid & Token.
type Account struct {
	act.Account
}

var validateShortcodeSid = regexp.MustCompile(`^SC[0-9a-z]{32}$`).MatchString

// Resource represents a short code message resource.
//
// https://www.twilio.com/docs/api/rest/short-codes
type Resource struct {
	Sid            string          `json:"sid"`
	DateCreated    common.JSONTime `json:"date_created"`
	DateUpdated    common.JSONTime `json:"date_updated"`
	FriendlyName   string          `json:"friendly_name"`
	AccountSid     string          `json:"account_sid"`
	ShortCode      string          `json:"short_code"`
	APIVersion     string          `json:"api_version"`
	SmsURL         string          `json:"sms_url"`
	SmsMethod      string          `json:"sms_method"`
	SmsFallbackURL string          `json:"sms_fallback_url"`
	URI            string          `json:"uri"`
}

// Get returns a shortcode message resource given a sid.
func (act Account) Get(sid string) (Resource, error) {
	var r Resource
	if !validateShortcodeSid(sid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(short.Get, act.AccountSid, sid), act, &r)
	return r, err
}

// ListFilter provides a way to filter the results returned by List()
type ListFilter struct {
	ShortCode    string
	FriendlyName string
}

// renders the query string of the filter
func (f ListFilter) getQueryString() string {
	v := url.Values{}
	if f.ShortCode != "" {
		v.Set("ShortCode", f.ShortCode)
	}
	if f.FriendlyName != "" {
		v.Set("FriendlyName", f.FriendlyName)
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

// ResourceList contains a list of all shortcode messages that matched the given filtered query.
type ResourceList struct {
	common.ListResponseCore
	ShortCodes *[]Resource
}

// List returns a list of all shortcode messages that matched the given filter
func (act Account) List(f ListFilter) (ResourceList, error) {
	var r ResourceList
	err := common.SendGetRequest(fmt.Sprintf(short.List+f.getQueryString(), act.AccountSid), act, &r)
	return r, err
}
