package transcription

// https://www.twilio.com/docs/api/rest/transcription

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

// holds url values used in queries
var transcriptions = struct {
	Get, List string
}{
	Get:  "/2010-04-01/Accounts/%s/Transcriptions/%s.json",
	List: "/2010-04-01/Accounts/%s/Transcriptions.json",
}

var validateTranscriptionSid = regexp.MustCompile(`^TR[0-9a-z]{32}$`).MatchString

// Account wraps the common Account struct to embed the AccountSid & Token.
type Account struct {
	common.Account
}

// Transcription contains data from a recording transcription.
type Transcription struct {
	common.ResponseCore2
	Status            string           `json:"status"`
	RecordingSid      string           `json:"recording_sid"`
	Duration          int64            `json:"duration,string"`
	TranscriptionText string           `json:"transcription_text"`
	Price             common.JSONFloat `json:"price"`
	PriceUnit         string           `json:"price_unit"`
}

// Get returns data about a transcription as a json.
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Get(trSid string) (Transcription, error) {
	var r Transcription
	if !validateTranscriptionSid(trSid) {
		return r, errors.New("Invalid sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(transcriptions.Get, act.AccountSid, string(trSid)), act, &r)
	return r, err
}

// List contains a list of transcription resources.
type List struct {
	common.ListResponseCore
	Transcriptions *[]Transcription `json:"transcriptions"`
	act            *Account
}

// List returns a list of all transcriptions associated with the account.
func (act Account) List() (List, error) {
	var tl List
	// err := act.getTranscriptionList(fmt.Sprintf(transcriptions.List, act.AccountSid), &tl)
	err := common.SendGetRequest(fmt.Sprintf(transcriptions.List, act.AccountSid), act, &tl)
	tl.act = &act
	return tl, err
}

// Next sets the List to the next page of the list resource, returns an error in the
// case that there are no more pages left.
func (tl *List) Next() error {
	if tl.Page == tl.NumPages-1 {
		return errors.New("No more new pages")
	}
	return common.SendGetRequest(tl.NextPageURI, *tl.act, tl)
}

// Delete removes a transcription resource from Twilio's server.
// Can get .mp3 or .wav of recording from the uri provided in Recording
func (act Account) Delete(trSid string) error {
	if !validateTranscriptionSid(trSid) {
		return errors.New("Invalid sid")
	}
	return common.SendDeleteRequest(fmt.Sprintf(transcriptions.Get, act.AccountSid, string(trSid)), act)
}
