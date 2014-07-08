package voice

import (
	"github.com/natebrennand/twilio-go/common"
	"time"
)

type VoiceResponseJson struct {
	common.ResponseCore
	JsonPrice          float64   `json:"price"`
	JsonDateCreated    string    `json:"date_created"`
	JsonDateUpdated    string    `json:"date_updated"`
	JsonParentCallSid  string    `json:"parent_call_sid"`
	JsonPhoneNumberSid string    `json:"phone_number_sid"`
	JsonStartTime      time.Time `json:"start_time"`
	JsonEndTime        time.Time `json:"end_time"`
	JsonDuration       float64   `json:"duration"`
	JsonAnsweredBy     string    `json:"answered_by"`
	JsonForwardedFrom  string    `json:"forwarded_from"`
	JsonCallerName     string    `json:"caller_name"`
}

type Response struct {
	common.ResponseCore
	Price          float64
	DateCreated    time.Time
	DateUpdated    time.Time
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      time.Time
	EndTime        time.Time
	Duration       float64
	AnsweredBy     string
	ForwardedFrom  string
	CallerName     string
}
