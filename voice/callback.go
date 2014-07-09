package voice

import (
	"github.com/natebrennand/twiliogo/common"
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
