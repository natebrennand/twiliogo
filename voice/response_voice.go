package voice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"io/ioutil"
	"net/http"
)

type Response struct {
	common.ResponseCore
	Price          common.JsonPrice `json:"price"`
	DateCreated    common.JsonTime  `json:"date_created"`
	DateUpdated    common.JsonTime  `json:"date_updated"`
	ParentCallSid  string
	PhoneNumberSid string
	StartTime      common.JsonTime `json:"start_time"`
	EndTime        common.JsonTime `json:"end_time"`
	Duration       float64         `json:"duration"`
	AnsweredBy     string          `json:"answered_by"`
	ForwardedFrom  string          `json:"fowarded_from"`
	CallerName     string          `json:"caller_name"`
}

func (r *Response) Build(resp *http.Response) error {
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
