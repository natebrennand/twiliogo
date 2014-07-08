package sms

import (
	"github.com/natebrennand/twiliogo/common"
)

type Response struct {
	common.ResponseCore
	NumSegments int              `json:"num_segments,string"`
	NumMedia    int              `json:"num_media,string"`
	Price       common.JsonPrice `json:"price"`
	DateCreated common.JsonTime  `json:"date_created"`
	DateSent    common.JsonTime  `json:"date_sent"`
	DateUpdated common.JsonTime  `json:"date_updated"`
}
