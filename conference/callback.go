package conference

import (
	"net/http"
)

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

func CallbackHandler(callbackChan chan Callback) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var cb Callback
		err := cb.Parse(req)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		resp.WriteHeader(http.StatusOK)
		callbackChan <- cb
	})
}
