package twiml

import (
	"bytes"
	"encoding/xml"
	"io"
)

var (
	preTwiml  = []byte(xml.Header + "\n<Response>\n")
	postTwiml = []byte("\n</Response>\n")
)

type twimlResponse interface {
	xml.Marshaler
	getContents() []interface{}
	appendContents(interface{})
}

type TwimlInterface interface {
	xml.Marshaler
	Render() ([]byte, error)
	RenderReader() (io.Reader, error)
	Say(SayOpts, ...string) TwimlInterface
	Play(PlayOpts, ...string) TwimlInterface
	Dial(DialOpts, DialBody) TwimlInterface
	Record(RecordOpts, string) TwimlInterface
	Gather(GatherOpts, GatherBody) TwimlInterface
	// Sms(...string) TwimlInterface
	// Hangup(...string) TwimlInterface
	// Queue(...string) TwimlInterface
	// Redirect(...string) TwimlInterface
	Pause(int) TwimlInterface
	// Reject(...string) TwimlInterface
}

type Response struct {
	baseTwiml
}

// Returns a TwiML representation of the previous calls on the struct as a byte
// slice.
func (t *Response) Render() (result []byte, err error) {
	result, err = xml.MarshalIndent(t, "\t", "\t")
	if err != nil {
		return
	}
	result = append(preTwiml, result...)
	result = append(result, postTwiml...)
	return
}

// Returns a TwiML representation of the previous calls on the struct, enclosed
// in a Reader interface.
func (t *Response) RenderReader() (io.Reader, error) {
	result, err := t.Render()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(result), nil
}
