package transcription

import (
	"encoding/json"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
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
	Uri               string           `json:"uri"`
}

type TranscriptionList struct {
	common.ListResponseCore
	Transcriptions *[]Transcription `json:"transcriptions"`
}

func validateTranscriptionSid(sid string) bool {
	match, _ := regexp.MatchString(`^TR[0-9a-z]{32}$`, string(sid))
	return match
}

func (r *Transcription) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading json from buffer => %s", err.Error())
	}

	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes))
	}
	return nil
}

func (r *TranscriptionList) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading json from buffer => %s", err.Error())
	}

	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes))
	}
	return nil
}
