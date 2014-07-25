package conference

import (
	"net/http"
)

// Callback represents the callback sent by Twilio which contains the
// url for the recording.
type Callback struct {
	RecordingURL string
}

// Parse the form encoded callback into a Callback struct
func (cb *Callback) Parse(req *http.Request) error {
	*cb = Callback{
		RecordingURL: req.PostFormValue("RecordingUrl"),
	}
	return nil
}

// CallbackHandler creates a http handler to parse incoming callbacks
// and pass them into a channel for consumption.
func CallbackHandler(callbackChan chan Callback) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var cb Callback
		cb.Parse(req)
		resp.WriteHeader(http.StatusOK)
		go func() {
			callbackChan <- cb
		}()
	})
}
