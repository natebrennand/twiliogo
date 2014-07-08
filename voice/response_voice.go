package voice

import (
	"encoding/json"
	"github.com/natebrennand/twiliogo/common"
	"time"
)

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

// Unmarshals a twilio sms response into a Response struct.
func Unmarshal(data []byte, msg *Response) error {
	return json.Unmarshal(data, msg)
}
