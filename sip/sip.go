package sip

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strings"
)

var sip = struct {
	DomainList, Domain, ControlList, Control, CredentialList, Credential string
}{
	DomainList:     "/2010-04-01/Accounts/%s/SIP/Domains.json",                                   // takes an AccountSid
	Domain:         "/2010-04-01/Accounts/%s/SIP/Domains/%s.json",                                // takes an AccountSid & SipDomainSid
	ControlList:    "/2010-04-01/Accounts/%s/SIP/Domains/%s/IpAccessControlListMappings.json",    // takes an AccountSid & SipDomainSid
	Control:        "/2010-04-01/Accounts/%s/SIP/Domains/%s/IpAccessControlListMappings/%s.json", // takes AccountSid & SipDomainSid & ALSid
	CredentialList: "/2010-04-01/Accounts/%s/SIP/Domains/%s/CredentialListMappings.json",         // takes an AccountSid & SipDomainSid
	Credential:     "/2010-04-01/Accounts/%s/SIP/Domains/%s/CredentialListMappings/%s.json",      // takes AccountSid & SipDomainSid & CLSid
}

type Account struct {
	common.Account
}

var (
	validateDomainSid     = regexp.MustCompile("^SD[a-z0-9]{32}$").MatchString
	validateMappingSid    = regexp.MustCompile("^AL[a-z0-9]{32}$").MatchString
	validateCredentialSid = regexp.MustCompile("^CL[a-z0-9]{32}$").MatchString
)

type capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"sms"`
	MMS   bool `json:"mms"`
}

type subresourceURIs struct {
	IPAccessControlListMappings string `json:"ip_access_control_list_mappings"`
	CredentialListMappings      string `json:"credential_list_mappings"`
}

type addressesSubresourceURI struct {
	Addresses string `json:"addresses"`
}

type credentialSubresourceURI struct {
	Credentials string `json:"credentials"`
}

type DomainList struct {
	common.ListResponseCore
	SipDomains *[]Domain `json:"sip_domains"`
	act        *Account
}

type CredentialList struct {
	common.ListResponseCore
	CredentialListMappings *[]Credential `json:"credential_list_mappings"`
	act                    *Account
}

// Common is a struct that contains common fields between a Domain, Credential, and Mapping
type Common struct {
	Sid          string          `json:"sid"`
	AccountSid   string          `json:"account_sid"`
	FriendlyName string          `json:"friendly_name"`
	DateCreated  common.JSONTime `json:"date_created"`
	DateUpdated  common.JSONTime `json:"date_updated"`
	URI          string          `json:"uri"`
}

// IPAccessControlList is a struct that contains fields in an IpAccessControlList resource
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-properties
type IPAccessControlList struct {
	Common
	SubresourceURIs addressesSubresourceURI `json:"subresource_uris"`
}

// CredentialList is a struct that contains fields in a CredentialList resource
// https://www.twilio.com/docs/api/rest/credential-list#instance-properties
type CredentialList struct {
	Common
	SubresourceURIs credentialSubresourceURI `json:"subresource_uris"`
}

// Domain contains fields for a Domain resource
// Find details here: https://www.twilio.com/docs/api/rest/domain#instance-properties
type Domain struct {
	Common
	ApiVersion                string          `json:"api_version"`
	DomainName                string          `json:"domain_name"`
	AuthType                  string          `json:"auth_type"`
	VoiceURL                  string          `json:"voice_url"`
	VoiceMethod               string          `json:"voice_method"`
	VoiceFallbackURL          string          `json:"voice_fallback_url"`
	VoiceFallbackMethod       string          `json:"voice_fallback_method"`
	VoiceStatusCallbackURL    string          `json:"voice_status_callback_url"`
	VoiceStatusCallbackMethod string          `json:"voice_status_callback_method"`
	SubresourceURIs           subresourceURIs `json:"subresource_uris"`
}

// List grabs a list of all SIP domains for this account
// https://www.twilio.com/docs/api/rest/domain#list-get
func (act Account) ListDomains() (DomainList, error) {
	var dl DomainList
	err := common.SendGetRequest(fmt.Sprintf(sip.DomainList, act.AccountSid), act, &dl)
	dl.act = &act
	return dl, err
}

// Next sets the DomainList to the next page of the list resource, returns an error in the
// case that there are no more pages left
func (dl *DomainList) Next() error {
	if dl.Page == dl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(dl.NextPageURI, *dl.act, dl)
}

// NewDomain contains fields for creating a new sip domain
type NewDomain struct {
	DomainName                string
	FriendlyName              string
	VoiceURL                  string
	VoiceMethod               string
	VoiceFallbackURL          string
	VoiceFallbackMethod       string
	VoiceStatusCallbackURL    string
	VoiceStatusCallbackMethod string
}

// Create a new SIP domain which will be added to list of domains via a post
// https://www.twilio.com/docs/api/rest/domain#list-post
func (act Account) CreateDomain(n NewDomain) (Domain, error) {
	var d Domain
	err := common.SendPostRequest(fmt.Sprintf(sip.DomainList, act.AccountSid), n, act, &d)
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
	if n.VoiceURL != "" {
		vals.Set("VoiceUrl", n.VoiceURL)
	}
	if n.VoiceMethod != "" {
		vals.Set("VoiceMethod", n.VoiceMethod)
	}
	if n.VoiceFallbackURL != "" {
		vals.Set("VoiceFallbackUrl", n.VoiceFallbackURL)
	}
	if n.VoiceFallbackMethod != "" {
		vals.Set("VoiceFallbackMethod", n.VoiceFallbackMethod)
	}
	if n.VoiceStatusCallbackURL != "" {
		vals.Set("VoiceStatusCallbackUrl", n.VoiceStatusCallbackURL)
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
	return nil
}

// Domain gets a domain with a given SIP sid for this account
// https://www.twilio.com/docs/api/rest/domain#instance-get
func (act Account) GetDomain(domainSid string) (Domain, error) {
	var d Domain
	if !validateDomainSid(domainSid) {
		return d, errors.New("Invalid SIP sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(sip.Domain, act.AccountSid, domainSid), act, &d)
	return d, err
}

// Update a SIP domain with any parameters in a NewDomain
// https://www.twilio.com/docs/api/rest/domain#instance-post
func (act Account) UpdateDomain(n NewDomain, domainSid string) (Domain, error) {
	var d Domain
	if !validateDomainSid(domainSid) {
		return d, errors.New("Invalid SIP sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(sip.Domain, act.AccountSid, domainSid), n, act, &d)
	return d, err
}

// Delete a sip domain with the given SIP sid
// https://www.twilio.com/docs/api/rest/domain#instance-delete
func (act Account) DeleteDomain(domainSid string) error {
	if !validateDomainSid(domainSid) {
		return errors.New("Invalid SIP sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(sip.Domain, act.AccountSid, domainSid), act)
}

// Mapping gets a control list mapping for this sid
func (act Account) GetMapping(mappingSid, domainSid string) (IPAccessControlList, error) {
	var i IPAccessControlList
	if !validateMappingSid(mappingSid) {
		return i, errors.New("Invalid control sid")
	}
	if !validateDomainSid(domainSid) {
		return i, errors.New("Invalid domain sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(sip.Control, act.AccountSid, domainSid, mappingSid), act, &i)
	return i, err
}

// ControlListUpdate contains fields for updating a control list mapping
type ControlListUpdate struct {
	IpAccessControlListSid string
}

// GetReader is needed for the common.twilioPost interface
func (c ControlListUpdate) GetReader() io.Reader {
	vals := url.Values{}
	if c.IpAccessControlListSid != "" {
		vals.Set("IpAccessControlListSid", c.IpAccessControlListSid)
	}
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (c ControlListUpdate) Validate() error {
	if c.IpAccessControlListSid == "" {
		return errors.New("Must include an ip access control list sid")
	}
	return nil
}

// Add a mapping to the control list on this domain
// https://www.twilio.com/docs/api/rest/domain#subresource-list-post-ipacl
func (act Account) AddMapping(c ControlListUpdate, domainSid string) (Mapping, error) {
	var m Mapping
	if !validateDomainSid(domainSid) {
		return m, errors.New("Invalid sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(sip.ControlList, act.AccountSid, domainSid), c, act, &m)
	return m, err
}

// Delete a mapping with the given sid from this domain
// https://www.twilio.com/docs/api/rest/domain#subresource-list-delete-ipacl
func (act Account) DeleteMapping(domainSid, mappingSid string) error {
	if !validateMappingSid(mappingSid) {
		return errors.New("Invalid control sid")
	}
	if !validateDomainSid(domainSid) {
		return errors.New("Invalid domain sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(sip.Control, act.AccountSid, domainSid, mappingSid), act)
}

// Next sets the CredentialList to the next page of the list resource, returns an error in the
// case that there are no more pages left
func (cl *CredentialList) Next() error {
	if cl.Page == cl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(cl.NextPageURI, *cl.act, cl)
}

// List grabs a list of all credential mappings for this account and domain
// https://www.twilio.com/docs/api/rest/domain#list-get-clm
func (act Account) ListCredentials(domainSid string) (CredentialList, error) {
	var cl CredentialList
	if !validateDomainSid(domainSid) {
		return cl, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(sip.CredentialList, act.AccountSid, domainSid), act, &cl)
	cl.act = &act
	return cl, err
}

// CredentialListUpdate contains fields for adding a credential
type CredentialListUpdate struct {
	CredentialListSid string
}

// GetReader is needed for the common.twilioPost interface
func (c CredentialListUpdate) GetReader() io.Reader {
	vals := url.Values{}
	if c.CredentialListSid != "" {
		vals.Set("CredentialListSid", c.CredentialListSid)
	}
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (c CredentialListUpdate) Validate() error {
	if c.CredentialListSid == "" {
		return errors.New("Must include a credential list sid")
	}
	return nil
}

// Add a credential list to the domain
// https://www.twilio.com/docs/api/rest/domain#list-post-clm
func (act Account) AddCredential(u CredentialListUpdate, domainSid string) (Credential, error) {
	var c Credential
	if !validateDomainSid(domainSid) {
		return c, errors.New("Invalid domain sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(sip.CredentialList, act.AccountSid, domainSid), u, act, &c)
	return c, err
}

// Delete a credential with the given sid from this domain
// https://www.twilio.com/docs/api/rest/domain#list-delete-clm
func (act Account) DeleteCredential(domainSid, credentialSid string) error {
	if !validateCredentialSid(credentialSid) {
		return errors.New("Invalid credential sid")
	}
	if !validateDomainSid(domainSid) {
		return errors.New("Invalid domain sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(sip.Credential, act.AccountSid, domainSid, credentialSid), act)
}
