package conference

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"net/url"
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

type ListFilter struct {
	Status       string
	FriendlyName string
	DateCreated  *common.JSONTime
	DateUpdated  *common.JSONTime
}

func (f ListFilter) GetQueryString() string {
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

func (act Account) getConference(destURL string, resp *Conference) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp)
}

// Get a info about a conference with confSid
func (act Account) Get(confSid string) (Conference, error) {
	var c Conference
	if !validateConferenceSid(confSid) {
		return c, errors.New("Invalid sid")
	}

	err := act.getConference(fmt.Sprintf(getURL, act.AccountSid, string(confSid)), &c)
	return c, err
}

func (act Account) getConferenceList(destURL string, f ListFilter, resp *ConferenceList) error {
	return common.SendGetRequest(destURL+f.GetQueryString(), act, resp)
}

// Get list of conferences for this account
func (act Account) List(f ListFilter) (ConferenceList, error) {
	var cl ConferenceList
	err := act.getConferenceList(fmt.Sprintf(listURL, act.AccountSid), f, &cl)
	return cl, err
}
