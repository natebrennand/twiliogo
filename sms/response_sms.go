package sms

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
	NumSegments int              `json:"num_segments,string"`
	NumMedia    int              `json:"num_media,string"`
	Price       common.JsonPrice `json:"price"`
	DateCreated common.JsonTime  `json:"date_created"`
	DateSent    common.JsonTime  `json:"date_sent"`
	DateUpdated common.JsonTime  `json:"date_updated"`
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
