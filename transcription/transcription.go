package transcription

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
)

func (act VoiceAccount) getTranscription(destUrl string, resp *Transcription) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp, 200)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act VoiceAccount) Transcription(trSid string) (Transcription, error) {
	var r Transcription
	if !validateTranscriptionSid(trSid) {
		return r, errors.New("Invalid sid")
	}

	err := act.getTranscription(fmt.Sprintf(transcriptionURL, act.AccountSid, string(trSid)), &r)
	return r, err
}

func (act VoiceAccount) getTranscriptionList(destUrl string, resp *TranscriptionList) error {
	return common.SendGetRequest(destUrl, act, resp, 200)
}

func (act VoiceAccount) TranscriptionList() (TranscriptionList, error) {
	var tl TranscriptionList
	err := act.getTranscriptionList(fmt.Sprintf(transcriptionListURL, act.AccountSid), &tl)
	return tl, err
}

func (act VoiceAccount) deleteTranscription(destUrl string) error {
	// send get request to twilio
	return common.SendDeleteRequest(destUrl, act, 204)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act VoiceAccount) DeleteTranscription(trSid string) error {
	if !validateTranscriptionSid(trSid) {
		return errors.New("Invalid sid")
	}

	return act.deleteTranscription(fmt.Sprintf(transcriptionURL, act.AccountSid, string(trSid)))
}
