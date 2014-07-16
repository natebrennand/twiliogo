package recording

import (
	"github.com/natebrennand/twiliogo/common"
)

type Recording struct {
	Sid         string          `json:"sid"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	AccountSid  string          `json:"account_sid"`
	CallSid     string          `json:"call_sid"`
	Duration    *int64          `json:"duration,string"`
	APIVersion  string          `json:"api_version"`
	URI         string          `json:"uri"`
}

type RecordingList struct {
	common.ListResponseCore
	Recordings *[]Recording `json:"recordings"`
}
