package common

import (
	"net/http"
)

func ParseStandardRequest(req *http.Request) StandardRequest {
	var msgLocation *Location = nil
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
