package usage

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var usage = struct {
	Get, Delete, Update, Create, List string
}{
	Get:    "/2010-04-01/Accounts/%s/Usage/Triggers/%s.json", // takes an AccountSid & TriggerSid
	Delete: "/2010-04-01/Accounts/%s/Usage/Triggers/%s.json", // takes an AccountSid & TriggerSid
	Update: "/2010-04-01/Accounts/%s/Usage/Triggers/%s.json", // takes an AccountSid & TriggerSid
	Create: "/2010-04-01/Accounts/%s/Usage/Triggers.json",    // takes an AccountSid
	List:   "/2010-04-01/Accounts/%s/Usage/Triggers.json",    // takes an AccountSid
}

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

var validateTriggerSid = regexp.MustCompile(`^UT[0-9a-z]{32}$`).MatchString

// Trigger represents a trigger that will cause Twilio to notify you.
//
// https://www.twilio.com/docs/api/rest/usage-triggers
type Trigger struct {
	common.ResourceInfo
	FriendlyName   string          `json:"friendly_name"`
	Recurring      string          `json:"recurring"`
	UsageCategory  string          `json:"usage_category"`
	TriggerBy      string          `json:"trigger_by"`
	TriggerValue   string          `json:"trigger_value"`
	CurrentValue   string          `json:"current_value"`
	UsageRecordURI string          `json:"usage_record_uri"`
	CallbackURL    string          `json:"callback_url"`
	CallbackMethod string          `json:"callback_method"`
	DateFired      common.JSONTime `json:"date_fired"`
}

// GetTrigger returns a Trigger object based on the sid.
func (act Account) GetTrigger(sid string) (Trigger, error) {
	var t Trigger
	if !validateTriggerSid(sid) {
		return t, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(usage.Get, act.AccountSid, sid), act, &t)
	return t, err
}

// DeleteTrigger removes a Twilio trigger.
func (act Account) DeleteTrigger(sid string) error {
	if !validateTriggerSid(sid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(usage.Delete, act.AccountSid, sid), act)
}

// TriggerUpdate is used to change properties of existing triggers.
type TriggerUpdate struct {
	FriendlyName   string
	CallbackURL    string
	CallbackMethod string
}

// GetReader implements the common.twilioPost interface
func (t TriggerUpdate) GetReader() io.Reader {
	v := url.Values{}
	if t.FriendlyName != "" {
		v.Add("FriendlyName", t.FriendlyName)
	}
	if t.CallbackURL != "" {
		v.Add("CallbackURL", t.CallbackURL)
	}
	if t.CallbackMethod != "" {
		v.Add("CallbackMethod", t.CallbackMethod)
	}
	return strings.NewReader(v.Encode())
}

// Validate implements the common.twilioPost interface
func (t TriggerUpdate) Validate() error {
	// TODO: bundled with top TODO
	return nil
}

// Update applies a trigger update.
func (act Account) Update(sid string, tu TriggerUpdate) (Trigger, error) {
	var t Trigger
	if !validateTriggerSid(sid) {
		return t, errors.New("Invalid sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(usage.Update, act.AccountSid, sid), tu, act, &t)
	return t, err
}

// TriggerListFilter is used to limit the responses of the TriggerList fn.
type TriggerListFilter struct {
	Recurring     string `json:"recurring"`
	UsageCategory string `json:"usage_category"`
	TriggerBy     string `json:"trigger_by"`
}

func (t TriggerListFilter) getQueryString() string {
	v := url.Values{}
	if t.Recurring != "" {
		v.Add("Recurring", t.Recurring)
	}
	if t.UsageCategory != "" {
		v.Add("UsageCategory", t.UsageCategory)
	}
	if t.TriggerBy != "" {
		v.Add("TriggerBy", t.TriggerBy)
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

// TriggerList contains a list of Triggers.
type TriggerList struct {
	common.ListResponseCore
	Triggers *[]Trigger `json:"triggers"`
	act      *Account
}

// ListTrigger returns a list of all triggers that satisfy the filter.
func (act Account) ListTrigger(f TriggerListFilter) (TriggerList, error) {
	var tl TriggerList
	err := common.SendGetRequest(fmt.Sprintf(usage.List, act.AccountSid)+f.getQueryString(), act, &tl)
	tl.act = &act
	return tl, err
}

// Next sets the MessageList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (tl *TriggerList) Next() error {
	if tl.Page == tl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(tl.NextPageURI, *tl.act, tl)
}

// NewTrigger is used to create a new Twilio usage trigger
type NewTrigger struct {
	UsageCategory string
	TriggerValue  string
	CallbackURL   string
	// Optional
	FriendlyName   string
	TriggerBy      string
	Recurring      string
	CallbackMethod string
}

// GetReader implements the common.twilioPost interface
func (nt NewTrigger) GetReader() io.Reader {
	v := url.Values{}
	if nt.UsageCategory != "" {
		v.Add("UsageCategory", nt.UsageCategory)
	}
	if nt.TriggerValue != "" {
		v.Add("TriggerValue", nt.TriggerValue)
	}
	if nt.CallbackURL != "" {
		v.Add("CallbackURL", nt.CallbackURL)
	}
	if nt.FriendlyName != "" {
		v.Add("FriendlyName", nt.FriendlyName)
	}
	if nt.TriggerBy != "" {
		v.Add("TriggerBy", nt.TriggerBy)
	}
	if nt.Recurring != "" {
		v.Add("Recurring", nt.Recurring)
	}
	if nt.CallbackMethod != "" {
		v.Add("CallbackMethod", nt.CallbackMethod)
	}
	return strings.NewReader(v.Encode())
}

// Validate implements the common.twilioPost interface
func (nt NewTrigger) Validate() error {
	errMsg := ""
	if nt.UsageCategory == "" {
		errMsg += `"UsageCategory" must be instantiated. `
	}
	if nt.TriggerValue == "" {
		errMsg += `"TriggerValue" must be instantiated. `
	}
	if nt.CallbackURL == "" {
		errMsg += `"CallbackURL" must be instantiated`
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}

// NewTrigger creates a new Twilio usage trigger.
func (act Account) NewTrigger(nt NewTrigger) (Trigger, error) {
	var t Trigger
	err := common.SendPostRequest(fmt.Sprintf(usage.Create, act.AccountSid), nt, act, &t)
	return t, err
}
