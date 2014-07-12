package conference

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"strconv"
	"strings"
)

type ParticipantAttr struct {
	Muted *bool
}

func (p ParticipantAttr) GetReader() io.Reader {
	vals := url.Values{}
	if p.Muted != nil {
		vals.Set("Muted", strconv.FormatBool(*p.Muted))
	}
	return strings.NewReader(vals.Encode())
}

func (p ParticipantAttr) Validate() error {
	if p.Muted == nil {
		return errors.New("Muted must be set, else nothing to modify")
	}
	return nil
}

func (f ParticipantAttr) GetParticipantQueryString() string {
	v := url.Values{}
	if f.Muted != nil {
		v.Set("Muted", strconv.FormatBool(*f.Muted))
	}
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

// Get a participant with callSid from conference with confSid
func (act Account) Participant(confSid string, callSid string) (Participant, error) {
	var p Participant
	if !validateConferenceSid(confSid) {
		return p, errors.New("Invalid conference sid")
	} else if !validateCallSid(callSid) {
		return p, errors.New("Invalid call sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(participantURL, act.AccountSid, confSid, callSid), act, &p)
	fmt.Println("p: ", p.CallSid)
	return p, err
}

// Mute or unmute participant with callSid in conference with confSid
func (act Account) SetMute(confSid string, callSid string, a ParticipantAttr) (Participant, error) {
	var p Participant
	if !validateConferenceSid(confSid) {
		return p, errors.New("Invalid conference sid")
	} else if !validateCallSid(callSid) {
		return p, errors.New("Invalid call sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(participantURL, act.AccountSid, confSid, callSid), a, act, &p)
	return p, err
}

// Kicks participant with callSid from conference with confSid
func (act Account) Kick(confSid string, callSid string) error {
	if !validateConferenceSid(confSid) {
		return errors.New("Invalid conference sid")
	} else if !validateCallSid(callSid) {
		return errors.New("Invalid call sid for participant")
	}
	return common.SendDeleteRequest(fmt.Sprintf(participantURL, act.AccountSid, confSid, callSid), act)
}

// Get list of participants in conference with confSid
func (act Account) Participants(f ParticipantAttr, confSid string) (ParticipantList, error) {
	var pl ParticipantList
	if !validateConferenceSid(confSid) {
		return pl, errors.New("Invalid conference sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(participantsURL, act.AccountSid, confSid)+f.GetParticipantQueryString(), act, &pl)
	return pl, err
}
