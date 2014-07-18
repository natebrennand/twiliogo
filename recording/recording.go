package recording

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
	"regexp"
)

// holds url values used in queries
var recordings = struct {
	Get, List string
}{
	Get:  "/2010-04-01/Accounts/%s/Recordings/%s.json",
	List: "/2010-04-01/Accounts/%s/Recordings.json",
}

var validateRecordingSid = regexp.MustCompile(`^RE[0-9a-z]{32}$`).MatchString

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

// Recording represents a voice recording from a phone call.
type Recording struct {
	Sid         string          `json:"sid"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	AccountSid  string          `json:"account_sid"`
	CallSid     string          `json:"call_sid"`
	Duration    int64           `json:"duration,string"`
	APIVersion  string          `json:"api_version"`
	URI         string          `json:"uri"`
}

// Get returns data about recording as json.
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Get(recSid string) (Recording, error) {
	var r Recording
	if !validateRecordingSid(recSid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(recordings.Get, act.AccountSid, recSid), act, &r)
	return r, err
}

// List holds a list of Reording references.
type List struct {
	common.ListResponseCore
	Recordings *[]Recording `json:"recordings"`
}

// ListFilter can be used to filter the results of a list recordings query.
type ListFilter struct {
	CallSid     string
	DateCreated *common.JSONTime
}

func (f ListFilter) getQueryString() string {
	v := url.Values{}
	if f.CallSid != "" {
		v.Set("CallSid", f.CallSid)
	}
	if f.DateCreated != nil {
		v.Set("DateCreated", f.DateCreated.Format(common.GMTTimeLayout))
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

// List returns a list of recordings that fulfill the provided filter.
func (act Account) List(f ListFilter) (List, error) {
	var rl List
	err := common.SendGetRequest(fmt.Sprintf(recordings.List, act.AccountSid)+f.getQueryString(), act, &rl)
	return rl, err
}

// Delete removes a recording from Twilio servers.
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Delete(recSid string) error {
	if !validateRecordingSid(recSid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(recordings.Get, act.AccountSid, recSid), act)
}
