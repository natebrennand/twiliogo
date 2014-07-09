package voice

import (
	"github.com/natebrennand/twiliogo/common"
	"net/http"
)

// Represents the callback sent everytime the status of the call is updated.
// https://www.twilio.com/docs/api/rest/making-calls#status-callback-parameter
type Callback struct {
	CallDuration      string
	RecordingUrl      string
	RecordingSid      string
	RecordingDuration string
	common.StandardRequest
	CallSid       string
	CallStatus    string
	ApiVersion    string
	Direction     string
	ForwardedFrom string
	CallerName    string
}

// Creates a Callback struct from a form
func (cb *Callback) Parse(req *http.Request) error {

	var msgLocation *common.Location = nil
	if req.PostFormValue("FromCity") != "" { // ignore location data if possible
		msgLocation = &common.Location{
			FromCity:    req.PostFormValue("FromCity"),
			FromState:   req.PostFormValue("FromState"),
			FromZip:     req.PostFormValue("FromZip"),
			FromCountry: req.PostFormValue("FromCountry"),
			ToCity:      req.PostFormValue("ToCity"),
			ToState:     req.PostFormValue("ToState"),
			ToZip:       req.PostFormValue("ToZip"),
			ToCountry:   req.PostFormValue("ToCountry"),
		}
	}

	// Construct callback
	*cb = Callback{
		CallDuration:      req.PostFormValue("CallDuration"),
		RecordingUrl:      req.PostFormValue("RecordingUrl"),
		RecordingSid:      req.PostFormValue("RecordingSid"),
		RecordingDuration: req.PostFormValue("RecordingDuration"),
		CallSid:           req.PostFormValue("CallSid"),
		CallStatus:        req.PostFormValue("CallStatus"),
		ApiVersion:        req.PostFormValue("ApiVersion"),
		Direction:         req.PostFormValue("Direction"),
		ForwardedFrom:     req.PostFormValue("ForwardedFrom"),
		CallerName:        req.PostFormValue("CallerName"),
		StandardRequest: common.StandardRequest{
			AccountSid: req.PostFormValue("AccountSid"),
			From:       req.PostFormValue("From"),
			To:         req.PostFormValue("To"),
			Location:   msgLocation,
		},
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
