package recording

import (
	"encoding/json"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Recording struct {
	Sid         string          `json:"sid"`
	DateCreated common.JSONTime `json:"date_created"`
	DateUpdated common.JSONTime `json:"date_updated"`
	AccountSid  string          `json:"account_sid"`
	CallSid     string          `json:"call_sid"`
	Duration    string          `json:"duration"`
	APIVersion  string          `json:"api_version"`
	URI         string          `json:"uri"`
}

type RecordingList struct {
	common.ListResponseCore
	Recordings *[]Recording `json:"recordings"`
}

func validateRecSid(sid string) bool {
	match, _ := regexp.MatchString(`^RE[0-9a-z]{32}$`, string(sid))
	return match
}

func (r *Recording) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading json from buffer => %s", err.Error())
	}

	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes))
	}
	return nil
}

func (r *RecordingList) Build(resp *http.Response) error {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error while reading json from buffer => %s", err.Error())
	}

	err = json.Unmarshal(bodyBytes, r)
	if err != nil {
		return fmt.Errorf("Error while decoding json => %s, recieved msg => %s", err.Error(), string(bodyBytes))
	}
	return nil
}
