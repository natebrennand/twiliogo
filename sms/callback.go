package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"strconv"
)

type Media struct {
	ContentType string
	Url         string
}

// Represents the callback sent everytime the status of the message is updated.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#status-callback-parameter for more detaiils
type Callback struct {
	MessageSid    string
	SmsSid        string
	Body          string
	NumMedia      int
	MessageStatus string
	ErrorCode     string
	MediaList     []Media
	common.StandardRequest
}

// Parses the form encoded callback into a Callback struct
func parseCallback(req *http.Request, cb *Callback) error {
	numMediaString := req.PostFormValue("NumMedia")
	numMedia, err := strconv.Atoi(numMediaString)
	if err != nil && numMediaString != "" {
		return errors.New(fmt.Sprintf("Error parsing NumMedia => %s", err.Error()))
	}

	// creates an array of Media Contents (typically empty)
	mediaArray := make([]Media, numMedia)
	for i := 0; i < numMedia; i++ {
		mediaArray[i] = Media{
			ContentType: req.PostFormValue(fmt.Sprintf("MediaContentType%d", i)),
			Url:         req.PostFormValue(fmt.Sprintf("MediaUrl%d", i)),
		}
	}

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

	*cb = Callback{
		MessageSid:    req.PostFormValue("MessageSid"),
		SmsSid:        req.PostFormValue("SmsSid"),
		Body:          req.PostFormValue("Body"),
		NumMedia:      numMedia,
		MessageStatus: req.PostFormValue("MessageStatus"),
		ErrorCode:     req.PostFormValue("ErrorCode"),
		MediaList:     mediaArray,
		StandardRequest: common.StandardRequest{
			AccountSid: req.PostFormValue("AccountSid"),
			From:       req.PostFormValue("From"),
			To:         req.PostFormValue("To"),
			Location:   msgLocation,
		},
	}
	return nil
}

// Creates a http Handler that will parse a Twilio callback and send it into the provided channel.
func CallbackHandler(callbackChan chan Callback) http.HandlerFunc {
	var cb Callback
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		err := parseCallback(req, &cb)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.WriteHeader(http.StatusOK)
		callbackChan <- cb
	})
}
