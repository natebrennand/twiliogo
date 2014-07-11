package transcription

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
)

type TranscriptionAccount struct {
	AccountSid string
	Token      string
	Client     http.Client
}

func (act TranscriptionAccount) GetSid() string {
	return act.AccountSid
}
func (act TranscriptionAccount) GetToken() string {
	return act.Token
}
func (act TranscriptionAccount) GetClient() http.Client {
	return act.Client
}

func (act TranscriptionAccount) getTranscription(destUrl string, resp *Transcription) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp, 200)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act TranscriptionAccount) Get(trSid string) (Transcription, error) {
	var r Transcription
	if !validateTranscriptionSid(trSid) {
		return r, errors.New("Invalid sid")
	}

	err := act.getTranscription(fmt.Sprintf(transcriptionURL, act.AccountSid, string(trSid)), &r)
	return r, err
}

func (act TranscriptionAccount) getTranscriptionList(destUrl string, resp *TranscriptionList) error {
	return common.SendGetRequest(destUrl, act, resp, 200)
}

func (act TranscriptionAccount) List() (TranscriptionList, error) {
	var tl TranscriptionList
	err := act.getTranscriptionList(fmt.Sprintf(transcriptionListURL, act.AccountSid), &tl)
	return tl, err
}

func (act TranscriptionAccount) deleteTranscription(destUrl string) error {
	// send get request to twilio
	return common.SendDeleteRequest(destUrl, act, 204)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act TranscriptionAccount) Delete(trSid string) error {
	if !validateTranscriptionSid(trSid) {
		return errors.New("Invalid sid")
	}

	return act.deleteTranscription(fmt.Sprintf(transcriptionURL, act.AccountSid, string(trSid)))
}
