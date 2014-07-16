package transcription

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
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

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Get(trSid string) (Transcription, error) {
	var r Transcription
	if !validateTranscriptionSid(trSid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(getURL, act.AccountSid, string(trSid)), act, &r)
	return r, err
}

func (act Account) List() (TranscriptionList, error) {
	var tl TranscriptionList
	// err := act.getTranscriptionList(fmt.Sprintf(listURL, act.AccountSid), &tl)
	err := common.SendGetRequest(fmt.Sprintf(listURL, act.AccountSid), act, &tl)
	return tl, err
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Delete(trSid string) error {
	if !validateTranscriptionSid(trSid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(getURL, act.AccountSid, string(trSid)), act)
}
