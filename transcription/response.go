package transcription

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

type Transcription struct {
	Sid               string           `json:"sid"`
	DateCreated       common.JSONTime  `json:"date_created"`
	DateUpdated       common.JSONTime  `json:"date_updated"`
	AccountSid        string           `json:"account_sid"`
	Status            string           `json:"status"`
	RecordingSid      string           `json:"recording_sid"`
	Duration          string           `json:"duration"`
	TranscriptionText string           `json:"transcription_text"`
	Price             common.JSONPrice `json:"price"`
	PriceUnit         string           `json:"price_unit"`
	URI               string           `json:"uri"`
}

type TranscriptionList struct {
	common.ListResponseCore
	Transcriptions *[]Transcription `json:"transcriptions"`
}

func validateTranscriptionSid(sid string) bool {
	match, _ := regexp.MatchString(`^TR[0-9a-z]{32}$`, string(sid))
	return match
}
