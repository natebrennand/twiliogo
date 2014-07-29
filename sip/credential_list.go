package sip

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"strings"
)

var credentialList = struct {
	List, Get, Post string
}{
	List: "/2010-04-01/Accounts/%s/SIP/CredentialLists.json",    // takes an AccountSid
	Get:  "/2010-04-01/Accounts/%s/SIP/CredentialLists/%s.json", // takes an AccountSid & CredentialListSid
	Post: "/2010-04-01/Accounts/%s/SIP/CredentialLists/%s.json", // takes an AccountSid & CredentialListSid
}

var credential = struct {
	List, Get, Post string
}{
	List: "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists/%s/Credentials.json",    // takes an AccountSid & CredentialListSid
	Get:  "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists/%s/Credentials/%s.json", // takes an AccountSid & CredentialListSid & CredentialSid
	Post: "/2010-04-01/Accounts/%s/SIP/IpAccessControlLists/%s/Credentials/%s.json", // takes an AccountSid & CredentialListSid & CredentialSid
}

// CredentialListResource represents an CredentialList
// https://www.twilio.com/docs/api/rest/credential-list#instance-properties
type CredentialListResource struct {
	common.ResourceInfo
	SubresourceURIs credentialSubresourceURI `json:"subresource_uris"`
	FriendlyName    string                   `json:"friendly_name"`
}

// CredentialLists represents a list of CredentialList
type CredentialLists struct {
	common.ListResponseCore
	Lists *[]CredentialListResource `json:"credential_lists"`
	act   *Account
}

// Next sets the CredentialLists to the next page of the list resource, returns an error in the
// case that there are no more pages left
func (cl *CredentialLists) Next() error {
	if cl.Page == cl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(cl.NextPageURI, *cl.act, cl)
}

// UpdateCredentialList contains properties to add or edit an CredentialList on your account
type UpdateCredentialList struct {
	FriendlyName string
}

// GetReader is needed for the common.twilioPost interface
func (u UpdateCredentialList) GetReader() io.Reader {
	vals := url.Values{}
	vals.Set("FriendlyName", u.FriendlyName)
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (u UpdateCredentialList) Validate() error {
	if u.FriendlyName == "" {
		return errors.New("Must include at least a friendly name for your new credential list")
	}
	return nil
}

// ListCredentialLists grabs a list of all CredentialList for this account
// https://www.twilio.com/docs/api/rest/credential-list#list-get
func (act Account) ListCredentialLists() (CredentialLists, error) {
	var cl CredentialLists
	err := common.SendGetRequest(fmt.Sprintf(credentialList.List, act.AccountSid), act, &cl)
	cl.act = &act
	return cl, err
}

// AddCredentialList allows you to add a new CredentialList to your account
// https://www.twilio.com/docs/api/rest/credential-list#list-post
func (act Account) AddCredentialList(u UpdateCredentialList) (CredentialListResource, error) {
	var c CredentialListResource
	err := common.SendPostRequest(fmt.Sprintf(credentialList.List, act.AccountSid), u, act, &c)
	return c, err
}

// GetCredentialList allows you to get an CredentialList
// https://www.twilio.com/docs/api/rest/credential-list#instance-get
func (act Account) GetCredentialList(clSid string) (CredentialListResource, error) {
	var c CredentialListResource
	if !validateCredentialSid(clSid) {
		return c, errors.New("Invalid credential list sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(credentialList.Get, act.AccountSid, clSid), act, &c)
	return c, err
}

// UpdateCredentialList allows you to update an CredentialList with a new FriendlyName
// https://www.twilio.com/docs/api/rest/credential-list#instance-post
func (act Account) UpdateCredentialList(u UpdateCredentialList, clSid string) (CredentialListResource, error) {
	var c CredentialListResource
	if !validateCredentialListSid(clSid) {
		return c, errors.New("Invalid ip access control sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(credentialList.Post, act.AccountSid, clSid), u, act, &c)
	return c, err
}

// DeleteCredentialList allows you to delete a CredentialList from your account
// https://www.twilio.com/docs/api/rest/credential-list#instance-delete
func (act Account) DeleteCredentialList(clSid string) error {
	if !validateCredentialListSid(clSid) {
		return errors.New("Invalid credential list sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(credentialList.Get, act.AccountSid, clSid), act)
}

// CredentialResource represents an Credential
// https://www.twilio.com/docs/api/rest/credential-list#list-credential
type CredentialResource struct {
	common.ResourceInfo
	Username string `json:"username"`
}

// CredentialsResource represents a list of Credentials
type CredentialsResource struct {
	common.ListResponseCore
	Credentials *[]CredentialResource `json:"credentials"`
	act         *Account
}

// ListCredentials grabs a list of all Credentials for this account
// https://www.twilio.com/docs/api/rest/credential-list#list-get-credential
func (act Account) ListCredentials(clSid string) (CredentialsResource, error) {
	var cl CredentialsResource
	if !validateCredentialSid(clSid) {
		return cl, errors.New("Invalid credential list sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(credential.List, act.AccountSid, clSid), act, &cl)
	cl.act = &act
	return cl, err
}

// CredentialUpdate contains properties to add or edit an Credential on your account
type CredentialUpdate struct {
	Username string
	Password string
}

// GetReader is needed for the common.twilioPost interface
func (u CredentialUpdate) GetReader() io.Reader {
	vals := url.Values{}
	vals.Set("Username", u.Username)
	vals.Set("Password", u.Password)
	return strings.NewReader(vals.Encode())
}

// Validate is needed for the common.twilioPost interface
func (u CredentialUpdate) Validate() error {
	if u.Password == "" || u.Username == "" {
		return errors.New("Must include both a username and password your credential")
	} else if len(u.Password) < 12 || u.Password == strings.ToUpper(u.Password) ||
		u.Password == strings.ToLower(u.Password) ||
		strings.ContainsAny(u.Password, "01234567890") {
		return errors.New("Password must be longer than 12 characters, have at least one mixed case, and at least one digit")
	}
	return nil
}

// AddCredential allows you to add a new credential to your account
// https://www.twilio.com/docs/api/rest/credential-list#list-post-credential
func (act Account) AddCredential(u CredentialUpdate, clSid string) (CredentialResource, error) {
	var c CredentialResource
	if !validateCredentialListSid(clSid) {
		return c, errors.New("Invalid credential list sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(credential.List, act.AccountSid, clSid), u, act, &c)
	return c, err
}

// GetCredential grabs a single IP Address with the given sid
// https://www.twilio.com/docs/api/rest/credential-list#instance-get-credential
func (act Account) GetCredential(clSid, credSid string) (CredentialResource, error) {
	var c CredentialResource
	if !validateCredentialListSid(clSid) {
		return c, errors.New("Invalid credential list sid")
	} else if !validateCredentialSid(credSid) {
		return c, errors.New("Invalid credential sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(credential.Get, act.AccountSid, clSid, credSid), act, &c)
	return c, err
}

// UpdateCredential allows you to change the password of a credential resource
// https://www.twilio.com/docs/api/rest/credential-list#instance-post-credential
func (act Account) UpdateCredential(clSid, credSid string, u CredentialUpdate) (CredentialResource, error) {
	var c CredentialResource
	if !validateCredentialListSid(clSid) {
		return c, errors.New("Invalid credential list sid")
	} else if !validateCredentialSid(credSid) {
		return c, errors.New("Invalid credential sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(credential.Post, act.AccountSid, clSid, credSid), u, act, &c)
	return c, err
}

// DeleteCredential allows you to delete a credential from your account
// https://www.twilio.com/docs/api/rest/credential-list#instance-delete-credential
func (act Account) DeleteCredential(clSid, credSid string) error {
	if !validateMappingSid(clSid) {
		return errors.New("Invalid credential list sid")
	} else if !validateCredentialSid(credSid) {
		return errors.New("Invalid credential sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(credential.Get, act.AccountSid, clSid, credSid), act)
}
