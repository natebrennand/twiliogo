package recording

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"net/url"
	"regexp"
)

var validateRecordingSid = regexp.MustCompile(`^RE[0-9a-z]{32}$`).MatchString

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
	CallSid     string
	DateCreated *common.JSONTime
}

func (f ListFilter) GetQueryString() string {
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

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Get(recSid string) (Recording, error) {
	var r Recording
	if !validateRecordingSid(recSid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, recSid), act, &r)
	return r, err
}

func (act Account) List(f ListFilter) (RecordingList, error) {
	var rl RecordingList
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid)+f.GetQueryString(), act, &rl)
	return rl, err
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Delete(recSid string) error {
	if !validateRecordingSid(recSid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(getURL, act.AccountSid, recSid), act)
}
