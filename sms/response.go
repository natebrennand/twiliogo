package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
	"regexp"
)

func validateSmsSid(sid string) bool {
	match, _ := regexp.MatchString(`^(SM|MM)[0-9a-z]{32}$`, string(sid))
	return match
}

type Message struct {
	common.ResponseCore
	Body        string           `json:"body"`
	DateSent    common.JsonTime  `json:"date_sent"`
	NumSegments int              `json:"num_segments,string"`
	NumMedia    int              `json:"num_media,string"`
	Price       common.JsonPrice `json:"price"`
}

func (r *Message) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}

type MessageList struct {
	common.ListResponseCore
	Messages *[]Message `json:"messages"`
}

func (l *MessageList) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while reading json from buffer => %s", err.Error()))
	}
	err = json.Unmarshal(bodyBytes, l)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes)))
	}
	return nil
}
