package twiml

import "encoding/xml"

type TwimlInterface interface {
	Render() (string, error)
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

type Twiml struct {
	contents []interface{}
}

// Returns a string representation of the TwiML generated previous calls to this object.
func (t Twiml) Render() (string, error) {
	result, err := xml.MarshalIndent(t.contents, "", "  ")
	return string(result), err
}

// Say some text during a phone call.
// https://www.twilio.com/docs/api/twiml/say
func (t *Twiml) Say(opts SayOpts, lines ...string) TwimlInterface {
	for _, line := range lines {
		t.contents = append(t.contents, &say{0, line, &opts})
	}

	return t
}

// Play an audio file during a phone call.
// https://www.twilio.com/docs/api/twiml/play
func (t *Twiml) Play(opts PlayOpts, urls ...string) TwimlInterface {
	for _, url := range urls {
		t.contents = append(t.contents, &play{0, url, &opts})
	}

	return t
}

// Record audio during a phone call.
// https://www.twilio.com/docs/api/twiml/record
func (t *Twiml) Record(opts RecordOpts, action string) TwimlInterface {
	newRecord := &record{0, action, &opts}
	t.contents = append(t.contents, newRecord)

	return t
}
