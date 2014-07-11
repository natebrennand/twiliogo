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

func (act Account) getTranscription(destUrl string, resp *Transcription) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Get(trSid string) (Transcription, error) {
	var r Transcription
	if !validateTranscriptionSid(trSid) {
		return r, errors.New("Invalid sid")
	}

	err := act.getTranscription(fmt.Sprintf(getURL, act.AccountSid, string(trSid)), &r)
	return r, err
}

func (act Account) getTranscriptionList(destUrl string, resp *TranscriptionList) error {
	return common.SendGetRequest(destUrl, act, resp)
}

func (act Account) List() (TranscriptionList, error) {
	var tl TranscriptionList
	err := act.getTranscriptionList(fmt.Sprintf(listURL, act.AccountSid), &tl)
	return tl, err
}

func (act Account) deleteTranscription(destUrl string) error {
	// send get request to twilio
	return common.SendDeleteRequest(destUrl, act)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Delete(trSid string) error {
	if !validateTranscriptionSid(trSid) {
		return errors.New("Invalid sid")
	}

	return act.deleteTranscription(fmt.Sprintf(getURL, act.AccountSid, string(trSid)))
}
