package applications

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

var connectApps = struct {
	Get, List string
}{
	Get:  "/2010-04-01/Accounts/%s/AuthorizedConnectApps/%s.json", // takes an AccountSid & ConnectAppSid
	List: "/2010-04-01/Accounts/%s/AuthorizedConnectApps.json",    // takes an AccountSid
}

var validateConnectAppSid = regexp.MustCompile(`^CN[0-9a-z]{32}$`).MatchString

// ConnectApp is a subaccount application that has specific permissions.
//
// https://www.twilio.com/docs/api/rest/authorized-connect-apps
type ConnectApp struct {
	DateCreated            common.JSONTime `json:"date_created"`
	DateUpdated            common.JSONTime `json:"date_updated"`
	AccountSid             string          `json:"account_sid"`
	ConnectAppSid          string          `json:"connect_app_sid"`
	ConnectAppFriendlyName string          `json:"connect_app_friendly_name"`
	ConnectAppDescription  string          `json:"connect_app_description"`
	ConnectAppCompanyName  string          `json:"connect_app_company_name"`
	ConnectAppHomepageURL  string          `json:"connect_app_homepage_url"`
	URI                    string          `json:"uri"`
	Permissions            []string        `json:"permissions"`
}

// GetConnectApp returns the resource for a specific ConnectApp.
func (act Account) GetConnectApp(sid string) (ConnectApp, error) {
	var c ConnectApp
	if !validateConnectAppSid(sid) {
		return c, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(connectApps.Get, act.AccountSid, sid), act, &c)
	return c, err
}

// ConnectAppList is a list of ConnectApps
type ConnectAppList struct {
	common.ListResponseCore
	AuthorizedConnectApps *[]ConnectApp
	act                   *Account
}

// ListConnectApp returns the list of ConnectApps associated with the account.
func (act Account) ListConnectApp() (ConnectAppList, error) {
	var cal ConnectAppList
	err := common.SendGetRequest(fmt.Sprintf(connectApps.List, act.AccountSid), act, &cal)
	cal.act = &act
	return cal, err
}

// Next sets the ConnectAppList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (cal *ConnectAppList) next() error {
	if cal.Page == cal.NumPages-1 {
		return errors.New("no more new pages")
	}
	return common.SendGetRequest(cal.NextPageURI, *cal.act, cal)
}
