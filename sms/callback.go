package sms

import (
	"errors"
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"strconv"
)

// Represents the callback sent everytime the status of the message is updated.
// Visit https://www.twilio.com/docs/api/rest/sending-messages#status-callback-parameter for more detaiils
type Callback struct {
	MessageSid    string
	SmsSid        string
	Body          string
	NumMedia      int
	MessageStatus string
	ErrorCode     string
	common.StandardRequest
}

func parseCallback(req *http.Request, cb *Callback) error {
	numMediaString := req.PostFormValue("NumMedia")
	numMedia, err := strconv.Atoi(numMediaString)
	if err != nil && numMediaString != "" {
		return errors.New(fmt.Sprintf("Error parsing NumMedia => %s", err.Error()))
	}

	mediaArray := make([]common.Media, numMedia)
	for i := 0; i < numMedia; i++ {
		mediaArray[i] = common.Media{
			ContentType: req.PostFormValue(fmt.Sprintf("MediaContentType%d", i)),
			Url:         req.PostFormValue(fmt.Sprintf("MediaUrl%d", i)),
		}
	}

	*cb = Callback{
		MessageSid:    req.PostFormValue("MessageSid"),
		SmsSid:        req.PostFormValue("SmsSid"),
		Body:          req.PostFormValue("Body"),
		NumMedia:      numMedia,
		MessageStatus: req.PostFormValue("MessageStatus"),
		ErrorCode:     req.PostFormValue("ErrorCode"),
		StandardRequest: common.StandardRequest{
			AccountSid:  req.PostFormValue("AccountSid"),
			From:        req.PostFormValue("From"),
			To:          req.PostFormValue("To"),
			MediaList:   mediaArray,
			FromCity:    req.PostFormValue("FromCity"),
			FromState:   req.PostFormValue("FromState"),
			FromZip:     req.PostFormValue("FromZip"),
			FromCountry: req.PostFormValue("FromCountry"),
			ToCity:      req.PostFormValue("ToCity"),
			ToState:     req.PostFormValue("ToState"),
			ToZip:       req.PostFormValue("ToZip"),
			ToCountry:   req.PostFormValue("ToCountry"),
		},
	}
	return nil
}

func CallbackHandler(callbackChan chan Callback) http.HandlerFunc {
	var cb Callback
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(200)
		err := parseCallback(req, &cb)
		if err != nil {
			resp.WriteHeader(400)
			return
		}
		callbackChan <- cb
	})
}
