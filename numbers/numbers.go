package numbers

// https://www.twilio.com/docs/api/rest/incoming-phone-numbers

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

var numbers = struct {
	Get, List, Local, TollFree, Mobile string
}{
	Get:      "/2010-04-01/Accounts/%s/IncomingPhoneNumbers/%s.json",       // takes act sid and incoming phone # sid
	List:     "/2010-04-01/Accounts/%s/IncomingPhoneNumbers.json",          // takes act sid
	Local:    "/2010-04-01/Accounts/%s/IncomingPhoneNumbers/Local.json",    // takes act sid
	TollFree: "/2010-04-01/Accounts/%s/IncomingPhoneNumbers/TollFree.json", // takes act sid
	Mobile:   "/2010-04-01/Accounts/%s/IncomingPhoneNumbers/Mobile.json",   // takes act sid
}

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

// core contains attributes common to Number, Update, NumberSelector & Selector
type core struct {
	FriendlyName         string `json:"friendly_name"`
	APIVersion           string `json:"api_version"`
	VoiceURL             string `json:"voice_url"`
	VoiceMethod          string `json:"voice_method"`
	VoiceFallbackURL     string `json:"voice_fallback_url"`
	VoiceFallbackMethod  string `json:"voice_fallback_method"`
	StatusCallback       string `json:"status_callback"`
	StatusCallbackMethod string `json:"status_callback_method"`
	VoiceCallerIDLookup  bool   `json:"voice_caller_id_lookup,string"`
	VoiceApplicationSid  string `json:"voice_application_sid"`
	SmsURL               string `json:"sms_url"`
	SmsMethod            string `json:"sms_method"`
	SmsFallbackURL       string `json:"sms_fallback_url"`
	SmsFallbackMethod    string `json:"sms_fallback_method"`
	SmsApplicationSid    string `json:"sms_application_sid"`
}

func setCoreValues(p core) url.Values {
	vals := url.Values{}
	if p.FriendlyName != "" {
		vals.Set("FriendlyName", p.FriendlyName)
	}
	if p.APIVersion != "" {
		vals.Set("ApiVersion", p.APIVersion)
	}
	if p.VoiceURL != "" {
		vals.Set("VoiceUrl", p.VoiceURL)
	}
	if p.VoiceMethod != "" {
		vals.Set("VoiceMethod", p.VoiceMethod)
	}
	if p.VoiceFallbackURL != "" {
		vals.Set("VoiceFallbackUrl", p.VoiceFallbackURL)
	}
	if p.VoiceFallbackMethod != "" {
		vals.Set("VoiceFallbackMethod", p.VoiceFallbackMethod)
	}
	if p.StatusCallback != "" {
		vals.Set("StatusCallback", p.StatusCallback)
	}
	if p.StatusCallbackMethod != "" {
		vals.Set("StatusCallbackMethod", p.StatusCallbackMethod)
	}
	vals.Set("VoiceCallerIdLookup", strconv.FormatBool(p.VoiceCallerIDLookup))
	if p.VoiceApplicationSid != "" {
		vals.Set("ApplicationSid", p.VoiceApplicationSid)
	}
	if p.SmsURL != "" {
		vals.Set("SmsUrl", p.SmsURL)
	}
	if p.SmsMethod != "" {
		vals.Set("SmsMethod", p.SmsMethod)
	}
	if p.SmsFallbackURL != "" {
		vals.Set("SmsFallbackUrl", p.SmsFallbackURL)
	}
	if p.SmsFallbackMethod != "" {
		vals.Set("SmsFallbackMethod", p.SmsFallbackMethod)
	}
	if p.SmsApplicationSid != "" {
		vals.Set("SmsApplicationSid", p.SmsApplicationSid)
	}
	return vals
}

var validateNumberSid = regexp.MustCompile(`^PN[0-9a-z]{32}$`).MatchString

// Number represent a number that your account is in control of.
type Number struct {
	core
	Sid          string          `json:"sid"`
	DateCreated  common.JSONTime `json:"date_created"`
	DateUpdated  common.JSONTime `json:"date_updated"`
	PhoneNumber  string          `json:"phone_number"`
	Capabilities capabilities    `json:"capabilities"`
	URI          string          `json:"uri"`
	AccountSid   string          `json:"account_sid"`
}

// helper struct
type capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"sms"`
	MMS   bool `json:"mms"`
}

// Get a info about a number with it's NumberSid
func (act Account) Get(numberSid string) (Number, error) {
	var p Number
	if !validateNumberSid(numberSid) {
		return p, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(numbers.Get, act.AccountSid, numberSid), act, &p)
	return p, err
}

// Update is used to apply an update to a current Number resource.
//
// https://www.twilio.com/docs/api/rest/incoming-phone-numbers#instance-post
type Update struct {
	core
	AccountSid string
}

// Update make modifications to properties of a phone number via a post
func (act Account) Update(numberSid string, u Update) (Number, error) {
	var p Number
	if !validateNumberSid(numberSid) {
		return p, errors.New("Invalid sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(numbers.Get, act.AccountSid, numberSid), u, act, &p)
	return p, err
}

// GetReader is needed for the common.twilioPost interface
func (p Update) GetReader() io.Reader {
	vals := setCoreValues(p.core)
	if p.AccountSid != "" {
		vals.Set("AccountSid", p.AccountSid)
	}
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (p Update) Validate() error {
	return nil // All params are optional
}

// Delete a number from your account
//
// https://www.twilio.com/docs/api/rest/incoming-phone-numbers#instance-delete
func (act Account) Delete(numberSid string) error {
	if !validateNumberSid(numberSid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(numbers.Get, act.AccountSid, numberSid), act)
}

// NumberSelector is used to purchase a new phone number for an account.
//
// https://www.twilio.com/docs/api/rest/incoming-phone-numbers#list-post
type NumberSelector struct {
	core
	PhoneNumber string
	AreaCode    string
}

// GetReader implemented for the common.twillioPost interface
func (n NumberSelector) GetReader() io.Reader {
	vals := setCoreValues(n.core)
	if n.PhoneNumber != "" {
		vals.Set("PhoneNumber", n.PhoneNumber)
	}
	if n.AreaCode != "" {
		vals.Set("AreaCode", n.AreaCode)
	}
	return strings.NewReader(vals.Encode())
}

// Validate implemented for the common.twillioPost interface
func (n NumberSelector) Validate() error {
	if n.PhoneNumber == "" && n.AreaCode == "" {
		return errors.New("Must set either the phone number or area code")
	} else if n.PhoneNumber != "" && n.AreaCode != "" {
		return errors.New("Cannot set both the phone number and area code")
	}
	return nil
}

// PurchaseNumber executes a purchase request directed by a NumberSelector.
func (act Account) PurchaseNumber(n NumberSelector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(numbers.List, act.AccountSid), n, act, &p)
	return p, err
}

// Selector used to detail a phone number purchase for a specific category.
type Selector struct {
	core
	PhoneNumber string
}

// GetReader implemented for the common.twillioPost interface
func (s Selector) GetReader() io.Reader {
	vals := setCoreValues(s.core)
	if s.PhoneNumber != "" {
		vals.Set("PhoneNumber", s.PhoneNumber)
	}
	return strings.NewReader(vals.Encode())
}

// Validate implemented for the common.twillioPost interface
func (s Selector) Validate() error {
	if s.PhoneNumber == "" {
		return errors.New("Must set phone number to purchase")
	}
	return nil
}

// PurchaseLocal buys a local number if an appropriate number is found
func (act Account) PurchaseLocal(n Selector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(numbers.Local, act.AccountSid), n, act, &p)
	return p, err
}

// PurchaseTollFree buys a tollfree number if an appropriate number is found
func (act Account) PurchaseTollFree(n Selector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(numbers.TollFree, act.AccountSid), n, act, &p)
	return p, err
}

// PurchaseMobile buys a mobile number if an appropriate number is found
func (act Account) PurchaseMobile(n Selector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(numbers.Mobile, act.AccountSid), n, act, &p)
	return p, err
}

// ListFilter provides a way to filter the response from List() to a subset of your numbers.
type ListFilter struct {
	PhoneNumber  string
	FriendlyName string
}

func (f ListFilter) getQueryString() string {
	v := url.Values{}
	if f.PhoneNumber != "" {
		v.Set("PhoneNumber", f.PhoneNumber)
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

// NumberList represents the list of numbers the account controls.
type NumberList struct {
	common.ListResponseCore
	act                  *Account
	IncomingPhoneNumbers *[]Number `json:"incoming_phone_numbers"`
}

// Next sets the NumberList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (nl *NumberList) Next() error {
	if nl.Page == nl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(nl.NextPageURI, *nl.act, nl)
}

// List grabs a list of phone numbers for a given account with optional filters
func (act Account) List(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(numbers.List, act.AccountSid)+f.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}

// ListLocal grabs a list of local phone numbers for a given account with optional filters - no toll free
func (act Account) ListLocal(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(numbers.Local, act.AccountSid)+f.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}

// ListTollFree grabs a list of toll free phone numbers for a given account with optional filters - no toll free
func (act Account) ListTollFree(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(numbers.TollFree, act.AccountSid)+f.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}

// ListMobile grabs a list of mobile phone numbers for a given account with optional filters - no toll free
func (act Account) ListMobile(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(numbers.Mobile, act.AccountSid)+f.getQueryString(), act, &nl)
	nl.act = &act
	return nl, err
}
