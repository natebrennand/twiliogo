package conference

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
)

var conference = struct {
	Get, List string
}{
	Get:  "/2010-04-01/Accounts/%s/Conferences/%s.json", // takes account sid, conference sid
	List: "/2010-04-01/Accounts/%s/Conferences.json",    // takes account sid
}

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

// kept private
type participants struct {
	Participants string `json:"participants"`
}

// Conference represents a conference call that can be updated using the
// participants resource.
//
// https://www.twilio.com/docs/api/rest/conference
type Conference struct {
	APIVersion      string          `json:"api_version"`
	Sid             string          `json:"sid"`
	FriendlyName    string          `json:"friendly_name"`
	Status          string          `json:"status"`
	DateCreated     common.JSONTime `json:"date_created"`
	DateUpdated     common.JSONTime `json:"date_updated"`
	AccountSid      string          `json:"account_sid"`
	SubResourceURIs participants    `json:"subresource_uris"`
	URI             string          `json:"uri"`
}

func (f ListFilter) getQueryString() string {
	v := url.Values{}
	if f.Status != "" {
		v.Set("Status", f.Status)
	}
	if f.FriendlyName != "" {
		v.Set("FriendlyName", f.FriendlyName)
	}
	if f.DateCreated != nil {
		v.Set("DateCreated", f.DateCreated.Format(common.GMTTimeLayout))
	}
	if f.DateUpdated != nil {
		v.Set("DateUpdated", f.DateUpdated.Format(common.GMTTimeLayout))
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

// Get info about a conference identified by a ConferenceSID
func (act Account) Get(confSid string) (Conference, error) {
	var c Conference
	if !validateConferenceSid(confSid) {
		return c, errors.New("Invalid sid")
	}

	err := common.SendGetRequest(fmt.Sprintf(conference.Get, act.AccountSid, confSid), act, &c)
	return c, err
}

// List contains a list of conference resources.
type List struct {
	common.ListResponseCore
	Conferences *[]Conference `json:"conferences"`
}

// ListFilter allows filtering of conference resource lists
type ListFilter struct {
	Status       string
	FriendlyName string
	DateCreated  *common.JSONTime
	DateUpdated  *common.JSONTime
}

// List the conferences for this account
func (act Account) List(f ListFilter) (List, error) {
	var cl List
	err := common.SendGetRequest(fmt.Sprintf(conference.List, act.AccountSid)+f.getQueryString(), act, &cl)
	return cl, err
}
