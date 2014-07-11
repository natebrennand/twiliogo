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

// Get a info about a conference with confSid
func (act Account) Get(confSid string) (Conference, error) {
	var c Conference
	if !validateConferenceSid(confSid) {
		return c, errors.New("Invalid sid")
	}

	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, confSid), act, &c)
	return c, err
}

// Get list of conferences for this account
func (act Account) List(f ListFilter) (ConferenceList, error) {
	var cl ConferenceList
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid)+f.GetQueryString(), act, &cl)
	return cl, err
}
