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

var (
	validateMediaSid = regexp.MustCompile(`^ME[0-9a-z]{32}$`).MatchString
	validateMmsSid   = regexp.MustCompile(`^MM[0-9a-z]{32}$`).MatchString
)

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
func (act Account) getMedia(destURL string, resp *Media) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp)
}

func (act Account) GetMedia(mmsSid, mediaSid string) (Media, error) {
	var m Media
	if !validateMmsSid(mmsSid) {
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
func (act Account) getMediaList(destURL string, resp *MediaList) error {
	// send get request to twilio
	return common.SendGetRequest(destURL, act, resp)
}

func (act Account) GetMediaList(mmsSid string) (MediaList, error) {
	var m MediaList
	if !validateMediaSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	}
	err := act.getMediaList(fmt.Sprintf(mediaListURL, act.AccountSid, mmsSid), &m)
	return m, err
}
