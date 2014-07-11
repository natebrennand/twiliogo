package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

const (
	mediaGetURL  = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages/%s/Media/%s.json" // takes an AccountSid, MessageSid & MediaSid
	mediaListURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages/%s/Media.json"    // takes an AccountSid & MessageSid

)

func validateMediaSid(sid string) bool {
	match, _ := regexp.MatchString(`^ME[0-9a-z]{32}$`, string(sid))
	return match
}

func validateMmsSid(sid string) bool {
	match, _ := regexp.MatchString(`^MM[0-9a-z]{32}$`, string(sid))
	return match
}

type Media struct {
	AccountSid  string          `json:"account_sid"`
	ParentSid   string          `json:"parent_sid,omitempty"`
	Sid         string          `json:"sid"`
	URI         string          `json:"uri"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	ContentType string          `json:"content-type"`
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) getMedia(destURL string, resp *Media) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp, 200)
}

func (act SmsAccount) GetMedia(mmsSid, mediaSid string) (Media, error) {
	var m Media
	if !validateMediaSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	} else if !validateMediaSid(mediaSid) {
		return m, errors.New("Invalid media sid")
	}

	err := act.getMedia(fmt.Sprintf(mediaGetURL, act.AccountSid, mmsSid, mediaSid), &m)
	return m, err
}

type MediaList struct {
	common.ListResponseCore
	MediaList *[]Media `json:"media_list"`
}

// Internal function for sending the post request to twilio.
func (act SmsAccount) getMediaList(destURL string, resp *MediaList) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp, 200)
}

func (act SmsAccount) GetMediaList(mmsSid string) (MediaList, error) {
	var m MediaList
	if !validateMediaSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	}
	err := act.getMediaList(fmt.Sprintf(mediaListURL, act.AccountSid, mmsSid), &m)
	return m, err
}
