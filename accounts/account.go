package account

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	getURL    = "https://api.twilio.com/2010-04-01/Accounts/%s.json" // takes an AccountSid
	updateURL = "https://api.twilio.com/2010-04-01/Accounts/%s.json" // takes an AccountSid
	listURL   = "https://api.twilio.com/2010-04-01/Accounts.json"    // takes nothing
)

var accountStatuses = map[string]bool{
	"closed":    true,
	"suspended": true,
	"active":    true,
}

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

func validateAccountSid(sid string) bool {
	match, _ := regexp.MatchString(`^AC[0-9a-z]{32}$`, string(sid))
	return match
}

// Resource represents an Account resource
//
// https://www.twilio.com/docs/api/rest/account
type Resource struct {
	Sid             string          `json:"sid"`
	DateCreated     common.JSONTime `json:"date_created"`
	DateUpdated     common.JSONTime `json:"date_updated"`
	FriendlyName    string          `json:"friendly_name"`
	Type            string          `json:"type"`
	Status          string          `json:"status"`
	AuthToken       string          `json:"auth_token"`
	URI             string          `json:"uri"`
	SubresourceURIs struct {
		AvailablePhoneNumbers string `json:"available_phone_numbers"`
		Calls                 string `json:"calls"`
		Conferences           string `json:"conferences"`
		IncomingPhoneNumbers  string `json:"incoming_phone_numbers"`
		Notifications         string `json:"notifications"`
		OutgoingCallerIds     string `json:"outgoing_caller_ids"`
		Recordings            string `json:"recordings"`
		Sandbox               string `json:"sandbox"`
		SmsPessages           string `json:"sms_messages"`
		Transcriptions        string `json:"transcriptions"`
	} `json:"subresource_uris"`
	OwnerAccountSid string `json:"owner_account_sid"`
}

func (act Account) Get() (Resource, error) {
	var r Resource
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid), act, &r, 200)
	return r, err
}

type Modification struct {
	FriendlyName string
	Status       string
}

func (m Modification) Validate() error {
	if len(m.FriendlyName) > 64 {
		return errors.New("Invalid FriendlyName, must be <= 64 characters")
	}
	if m.Status != "" {
		_, exists := accountStatuses[m.Status]
		if !exists {
			return errors.New("Invalid updated status for account")
		}
	}
	return nil
}

func (m Modification) GetReader() io.Reader {
	v := url.Values{}
	if m.FriendlyName != "" {
		v.Add("FriendlyName", m.FriendlyName)
	}
	if m.Status != "" {
		v.Add("Status", m.Status)
	}
	return strings.NewReader(v.Encode())
}

func (act Account) Modify(m Modification) (Resource, error) {
	var r Resource
	err := common.SendPostRequest(fmt.Sprintf(updateURL, act.AccountSid), m, act, &r, 200)
	return r, err
}

type ResourceList struct {
	common.ListResponseCore
	Accounts *[]Resource `json:"accounts"`
}

type ListFilter struct {
	FriendlyName string
	Status       string
}

func (f ListFilter) GetQueryString() string {
	v := url.Values{}
	if f.FriendlyName != "" {
		v.Add("FriendlyName", f.FriendlyName)
	}
	if f.Status != "" {
		v.Add("Status", f.Status)
	}
	qs := v.Encode()
	if qs != "" {
		return "?" + qs
	}
	return qs
}

func (act Account) List(f ListFilter) (ResourceList, error) {
	var rl ResourceList
	err := common.SendGetRequest(listURL+f.GetQueryString(), act, &rl, 200)
	return rl, err
}
