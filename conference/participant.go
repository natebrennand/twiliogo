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

func (act Account) getParticipant(destURL string, resp *Participant) error {
	return common.SendGetRequest(destURL, act, resp)
}

// Get a participant with callSid from conference with confSid
func (act Account) Participant(confSid string, callSid string) (Participant, error) {
	var p Participant
	if !validateConferenceSid(confSid) {
		return p, errors.New("Invalid conference sid")
	}
	if !validateCallSid(callSid) {
		return p, errors.New("Invalid call sid")
	}

	err := act.getParticipant(fmt.Sprintf(getURL, act.AccountSid, string(confSid), string(callSid)), &p)
	return p, err
}

func (act Account) setParticipantMute(dest string, msg ParticipantAttr, resp *Participant) error {
	return common.SendPostRequest(dest, msg, act, resp)
}

// Mute or unmute participant with callSid in conference with confSid
func (act Account) SetMute(confSid string, callSid string, a ParticipantAttr) (Participant, error) {
	var p Participant
	err := act.setParticipantMute(fmt.Sprintf(modifyURL, act.AccountSid, string(confSid), string(callSid)), a, &p)
	return p, err
}

func (act Account) kickParticipant(destURL string) error {
	return common.SendDeleteRequest(destURL, act)
}

// Kicks participant with callSid from conference with confSid
func (act Account) Kick(confSid string, callSid string) error {
	if !validateConferenceSid(confSid) {
		return errors.New("Invalid conference sid")
	}
	if !validateCallSid(callSid) {
		return errors.New("Invalid call sid for participant")
	}
	return act.kickParticipant(fmt.Sprintf(modifyURL, act.AccountSid, string(confSid), string(callSid)))
}

func (act Account) getParticipantList(destURL string, f ParticipantAttr, resp *ParticipantList) error {
	return common.SendGetRequest(destURL+f.GetParticipantQueryString(), act, resp)
}

// Get list of participants in conference with confSid
func (act Account) Participants(f ParticipantAttr, confSid string) (ParticipantList, error) {
	var pl ParticipantList
	err := act.getParticipantList(fmt.Sprintf(participantsURL, act.AccountSid, string(confSid)), f, &pl)
	return pl, err
}
