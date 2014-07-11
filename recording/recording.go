package recording

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

type RecordingListFilter struct {
	CallSid     string
	DateCreated *common.JSONTime
}

func (f RecordingListFilter) GetQueryString() string {
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

func (act Account) getRecording(destURL string, resp *Recording) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp, 200)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Get(recSid string) (Recording, error) {
	var r Recording
	if !validateRecSid(recSid) {
		return r, errors.New("Invalid sid")
	}

	err := act.getRecording(fmt.Sprintf(getURL, act.AccountSid, string(recSid)), &r)
	return r, err
}

func (act Account) getRecordingList(destURL string, f RecordingListFilter, resp *RecordingList) error {
	return common.SendGetRequest(destURL+f.GetQueryString(), act, resp, 200)
}

func (act Account) List(f RecordingListFilter) (RecordingList, error) {
	var rl RecordingList
	err := act.getRecordingList(fmt.Sprintf(listURL, act.AccountSid), f, &rl)
	return rl, err
}

func (act Account) deleteRecording(destURL string) error {
	// send get request to twilio
	return common.SendDeleteRequest(destURL, act, 204)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Delete(recSid string) error {
	if !validateRecSid(recSid) {
		return errors.New("Invalid sid")
	}

	return act.deleteRecording(fmt.Sprintf(getURL, act.AccountSid, string(recSid)))
}
