package sms

import (
	"fmt"
	"github.com/natebrennand/twiliogo/common"
	"net/http"
	"strconv"
)

// MediaReference contains the content type and URL for a MMS Media object
type MediaReference struct {
	ContentType string
	URL         string
}

// Callback represents the callback sent everytime the status of the message is updated.
//
// Visit https://www.twilio.com/docs/api/rest/sending-messages#status-callback-parameter for more details
type Callback struct {
	MessageSid    string
	SmsSid        string
	Body          string
	NumMedia      int
	MessageStatus string
	ErrorCode     string
	MediaList     []MediaReference
	common.StandardRequest
}

// Parse the form encoded callback into a Callback struct
func (cb *Callback) Parse(req *http.Request) error {
	numMediaString := req.PostFormValue("NumMedia")
	numMedia, err := strconv.Atoi(numMediaString)
	if err != nil && numMediaString != "" {
		return fmt.Errorf("Error parsing NumMedia => %s", err.Error())
	}

	// creates an array of MediaReference Contents (typically empty)
	mediaArray := make([]MediaReference, numMedia)
	for i := 0; i < numMedia; i++ {
		mediaArray[i] = MediaReference{
			ContentType: req.PostFormValue(fmt.Sprintf("MediaContentType%d", i)),
			URL:         req.PostFormValue(fmt.Sprintf("MediaURL%d", i)),
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

// CallbackHandler creates a http Handler that will parse a Twilio callback and send it into the provided channel.
func CallbackHandler(callbackChan chan Callback) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var cb Callback
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
