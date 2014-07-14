package applications

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

const (
	getConnectAppURL  = "https://api.twilio.com/2010-04-01/Accounts/%s/AuthorizedConnectApps/%s.json" // takes an AccountSid & ConnectAppSid
	listConnectAppURL = "https://api.twilio.com/2010-04-01/Accounts/%s/AuthorizedConnectApps.json"    // takes an AccountSid
)

var validateConnectAppSid = regexp.MustCompile(`^CN[0-9a-z]{32}$`).MatchString

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

func (act Account) GetConnectApp(sid string) (ConnectApp, error) {
	var c ConnectApp
	if !validateConnectAppSid(sid) {
		return c, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, sid), act, &c)
	return c, err
}

type ConnectAppList struct {
	common.ListResponseCore
	AuthorizedConnectApps *[]ConnectApp
}

func (act Account) ListConnectApp() (ConnectAppList, error) {
	var cnl ConnectAppList
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid), act, &cnl)
	return cnl, err
}
