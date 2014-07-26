package conference

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	validateConferenceSid = regexp.MustCompile(`^CF[0-9a-z]{32}$`).MatchString
	validateCallSid       = regexp.MustCompile(`^CA[0-9a-z]{32}$`).MatchString
)

var participant = struct {
	Get, List string
}{
	List: "/2010-04-01/Accounts/%s/Conferences/%s/Participants.json",    // takes account sid, conference sid
	Get:  "/2010-04-01/Accounts/%s/Conferences/%s/Participants/%s.json", // takes account sid, conference sid, callsid
}

// Participant represents a user on a conference call.
//
// https://www.twilio.com/docs/api/rest/participant
type Participant struct {
	CallSid                string          `json:"call_sid"`
	ConferenceSid          string          `json:"conference_sid"`
	DateCreated            common.JSONTime `json:"date_created"`
	DateUpdated            common.JSONTime `json:"date_updated"`
	AccountSid             string          `json:"account_sid"`
	Muted                  bool            `json:"muted"`
	StartConferenceOnEnter bool            `json:"start_conference_on_enter"`
	EndConferenceOnExit    bool            `json:"end_conference_on_exit"`
	URI                    string          `json:"uri"`
}

// ParticipantUpdate is used to apply an update a participant in a call.
type ParticipantUpdate struct {
	Muted bool
}

// GetReader creates an io.Reader for the ParticipantUpdate resource.
func (p ParticipantUpdate) GetReader() io.Reader {
	vals := url.Values{}
	vals.Set("Muted", strconv.FormatBool(p.Muted))
	return strings.NewReader(vals.Encode())
}

// Validate helps Participant Update implement the twilioPost interface in common.
func (p ParticipantUpdate) Validate() error {
	return nil
}

func (p ParticipantUpdate) getParticipantQueryString() string {
	v := url.Values{}
	v.Set("Muted", strconv.FormatBool(p.Muted))
	encoded := v.Encode()
	if encoded != "" {
		encoded = "?" + encoded
	}
	return encoded
}

// GetParticipant returns the participant with callSid from conference with confSid
func (act Account) GetParticipant(confSid string, callSid string) (Participant, error) {
	var p Participant
	if !validateConferenceSid(confSid) {
		return p, errors.New("Invalid conference sid")
	} else if !validateCallSid(callSid) {
		return p, errors.New("Invalid call sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(participant.Get, act.AccountSid, confSid, callSid), act, &p)
	return p, err
}

// SetMute allows (un)muting of participant with callSid in conference with confSid
func (act Account) SetMute(confSid string, callSid string, a ParticipantUpdate) (Participant, error) {
	var p Participant
	if !validateConferenceSid(confSid) {
		return p, errors.New("Invalid conference sid")
	} else if !validateCallSid(callSid) {
		return p, errors.New("Invalid call sid")
	}
	err := common.SendPostRequest(fmt.Sprintf(participant.Get, act.AccountSid, confSid, callSid), a, act, &p)
	return p, err
}

// Kick a participant with callSid from conference with confSid.
func (act Account) Kick(confSid string, callSid string) error {
	if !validateConferenceSid(confSid) {
		return errors.New("Invalid conference sid")
	} else if !validateCallSid(callSid) {
		return errors.New("Invalid call sid for participant")
	}
	return common.SendDeleteRequest(fmt.Sprintf(participant.Get, act.AccountSid, confSid, callSid), act)
}

// ParticipantList lists all current participants in a call.
type ParticipantList struct {
	common.ListResponseCore
	act          *Account
	Participants *[]Participant `json:"participants"`
}

// ListParticipants queries for a list of participants in the conference with confSid.
func (act Account) ListParticipants(confSid string) (ParticipantList, error) {
	var pl ParticipantList
	if !validateConferenceSid(confSid) {
		return pl, errors.New("Invalid conference sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(participant.List, act.AccountSid, confSid), act, &pl)
	pl.act = &act
	return pl, err
}

// Next sets the ParticipantList to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (pl *ParticipantList) Next() error {
	if pl.Page == pl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(pl.NextPageURI, *pl.act, pl)
}
