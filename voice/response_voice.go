package voice

import (
	"encoding/json"
	"github.com/natebrennand/twiliogo/common"
)

type Response struct {
	common.ResponseCore
	Price          common.JsonPrice `json:"price"`
	DateCreated    common.JsonTime  `json:"date_created"`
	DateUpdated    common.JsonTime  `json:"date_updated"`
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      common.JsonTime `json:"start_time"`
	EndTime        common.JsonTime `json:"end_time"`
	Duration       float64         `json:"duration"`
	AnsweredBy     string          `json:"answered_by"`
	ForwardedFrom  string          `json:"fowarded_from"`
	CallerName     string          `json:"caller_name"`
}

// Unmarshals a twilio sms response into a Response struct.
func Unmarshal(data []byte, msg *Response) error {
	return json.Unmarshal(data, msg)
}
