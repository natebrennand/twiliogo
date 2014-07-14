package usage

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
	getTriggerURL    = "https://api.twilio.com/2010-04-01/Accounts/%s/Usage/Triggers/%s.json" // takes an AccountSid & TriggerSid
	deleteTriggerURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Usage/Triggers/%s.json" // takes an AccountSid & TriggerSid
	updateTriggerURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Usage/Triggers/%s.json" // takes an AccountSid & TriggerSid
	createTriggerURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Usage/Triggers.json"    // takes an AccountSid
	listTriggerURL   = "https://api.twilio.com/2010-04-01/Accounts/%s/Usage/Triggers.json"    // takes an AccountSid
)

// TODO: add checks for values that have a limited number of constants

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

func validateTriggerSid(sid string) bool {
	match, _ := regexp.MatchString(`^UT[0-9a-z]{32}$`, sid)
	return match
}

type Trigger struct {
	Sid            string          `json:"sid"`
	AccountSid     string          `json:"account_sid"`
	FriendlyName   string          `json:"friendly_name"`
	Recurring      string          `json:"recurring"`
	UsageCategory  string          `json:"usage_category"`
	TriggerBy      string          `json:"trigger_by"`
	TriggerValue   string          `json:"trigger_value"`
	CurrentValue   string          `json:"current_value"`
	UsageRecordURI string          `json:"usage_record_uri"`
	CallbackURL    string          `json:"callback_url"`
	CallbackMethod string          `json:"callback_method"`
	URI            string          `json:"uri"`
	DateFired      common.JSONTime `json:"date_fired"`
	DateCreated    common.JSONTime `json:"date_created"`
	DateUpdated    common.JSONTime `json:"date_updated"`
}

func (act Account) GetTrigger(sid string) (Trigger, error) {
	var t Trigger
	if !validateTriggerSid(sid) {
		return t, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getTriggerURL, act.AccountSid, sid), act, &t)
	return t, err
}

func (act Account) DeleteTrigger(sid string) error {
	if !validateTriggerSid(sid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(deleteTriggerURL, act.AccountSid, sid), act)
}

type TriggerUpdate struct {
	FriendlyName   string
	CallbackURL    string
	CallbackMethod string
}

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

func (t TriggerUpdate) Validate() error {
	// TODO: bundled with top TODO
	return nil
}

func (act Account) Update(sid string, tu TriggerUpdate) (Trigger, error) {
	var t Trigger
	if !validateTriggerSid(sid) {
		return t, errors.New("Invalid sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(updateTriggerURL, act.AccountSid, sid), tu, act, &t)
	return t, err
}

type TriggerListFilter struct {
	Recurring     string `json:"recurring"`
	UsageCategory string `json:"usage_category"`
	TriggerBy     string `json:"trigger_by"`
}

func (t TriggerListFilter) GetQueryString() string {
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

type TriggerList struct {
	common.ListResponseCore
	Triggers *[]Trigger `json:"triggers"`
}

func (act Account) ListTrigger(f TriggerListFilter) (TriggerList, error) {
	var tl TriggerList
	err := common.SendGetRequest(fmt.Sprintf(listTriggerURL, act.AccountSid)+f.GetQueryString(), act, &tl)
	return tl, err
}

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

func (act Account) NewTrigger(nt NewTrigger) (Trigger, error) {
	var t Trigger
	err := common.SendPostRequest(fmt.Sprintf(createTriggerURL, act.AccountSid), nt, act, &t)
	return t, err
}
