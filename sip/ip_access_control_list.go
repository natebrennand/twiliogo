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

var ipAccessControlList = struct {
	ControlLists, ControlList, IPAddresses, IPAddress string
}{
	ControlLists: "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists.json",                   // takes an AccountSid
	ControlList:  "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists/%s.json",                // takes an AccountSid & IpAccessControlListSid
	IPAddresses:  "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists/%s/IpAddresses.json",    // takes an AccountSid & IpAccessControlListSid
	IPAddress:    "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists/%s/IpAddresses/%s.json", // takes an AccountSid & IpAccessControlListSid & IpAddressSid
}

var (
	validateIPSid = regexp.MustCompile("^IP[a-z0-9]{32}$").MatchString
)

// IPAccessControlList represents an IPAccessControlList
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-properties
type IPAccessControlList struct {
	common.ResourceInfo
	SubresourceURIs addressesSubresourceURI `json:"subresource_uris"`
	FriendlyName    string                  `json:"friendly_name"`
}

// IPAccessControlLists represents a list of IPAccessControlList
type IPAccessControlLists struct {
	common.ListResponseCore
	ControlLists *[]IPAccessControlList `json:"ip_access_control_lists"`
	act          *Account
}

// Next sets the IPAccessControlLists to the next page of the list resource, returns an error in the
// case that there are no more pages left
func (il *IPAccessControlLists) Next() error {
	if il.Page == il.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(il.NextPageURI, *il.act, il)
}

// UpdateControlList contains properties to add or edit an IPAccessControlList on your account
type UpdateControlList struct {
	FriendlyName string
}

// GetReader is needed for the common.twilioPost interface
func (u UpdateControlList) GetReader() io.Reader {
	vals := url.Values{}
	if u.FriendlyName != "" {
		vals.Set("FriendlyName", u.FriendlyName)
	}
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (u UpdateControlList) Validate() error {
	if u.FriendlyName == "" {
		return errors.New("Must include at least a friendly name for your new control list")
	}
	return nil
}

// ListControlLists grabs a list of all IPAccessControlList for this account
// https://www.twilio.com/docs/api/rest/ip-access-control-list#list
func (act Account) ListControlLists() (IPAccessControlLists, error) {
	var il IPAccessControlLists
	err := common.SendGetRequest(fmt.Sprintf(ipAccessControlList.ControlLists, act.AccountSid), act, &il)
	il.act = &act
	return il, err
}

// AddNewControlList allows you to add a new IPAccessControlList to your account
// https://www.twilio.com/docs/api/rest/ip-access-control-list#list-post
func (act Account) AddNewControlList(u UpdateControlList) (IPAccessControlList, error) {
	var i IPAccessControlList
	err := common.SendPostRequest(fmt.Sprintf(ipAccessControlList.ControlLists, act.AccountSid), u, act, &i)
	return i, err
}

// GetControlList allows you to get an IPAccessControlList
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-get
func (act Account) GetControlList(alSid string) (IPAccessControlList, error) {
	var i IPAccessControlList
	if !validateMappingSid(alSid) {
		return i, errors.New("Invalid ip access control sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(ipAccessControlList.ControlList, act.AccountSid, alSid), act, i)
	return i, err
}

// UpdateControlList allows you to update an IPAccessControlList with a new FriendlyName
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-post
func (act Account) UpdateControlList(u UpdateControlList, alSid string) (IPAccessControlList, error) {
	var i IPAccessControlList
	if !validateMappingSid(alSid) {
		return i, errors.New("Invalid ip access control sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(ipAccessControlList.ControlList, act.AccountSid, alSid), u, act, &i)
	return i, err
}

// DeleteControlList allows you to delete an IPAccessControlList from your account
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-delete
func (act Account) DeleteControlList(alSid string) error {
	if !validateMappingSid(alSid) {
		return errors.New("Invalid ip access control sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(ipAccessControlList.ControlList, act.AccountSid, alSid), act)
}

// IPAddressResource represents an IPAddress
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-properties-ipaddress
type IPAddressResource struct {
	common.ResourceInfo
	FriendlyName string `json:"friendly_name"`
	IPAddress    string `json:"ip_address"`
}

// IPAddressList represents a list of IPAccessControlList
type IPAddressList struct {
	common.ListResponseCore
	IPAddresses *[]IPAddressResource `json:"ip_addresses"`
	act         *Account
}

// Next sets the IPAddressList to the next page of the list resource, returns an error in the
// case that there are no more pages left
func (il *IPAddressList) Next() error {
	if il.Page == il.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(il.NextPageURI, *il.act, il)
}

// ListIPAddresses grabs a list of all IPAddresses for this account
// https://www.twilio.com/docs/api/rest/ip-access-control-list#list-get-ipaddress
func (act Account) ListIPAddresses(alSid string) (IPAddressList, error) {
	var il IPAddressList
	if !validateMappingSid(alSid) {
		return il, errors.New("Invalid ip access control sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(ipAccessControlList.IPAddresses, act.AccountSid, alSid), act, &il)
	il.act = &act
	return il, err
}

// IPAddressUpdate contains properties to add or edit an IPAddress on your account
type IPAddressUpdate struct {
	FriendlyName string
	IPAddress    string
}

// GetReader is needed for the common.twilioPost interface
func (u IPAddressUpdate) GetReader() io.Reader {
	vals := url.Values{}
	if u.FriendlyName != "" {
		vals.Set("FriendlyName", u.FriendlyName)
	}
	if u.IPAddress != "" {
		vals.Set("IpAddress", u.IPAddress)
	}
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (u IPAddressUpdate) Validate() error {
	if u.FriendlyName == "" || u.IPAddress == "" {
		return errors.New("Must include both a friendly name and IP addressfor your new IP Address")
	}
	return nil
}

// AddIPAddress allows you to add a new ip address to your account
// https://www.twilio.com/docs/api/rest/ip-access-control-list#list-post-ipaddress
func (act Account) AddIPAddress(u IPAddressUpdate, alSid string) (IPAddressResource, error) {
	var i IPAddressResource
	if !validateMappingSid(alSid) {
		return i, errors.New("Invalid ip access control sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(ipAccessControlList.IPAddresses, act.AccountSid, alSid), u, act, &i)
	return i, err
}

// GetIPAddress grabs a single IP Address with the given sid
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-get-ipaddress
func (act Account) GetIPAddress(alSid, ipSid string) (IPAddressResource, error) {
	var i IPAddressResource
	if !validateMappingSid(alSid) {
		return i, errors.New("Invalid ip access control sid")
	}
	if !validateIPSid(ipSid) {
		return i, errors.New("Invalid ip sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(ipAccessControlList.IPAddress, act.AccountSid, alSid, ipSid), act, &i)
	return i, err
}

// UpdateIPAddress allows you to change the ip address and friendly name of an ipaddress resource
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-post-ipaddress
func (act Account) UpdateIPAddress(alSid, ipSid string, u IPAddressUpdate) (IPAddressResource, error) {
	var i IPAddressResource
	if !validateMappingSid(alSid) {
		return i, errors.New("Invalid ip access control sid")
	}
	if !validateIPSid(ipSid) {
		return i, errors.New("Invalid ip sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(ipAccessControlList.IPAddress, act.AccountSid, alSid, ipSid), u, act, &i)
	return i, err
}

// DeleteIPAddress allows you to delete an ip address from your account
// https://www.twilio.com/docs/api/rest/ip-access-control-list#instance-delete-ipaddress
func (act Account) DeleteIPAddress(alSid, ipSid string) error {
	if !validateMappingSid(alSid) {
		return errors.New("Invalid ip access control sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(ipAccessControlList.IPAddress, act.AccountSid, alSid, ipSid), act)
}
