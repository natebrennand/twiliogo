package account

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var account = struct {
	Get, Update, List string
}{
	Get:    "/2010-04-01/Accounts/%s.json", // takes an AccountSid
	Update: "/2010-04-01/Accounts/%s.json", // takes an AccountSid
	List:   "/2010-04-01/Accounts.json",    // takes nothing
}

var validateAccountSid = regexp.MustCompile(`^AC[0-9a-z]{32}$`).MatchString

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
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

// Get retrieves the accunt resource
func (act Account) Get(sid string) (Resource, error) {
	var r Resource
	if !validateAccountSid(sid) {
		return r, errors.New("Invalid account sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(account.Get, sid), act, &r)
	return r, err
}

// private type so that it must be set using a pre-defined constant
type accountStatus string

// These constants are for changing the status of an account.
const (
	Closed    accountStatus = "closed"
	Suspended accountStatus = "suspended"
	Active    accountStatus = "active"
)

// Modification represents an update to an account's info.
type Modification struct {
	FriendlyName string
	Status       accountStatus
}

// Validate guarantees that a proposed account update is valid
func (m Modification) Validate() error {
	if len(m.FriendlyName) > 64 {
		return errors.New("Invalid FriendlyName, must be <= 64 characters")
	}
	return nil
}

// GetReader encodes the Post into an io.Reader for consumption while building a HTTP request to
// Twilio.
func (m Modification) GetReader() io.Reader {
	v := url.Values{}
	if m.FriendlyName != "" {
		v.Add("FriendlyName", m.FriendlyName)
	}
	if m.Status != "" {
		v.Add("Status", string(m.Status))
	}
	return strings.NewReader(v.Encode())
}

// Modify sends an update to the account for the info of an account. The new version of the
// account is returned.
func (act Account) Modify(m Modification) (Resource, error) {
	var r Resource
	err := common.SendPostRequest(fmt.Sprintf(account.Update, act.AccountSid), m, act, &r)
	return r, err
}

// ResourceList represents the list of all accounts controlled by the querying account.
type ResourceList struct {
	common.ListResponseCore
	Accounts *[]Resource `json:"accounts"`
	act      *Account
}

// Next sets the ResourceList to the next page of the list, returns an error in the
// case that there are no more pages left.
func (rl *ResourceList) Next() error {
	if rl.Page == rl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(rl.NextPageURI, *rl.act, rl)
}

// ListFilter allows filtering of accounts controlled by the querying account in a List()
// query.
type ListFilter struct {
	FriendlyName string
	Status       string
}

// generates a querystring to filter the accounts returned by List()
func (f ListFilter) getQueryString() string {
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

// List returns a list of all accounts that pass the filter.
func (act Account) List(f ListFilter) (ResourceList, error) {
	var rl ResourceList
	err := common.SendGetRequest(account.List+f.getQueryString(), act, &rl)
	rl.act = &act
	return rl, err
}
