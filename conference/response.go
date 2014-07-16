package conference

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

// kept private
type participants struct {
	Participants string `json:"participants"`
}

type Conference struct {
	APIVersion      string          `json:"api_version"`
	Sid             string          `json:"sid"`
	FriendlyName    string          `json:"friendly_name"`
	Status          string          `json:"status"`
	DateCreated     common.JSONTime `json:"date_created"`
	DateUpdated     common.JSONTime `json:"date_updated"`
	AccountSid      string          `json:"account_sid"`
	SubResourceURIs participants    `json:"subresource_uris"`
	URI             string          `json:"uri"`
}

type ConferenceList struct {
	common.ListResponseCore
	Conferences *[]Conference `json:"conferences"`
}

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

type ParticipantList struct {
	common.ListResponseCore
	Participants *[]Participant `json:"participants"`
}

func validateConferenceSid(sid string) bool {
	match, _ := regexp.MatchString(`^CF[0-9a-z]{32}$`, string(sid))
	return match
}

func validateCallSid(sid string) bool {
	match, _ := regexp.MatchString(`^CA[0-9a-z]{32}$`, string(sid))
	return match
}
