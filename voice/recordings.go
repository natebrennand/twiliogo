package voice

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/url"
)

type RecordingListFilter struct {
	CallSid     string
	DateCreated *common.JsonTime
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

func (act VoiceAccount) getRecording(destUrl string, resp *Recording) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp, 200)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act VoiceAccount) Recording(recSid string) (Recording, error) {
	var r Recording
	if !validateRecSid(recSid) {
		return r, errors.New("Invalid sid")
	}

	err := act.getRecording(fmt.Sprintf(recordingUrl, act.AccountSid, string(recSid)), &r)
	return r, err
}

func (act VoiceAccount) getRecordingList(destUrl string, f RecordingListFilter, resp *RecordingList) error {
	return common.SendGetRequest(destUrl+f.GetQueryString(), act, resp, 200)
}

func (act VoiceAccount) RecordingList(f RecordingListFilter) (RecordingList, error) {
	var rl RecordingList
	err := act.getRecordingList(fmt.Sprintf(recordingListUrl, act.AccountSid), f, &rl)
	return rl, err
}

func (act VoiceAccount) deleteRecording(destUrl string) error {
	// send get request to twilio
	return common.SendDeleteRequest(destUrl, act, 204)
}

// Returns data about recording as json
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act VoiceAccount) Delete(recSid string) error {
	if !validateRecSid(recSid) {
		return errors.New("Invalid sid")
	}

	return act.deleteRecording(fmt.Sprintf(recordingUrl, act.AccountSid, string(recSid)))
}
