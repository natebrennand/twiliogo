package numbers

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

type Number struct {
	Sid                  string          `json:"sid"`
	DateCreated          common.JSONTime `json:"date_created"`
	DateUpdated          common.JSONTime `json:"date_updated"`
	PhoneNumber          string          `json:"phone_number"`
	Capabilities         []string        `json:"capabilities"`
	URI                  string          `json:"uri"`
	FriendlyName         string          `json:"friendly_name"`
	APIVersion           string          `json:"api_version"`
	VoiceURL             string          `json:"voice_url"`
	VoiceMethod          string          `json:"voice_method"`
	VoiceFallbackURL     string          `json:"voice_fallback_url"`
	VoiceFallbackMethod  string          `json:"voice_fallback_method"`
	StatusCallback       string          `json:"status_callback"`
	StatusCallbackMethod string          `json:"status_callback_method"`
	VoiceCallerIdLookup  bool            `json:"voice_caller_id_lookup,string"`
	VoiceApplicationSid  string          `json:"voice_application_sid"`
	SmsURL               string          `json:"sms_url"`
	SmsMethod            string          `json:"sms_method"`
	SmsFallbackURL       string          `json:"sms_fallback_url"`
	SmsFallbackMethod    string          `json:"sms_fallback_method"`
	SmsApplicationSid    string          `json:"sms_application_sid"`
	AccountSid           string          `json:"account_sid"`
}

type NumberList struct {
	common.ListResponseCore
	IncomingPhoneNumbers *[]Number `json:"incoming_phone_numbers"`
}

var validateNumberSid = regexp.MustCompile(`^PN[0-9a-z]{32}$`).MatchString(sid)
