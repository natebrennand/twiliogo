package sms

import (
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

func validateSmsSid(sid string) bool {
	match, _ := regexp.MatchString(`^(SM|MM)[0-9a-z]{32}$`, string(sid))
	return match
}

type Message struct {
	common.ResponseCore
	Body        string           `json:"body"`
	DateSent    common.JSONTime  `json:"date_sent"`
	NumSegments int              `json:"num_segments,string"`
	NumMedia    int              `json:"num_media,string"`
	Price       common.JSONPrice `json:"price"`
}

type MessageList struct {
	common.ListResponseCore
	Messages *[]Message `json:"messages"`
}
