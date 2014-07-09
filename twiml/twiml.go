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

type TwimlInterface interface {
	Render() ([]byte, error)
	RenderReader() (io.Reader, error)
	Say(SayOpts, ...string) TwimlInterface
	Play(PlayOpts, ...string) TwimlInterface
	// Dial(...string) TwimlInterface
	Record(RecordOpts, string) TwimlInterface
	// Gather(...string) TwimlInterface
	// Sms(...string) TwimlInterface
	// Hangup(...string) TwimlInterface
	// Queue(...string) TwimlInterface
	// Redirect(...string) TwimlInterface
	// Pause(...string) TwimlInterface
	// Reject(...string) TwimlInterface
}

type Response struct {
	contents []interface{}
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

// Say some text during a phone call.
//
// https://www.twilio.com/docs/api/twiml/say
func (t *Response) Say(opts SayOpts, lines ...string) TwimlInterface {
	for _, line := range lines {
		t.contents = append(t.contents, &say{0, line, &opts})
	}

	return t
}

// Play an audio file during a phone call.
//
// https://www.twilio.com/docs/api/twiml/play
func (t *Response) Play(opts PlayOpts, urls ...string) TwimlInterface {
	for _, url := range urls {
		t.contents = append(t.contents, &play{0, url, &opts})
	}

	return t
}

// Record audio during a phone call.
//
// https://www.twilio.com/docs/api/twiml/record
func (t *Response) Record(opts RecordOpts, action string) TwimlInterface {
	newRecord := &record{0, action, &opts}
	t.contents = append(t.contents, newRecord)

	return t
}
