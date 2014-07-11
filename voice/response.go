package voice

import (
	"encoding/json"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Call struct {
	common.ResponseCore
	Price          common.JSONPrice `json:"price"`
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      common.JSONTime `json:"start_time"`
	EndTime        common.JSONTime `json:"end_time"`
	Duration       string          `json:"duration"`
	AnsweredBy     string          `json:"answered_by"`
	ForwardedFrom  string          `json:"fowarded_from"`
	CallerName     string          `json:"caller_name"`
}

type Recording struct {
	Sid         string          `json:"sid"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	AccountSid  string          `json:"account_sid"`
	CallSid     string          `json:"call_sid"`
	Duration    string          `json:"duration"`
	APIVersion  string          `json:"api_version"`
	URI         string          `json:"uri"`
}

type Transcription struct {
	Sid               string           `json:"sid"`
	DateCreated       common.JsonTime  `json:"date_created"`
	DateUpdated       common.JsonTime  `json:"date_updated"`
	AccountSid        string           `json:"account_sid"`
	Status            string           `json:"status"`
	RecordingSid      string           `json:"recording_sid"`
	Duration          string           `json:"duration"`
	TranscriptionText string           `json:"transcription_text"`
	Price             common.JsonPrice `json:"price"`
	PriceUnit         string           `json:"price_unit"`
	Uri               string           `json:"uri"`
}

type TranscriptionList struct {
	common.ListResponseCore
	Transcriptions *[]Transcription `json:"transcriptions"`
}

type RecordingList struct {
	common.ListResponseCore
	Recordings *[]Recording `json:"recordings"`
}

type CallList struct {
	common.ListResponseCore
	Calls *[]Call `json:"calls"`
}

func validateCallSid(sid string) bool {
	match, _ := regexp.MatchString(`^CA[0-9a-z]{32}$`, string(sid))
	return match
}

func validateRecSid(sid string) bool {
	match, _ := regexp.MatchString(`^RE[0-9a-z]{32}$`, string(sid))
	return match
}

func validateTranscriptionSid(sid string) bool {
	match, _ := regexp.MatchString(`^TR[0-9a-z]{32}$`, string(sid))
	return match
}

func (r *Call) Build(resp *http.Response) error {
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

func (r *Recording) Build(resp *http.Response) error {
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

func (r *RecordingList) Build(resp *http.Response) error {
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

func (l *CallList) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading json from buffer => %s", err.Error())
	}
	err = json.Unmarshal(bodyBytes, l)
	if err != nil {
		return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes))
	}
	return nil
}
