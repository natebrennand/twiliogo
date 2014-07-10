package voice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Call struct {
	common.ResponseCore
	Price          common.JsonPrice `json:"price"`
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      common.JsonTime `json:"start_time"`
	EndTime        common.JsonTime `json:"end_time"`
	Duration       string          `json:"duration"`
	AnsweredBy     string          `json:"answered_by"`
	ForwardedFrom  string          `json:"fowarded_from"`
	CallerName     string          `json:"caller_name"`
}

func validateCallSid(sid string) bool {
	match, _ := regexp.MatchString(`^CA[0-9a-z]{32}$`, string(sid))
	return match
}

func (r *Call) Build(resp *http.Response) error {
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

type CallList struct {
	//common.ListResponseCore
	Calls *[]Call `json:"calls"`
}

func (l *CallList) Build(resp *http.Response) error {
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
