package sms

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

var validateSmsSid = regexp.MustCompile(`^(SM|MM)[0-9a-z]{32}$`).MatchString

type Message struct {
	common.ResponseCore
	Body        string           `json:"body"`
	DateSent    common.JSONTime  `json:"date_sent"`
	NumSegments int              `json:"num_segments,string"`
	NumMedia    int              `json:"num_media,string"`
	Price       common.JSONFloat `json:"price"`
}

type MessageList struct {
	common.ListResponseCore
	Messages *[]Message `json:"messages"`
}
