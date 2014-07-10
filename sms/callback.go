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
func (cb *Callback) Parse(req *http.Request) error {
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

	*cb = Callback{
		MessageSid:      req.PostFormValue("MessageSid"),
		SmsSid:          req.PostFormValue("SmsSid"),
		Body:            req.PostFormValue("Body"),
		NumMedia:        numMedia,
		MessageStatus:   req.PostFormValue("MessageStatus"),
		ErrorCode:       req.PostFormValue("ErrorCode"),
		MediaList:       mediaArray,
		StandardRequest: common.ParseStandardRequest(req),
	}
	return nil
}

// Creates a http Handler that will parse a Twilio callback and send it into the provided channel.
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
