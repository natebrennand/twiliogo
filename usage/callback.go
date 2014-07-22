package usage

import (
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"time"
)

// TriggerCallback contains the information sent by a usage trigger.
type TriggerCallback struct {
	AccountSid        string          `json:"account_sid"`
	UsageTriggerSid   string          `json:"usage_trigger_sid"`
	Recurring         string          `json:"recurring"`
	UsageCategory     string          `json:"usage_category"`
	TriggerBy         string          `json:"trigger_by"`
	TriggerValue      string          `json:"trigger_value"`
	CurrentUsageValue string          `json:"current_usage_value"`
	UsageRecordURI    string          `json:"usage_record_uri"`
	URI               string          `json:"uri"`
	DateFired         common.JSONTime `json:"date_fired"`
	DateCreated       common.JSONTime `json:"date_created"`
	DateUpdated       common.JSONTime `json:"date_updated"`
}

// Parse the form encoded callback into a TriggerCallback struct
func (tcb *TriggerCallback) Parse(req *http.Request) error {
	dateFired, err := time.Parse(common.GMTTimeLayout, req.PostFormValue("DateFired"))
	if err != nil {
		return fmt.Errorf("Problem parsing DateFired => %s", err.Error())
	}
	dateCreated, err := time.Parse(common.GMTTimeLayout, req.PostFormValue("DateCreated"))
	if err != nil {
		return fmt.Errorf("Problem parsing DateCreated => %s", err.Error())
	}
	dateUpdated, err := time.Parse(common.GMTTimeLayout, req.PostFormValue("DateUpdated"))
	if err != nil {
		return fmt.Errorf("Problem parsing DateUpdated => %s", err.Error())
	}

	*tcb = TriggerCallback{
		AccountSid:        req.PostFormValue("AccountSid"),
		UsageTriggerSid:   req.PostFormValue("UsageTriggerSid"),
		Recurring:         req.PostFormValue("Recurring"),
		UsageCategory:     req.PostFormValue("UsageCategory"),
		TriggerBy:         req.PostFormValue("TriggerBy"),
		TriggerValue:      req.PostFormValue("TriggerValue"),
		CurrentUsageValue: req.PostFormValue("CurrentUsageValue"),
		UsageRecordURI:    req.PostFormValue("UsageRecordURI"),
		URI:               req.PostFormValue("URI"),
		DateFired:         common.JSONTime{Time: dateFired},
		DateCreated:       common.JSONTime{Time: dateCreated},
		DateUpdated:       common.JSONTime{Time: dateUpdated},
	}
	return nil
}

// CallbackHandler creates a http Handler that will parse a Twilio callback and send it into the provided channel.
func CallbackHandler(callbackChan chan TriggerCallback) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var cb TriggerCallback
		err := cb.Parse(req)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.WriteHeader(http.StatusOK)
		// start seperate goroutine to allow http request to return.
		go func() {
			callbackChan <- cb
		}()
	})
}
