package numbers

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

type Core struct {
	FriendlyName         string
	APIVersion           string
	VoiceURL             string
	VoiceMethod          string
	VoiceFallbackURL     string
	VoiceFallbackMethod  string
	StatusCallback       string
	StatusCallbackMethod string
	VoiceCallerIdLookup  *bool
	VoiceApplicationSid  string
	SmsURL               string
	SmsMethod            string
	SmsFallbackURL       string
	SmsFallbackMethod    string
	SmsApplicationSid    string
}

type Post struct {
	Core
	AccountSid string
}

type NumberSelector struct {
	Core
	PhoneNumber string
	AreaCode    string
}

type Selector struct {
	Core
	PhoneNumber string
}

func setCoreValues(p Core) url.Values {
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
	if p.VoiceCallerIdLookup != nil {
		vals.Set("VoiceCallerIdLookup", strconv.FormatBool(*p.VoiceCallerIdLookup))
	}
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

func (p Post) GetReader() io.Reader {
	vals := setCoreValues(p.Core)
	if p.AccountSid != "" {
		vals.Set("AccountSid", p.AccountSid)
	}
	return strings.NewReader(vals.Encode())
}

func (n NumberSelector) GetReader() io.Reader {
	vals := setCoreValues(n.Core)
	if n.PhoneNumber != "" {
		vals.Set("PhoneNumber", n.PhoneNumber)
	}
	if n.AreaCode != "" {
		vals.Set("AreaCode", n.AreaCode)
	}
	return strings.NewReader(vals.Encode())

}

func (s Selector) GetReader() io.Reader {
	vals := setCoreValues(s.Core)
	if s.PhoneNumber != "" {
		vals.Set("PhoneNumber", s.PhoneNumber)
	}

	return strings.NewReader(vals.Encode())

}

func (p Post) Validate() error {
	// All params are optional
	return nil
}

func (n NumberSelector) Validate() error {
	if n.PhoneNumber == "" && n.AreaCode == "" {
		return errors.New("Must set either the phone number or area code")
	}
	return nil
}

func (p Selector) Validate() error {
	// All params are optional
	if p.PhoneNumber == "" {
		return errors.New("Must set phone number to purchase")
	}
	return nil
}

// Get a info about a conference with confSid
func (act Account) Get(numberSid string) (Number, error) {
	var p Number
	if !validateNumberSid(numberSid) {
		return p, errors.New("Invalid sid")
	}

	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, numberSid), act, &p)
	return p, err
}

// Updates properties of a phone number via a post
func (act Account) Post(numberSid string, update Post) (Number, error) {
	var p Number
	if !validateNumberSid(numberSid) {
		return p, errors.New("Invalid sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(getURL, act.AccountSid, numberSid), update, act, &p)
	return p, err
}

func (act Account) PurchaseNumber(n NumberSelector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(listURL, act.AccountSid), n, act, &p)
	return p, err
}

// Updates properties of a phone number via a put
func (act Account) Put(numberSid string, update Post) (Number, error) {
	var p Number
	if !validateNumberSid(numberSid) {
		return p, errors.New("Invalid sid")
	}
	err := common.SendPutRequest(fmt.Sprintf(getURL, act.AccountSid, numberSid), update, act, &p)
	return p, err
}

// Delete a number from your account
func (act Account) Delete(numberSid string) error {
	if !validateNumberSid(numberSid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(getURL, act.AccountSid, numberSid), act)
}

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

// Grabs a list of phone numbers for a given account with optional filters
func (act Account) List(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid)+f.getQueryString(), act, &nl)
	return nl, err
}

// Grabs a list of local phone numbers for a given account with optional filters - no toll free
// Not sure if can filter
func (act Account) ListLocal(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(localURL, act.AccountSid)+f.getQueryString(), act, &nl)
	return nl, err
}

// Posts to list of local numbers if an appropriate number is found
func (act Account) PurchaseLocal(n Selector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(localURL, act.AccountSid), n, act, &p)
	return p, err
}

// Grabs a list of toll free phone numbers for a given account with optional filters - no toll free
// Not sure if can filter
func (act Account) ListTollFree(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(tollFreeURL, act.AccountSid)+f.getQueryString(), act, &nl)
	return nl, err
}

// Posts to list of toll free numbers if an appropriate number is found
func (act Account) PurchaseTollFree(n Selector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(tollFreeURL, act.AccountSid), n, act, &p)
	return p, err
}

// Grabs a list of mobile phone numbers for a given account with optional filters - no toll free
// Not sure if can filter
func (act Account) ListMobile(f ListFilter) (NumberList, error) {
	var nl NumberList
	err := common.SendGetRequest(fmt.Sprintf(mobileURL, act.AccountSid)+f.getQueryString(), act, &nl)
	return nl, err
}

// Posts to list of mobile numbers if an appropriate number is found
func (act Account) PurchaseMobile(n Selector) (Number, error) {
	var p Number
	err := common.SendPostRequest(fmt.Sprintf(mobileURL, act.AccountSid), n, act, &p)
	return p, err
}
