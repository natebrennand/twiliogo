package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"regexp"
)

var media = struct {
	Get, List string
}{
	Get:  "/2010-04-01/Accounts/%s/Messages/%s/Media/%s.json", // takes an AccountSid, MessageSid & MediaSid,
	List: "/2010-04-01/Accounts/%s/Messages/%s/Media.json",    // takes an AccountSid & MessageSid
}

var (
	validateMediaSid = regexp.MustCompile(`^ME[0-9a-z]{32}$`).MatchString
	validateMmsSid   = regexp.MustCompile(`^MM[0-9a-z]{32}$`).MatchString
)

// Media represents a Media resource
type Media struct {
	AccountSid  string          `json:"account_sid"`
	ParentSid   string          `json:"parent_sid,omitempty"`
	Sid         string          `json:"sid"`
	URI         string          `json:"uri"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	ContentType string          `json:"content-type"`
}

// GetMedia returns the information for a piece of media given it's parent MMS's sid and the sid
// for the exact piece of media.
func (act Account) GetMedia(mmsSid, mediaSid string) (Media, error) {
	var m Media
	if !validateMmsSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	} else if !validateMediaSid(mediaSid) {
		return m, errors.New("Invalid media sid")
	}

	err := common.SendGetRequest(fmt.Sprintf(media.Get, act.AccountSid, mmsSid, mediaSid), act, &m)
	return m, err
}

// MediaList contains a list of Media instances as well as paging information for further querying.
// MediaList does not have a Next() method because there is a constant limit of the number of
// media items that may be associated with a Media resource.
type MediaList struct {
	common.ListResponseCore
	MediaList *[]Media `json:"media_list"`
}

// GetMediaList returns a list of media contained in the MMS message identified by the provided sid.
func (act Account) GetMediaList(mmsSid string) (MediaList, error) {
	var m MediaList
	if !validateMediaSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	}
	err := common.SendGetRequest(fmt.Sprintf(media.List, act.AccountSid, mmsSid), act, &m)
	return m, err
}
