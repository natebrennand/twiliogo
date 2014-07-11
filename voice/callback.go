package voice

import (
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"strconv"
)

// Represents the callback sent everytime the status of the call is updated.
// https://www.twilio.com/docs/api/rest/making-calls#status-callback-parameter
type Callback struct {
	CallDuration      int `json:"call_duration"`
	RecordingURL      string
	RecordingSid      string
	RecordingDuration int `json:"recording_duration"`
	common.StandardRequest
	CallSid       string
	CallStatus    string
	APIVersion    string
	Direction     string
	ForwardedFrom string
	CallerName    string
}

// Creates a Callback struct from a form
func (cb *Callback) Parse(req *http.Request) error {
	callDurString := req.PostFormValue("CallDuration")

	var err error
	var callDur = 0
	var recDur = 0

	if callDurString != "" {
		callDur, err = strconv.Atoi(callDurString)
		if err != nil {
			return fmt.Errorf("Error parsing CallDuration => %s", err.Error())
		}
	}

	recDurString := req.PostFormValue("RecordingDuration")
	if recDurString != "" {
		recDur, err = strconv.Atoi(recDurString)
		if err != nil {
			return fmt.Errorf("Error parsing RecordingDuration => %s", err.Error())
		}
	}

	// Construct callback
	*cb = Callback{
		CallDuration:      callDur,
		RecordingURL:      req.PostFormValue("RecordingURL"),
		RecordingSid:      req.PostFormValue("RecordingSid"),
		RecordingDuration: recDur,
		CallSid:           req.PostFormValue("CallSid"),
		CallStatus:        req.PostFormValue("CallStatus"),
		APIVersion:        req.PostFormValue("APIVersion"),
		Direction:         req.PostFormValue("Direction"),
		ForwardedFrom:     req.PostFormValue("ForwardedFrom"),
		CallerName:        req.PostFormValue("CallerName"),
		StandardRequest:   common.ParseStandardRequest(req),
	}

	return nil
}

func CallbackHandler(callbackChan chan Callback) http.HandlerFunc {
	var cb Callback
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		err := cb.Parse(req)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.WriteHeader(http.StatusOK)
		callbackChan <- cb
	})
}
