package common

import (
	"net/http"
)

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

// Standard request parameters for Twiml responses
// https://www.twilio.com/docs/api/twiml/sms/twilio_request#request-parameters
type StandardRequest struct {
	AccountSid string
	From       string
	To         string
	Location   *Location // Only sent when Twilio can look up the geographic data.
}

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
