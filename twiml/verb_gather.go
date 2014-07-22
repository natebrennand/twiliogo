package twiml

import "encoding/xml"

type gather struct {
	XMLName int `xml:"Gather"`
	*GatherOpts
	Body *gatherBody
}

// GatherOpts sets the options on a gather verb.
type GatherOpts struct {
	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     int    `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int    `xml:"numDigits,attr,omitempty"`
}

type gatherBody interface {
	xml.Marshaler
	Say(SayOpts, ...string) gatherBody
	Play(PlayOpts, ...string) gatherBody
	Pause(int) gatherBody
}

// GatherTwiml contains the TwiML in a gather block.
type GatherTwiml struct {
	baseTwiml
}

func addGather(t twimlResponse, opts *GatherOpts, body *gatherBody) {
	t.appendContents(&gather{GatherOpts: opts, Body: body})
}

// MarshalXML implements the XML marshaler interface
func (t *GatherTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}

// Play adds a Play instruction to the gather block
func (t *GatherTwiml) Play(opts PlayOpts, urls ...string) gatherBody {
	addPlay(t, &opts, urls)
	return t
}

// Say adds a Say instruction to the gather block
func (t *GatherTwiml) Say(opts SayOpts, lines ...string) gatherBody {
	addSay(t, &opts, lines)
	return t
}

// Pause adds a Pause instruction to the gather block
func (t *GatherTwiml) Pause(length int) gatherBody {
	addPause(t, length)
	return t
}
