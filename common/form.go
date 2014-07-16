package common

import (
	"net/http"
)

// Location holds information about a message that Twilio will instantiate if they
// can determine it.
type Location struct {
	FromCity    string
	FromState   string
	FromZip     string
	FromCountry string
	ToCity      string
	ToState     string
	ToZip       string
	ToCountry   string
}

// StandardRequest parameters for Twiml responses
// https://www.twilio.com/docs/api/twiml/sms/twilio_request#request-parameters
type StandardRequest struct {
	AccountSid string
	From       string
	To         string
	Location   *Location // Only sent when Twilio can look up the geographic data.
}

// ParseStandardRequest parses out the standard request parameters from an http request.
func ParseStandardRequest(req *http.Request) StandardRequest {
	var msgLocation *Location
	if req.PostFormValue("FromCity") != "" { // ignore location data if possible
		msgLocation = &Location{
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

	return StandardRequest{
		AccountSid: req.PostFormValue("AccountSid"),
		From:       req.PostFormValue("From"),
		To:         req.PostFormValue("To"),
		Location:   msgLocation,
	}
}
