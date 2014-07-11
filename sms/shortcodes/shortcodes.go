package shortcodes

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"net/url"
	"regexp"
)

const (
	getURL  = "https://api.twilio.com/2010-04-01/Accounts/%s/SMS/ShortCodes/%s.json" // takes an AccountSid & ShortcodeSid
	listURL = "https://api.twilio.com/2010-04-01/Accounts/%s/SMS/ShortCodes.json"    // takes an AccountSid

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

func validateShortcodeSid(sid string) bool {
	match, _ := regexp.MatchString(`^SC[0-9a-z]{32}$`, string(sid))
	return match
}

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

func (act Account) Get(sid string) (Resource, error) {
	var r Resource
	if !validateShortcodeSid(sid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, sid), act, &r)
	return r, err
}

type ListFilter struct {
	ShortCode    string
	FriendlyName string
}

func (f ListFilter) GetQueryString() string {
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

type ResourceList struct {
	common.ListResponseCore
	ShortCodes *[]Resource
}

func (act Account) GetList(f ListFilter) (ResourceList, error) {
	var r ResourceList
	err := common.SendGetRequest(fmt.Sprintf(listURL+f.GetQueryString(), act.AccountSid), act, &r)
	return r, err
}
