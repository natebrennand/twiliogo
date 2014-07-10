package media

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	getUrl  = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages/%s/Media/%s.json" // takes an AccountSid, MessageSid & MediaSid
	listUrl = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages/%s/Media.json"    // takes an AccountSid & MessageSid

)

type MediaAccount struct {
	AccountSid string
	Token      string
	Client     http.Client
}

func (m MediaAccount) GetSid() string {
	return m.AccountSid
}
func (m MediaAccount) GetToken() string {
	return m.Token
}
func (m MediaAccount) GetClient() http.Client {
	return m.Client
}

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
	Uri         string          `json:"uri"`
	DateCreated common.JsonTime `json:"date_created"`
	DateUpdated common.JsonTime `json:"date_updated"`
	ContentType string          `json:"content-type"`
}

func (m *Media) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, m)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act MediaAccount) getMedia(destUrl string, resp *Media) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp, 200)
}

func (act MediaAccount) Get(mmsSid, mediaSid string) (Media, error) {
	var m Media
	if !validateMediaSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	} else if !validateMediaSid(mediaSid) {
		return m, errors.New("Invalid media sid")
	}

	err := act.getMedia(fmt.Sprintf(getUrl, act.AccountSid, mmsSid, mediaSid), &m)
	return m, err
}

type MediaList struct {
	// common.ListResponseCore
	MediaList *[]Media `json:"media_list"`
}

func (m *MediaList) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, m)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}

// Internal function for sending the post request to twilio.
func (act MediaAccount) getMediaList(destUrl string, resp *Media) error {
	// send get request to twilio
	return common.SendGetRequest(destUrl, act, resp, 200)
}

func (act MediaAccount) GetList(mmsSid string) (Media, error) {
	var m Media
	if !validateMediaSid(mmsSid) {
		return m, errors.New("Invalid mms message sid")
	}
	err := act.getMedia(fmt.Sprintf(getUrl, act.AccountSid, mmsSid), &m)
	return m, err
}
