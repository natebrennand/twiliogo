package voice

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
)

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

func (act VoiceAccount) getRecording(destURL string, resp *Recording) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp, 200)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act VoiceAccount) Recording(recSid string) (Recording, error) {
	var r Recording
	if !validateRecSid(recSid) {
		return r, errors.New("Invalid sid")
	}

	err := act.getRecording(fmt.Sprintf(recordingURL, act.AccountSid, string(recSid)), &r)
	return r, err
}

func (act VoiceAccount) getRecordingList(destURL string, f RecordingListFilter, resp *RecordingList) error {
	return common.SendGetRequest(destURL+f.GetQueryString(), act, resp, 200)
}

func (act VoiceAccount) RecordingList(f RecordingListFilter) (RecordingList, error) {
	var rl RecordingList
	err := act.getRecordingList(fmt.Sprintf(recordingListURL, act.AccountSid), f, &rl)
	return rl, err
}

func (act VoiceAccount) deleteRecording(destURL string) error {
	// send get request to twilio
	return common.SendDeleteRequest(destURL, act, 204)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act VoiceAccount) Delete(recSid string) error {
	if !validateRecSid(recSid) {
		return errors.New("Invalid sid")
	}

	return act.deleteRecording(fmt.Sprintf(recordingURL, act.AccountSid, string(recSid)))
}
