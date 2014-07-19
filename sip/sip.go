package sip

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
)

var sip = struct {
	All, Domain, ControlList, Control, CredentialList, Credential string
}{
	All:            "/2010-04-01/Accounts/%s/SIP/Domains.json",                                   // takes an AccountSid
	Domain:         "/2010-04-01/Accounts/%s/SIP/Domains/%s.json",                                // takes an AccountSid & SipDomainSid
	ControlList:    "/2010-04-01/Accounts/%s/SIP/Domains/%s/IpAccessControlListMappings.json",    // takes an AccountSid & SipDomainSid
	Control:        "/2010-04-01/Accounts/%s/SIP/Domains/%s/IpAccessControlListMappings/%s.json", // takes AccountSid & SipDomainSid & ALSid
	CredentialList: "/2010-04-01/Accounts/%s/SIP/Domains/%s/CredentialListMappings.json",         // takes an AccountSid & SipDomainSid
	Credential:     "/2010-04-01/Accounts/%s/SIP/Domains/%s/CredentialListMappings/%s.json",      // takes AccountSid & SipDomainSid & CLSid
}

type Account struct {
	common.Account
}

type capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"sms"`
	MMS   bool `json:"mms"`
}

type DomainList struct {
	common.ListResponseCore
	SipDomains *[]Domain `json:"sip_domains"`
}

type Domain struct {
	Sid                       string          `json:"sid"`
	FriendlyName              string          `json:"friendly_name"`
	AccountSid                string          `json:"account_sid"`
	ApiVersion                string          `json:"api_version"`
	DomainName                string          `json:"domain_name"`
	AuthType                  string          `json:"auth_type"`
	VoiceUrl                  string          `json:"voice_url"`
	VoiceMethod               string          `json:"voice_method"`
	VoiceFallbackUrl          string          `json:"voice_fallback_url"`
	VoiceFallbackMethod       string          `json:"voice_fallback_method"`
	VoiceStatusCallback       string          `json:"voice_status_callback"`
	VoiceStatusCallbackMethod string          `json:"voice_status_callback_method"`
	DateCreated               common.JSONTime `json:"date_created"`
	DateUpdated               common.JSONTime `json:"date_updated"`
	Uri                       string          `json:"uri"`
}

// List grabs a list of all SIP domains for this account
func (act Account) List() (DomainList, error) {
	var dl DomainList
	err := common.SendGetRequest(fmt.Sprintf(sip.All, act.AccountSid), act, &dl)
	dl.act = &act
	return dl, err
}

type NewDomain struct {
	DomainName                string
	FriendlyName              string
	VoiceUrl                  string
	VoiceMethod               string
	VoiceFallbackUrl          string
	VoiceFallbackMethod       string
	VoiceStatusCallback       string
	VoiceStatusCallbackMethod string
}

// Create a new SIP domain which will be added to list of domains via a post
func (act Account) Create(n NewDomain) (Domain, err) {
	var d Domain
	err := common.SendPostRequest(fmt.Sprintf(sip.All, act.AccountSid), n, act, &d)
	return d, err
}

// GetReader is needed for the common.twilioPost interface
func (n NewDomain) GetReader() io.Reader {
	vals := url.Values{}
	if n.DomainName != "" {
		vals.Set("DomainName", n.DomainName)
	}
	if n.FriendlyName != "" {
		vals.Set("FriendlyName", n.FriendlyName)
	}
	if n.VoiceUrl != "" {
		vals.Set("VoiceUrl", n.VoiceUrl)
	}
	if n.VoiceMethod != "" {
		vals.Set("VoiceMethod", n.VoiceMethod)
	}
	if n.VoiceFallbackUrl != "" {
		vals.Set("VoiceFallbackUrl", n.VoiceFallbackUrl)
	}
	if n.VoiceFallbackMethod != "" {
		vals.Set("VoiceFallbackMethod", n.VoiceFallbackMethod)
	}
	if n.VoiceStatusCallback != "" {
		vals.Set("VoiceStatusCallback", n.VoiceStatusCallback)
	}
	if n.VoiceStatusCallbackMethod != "" {
		vals.Set("VoiceStatusCallbackMethod", n.VoiceStatusCallbackMethod)
	}
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (n NewDomain) Validate() error {
	if n.DomainName == "" {
		return errors.New("Must include at least a domain name for your new domain")
	}
	return nil // All params are optional
}
