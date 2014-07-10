package twiml

import "encoding/xml"

type gather struct {
	XMLName int `xml:"Gather"`
	*GatherOpts
	Body *GatherBody
}

type GatherOpts struct {
	Action      string `xml:"action,attr,omitempty"`
	Method      string `xml:"method,attr,omitempty"`
	Timeout     int    `xml:"timeout,attr,omitempty"`
	FinishOnKey string `xml:"finishOnKey,attr,omitempty"`
	NumDigits   int    `xml:"numDigits,attr,omitempty"`
}

type GatherBody interface {
	xml.Marshaler
	Say(SayOpts, ...string) GatherBody
	Play(PlayOpts, ...string) GatherBody
	Pause(int) GatherBody
}

type GatherTwiml struct {
	baseTwiml
}

func addGather(t twimlResponse, opts *GatherOpts, body *GatherBody) {
	t.appendContents(&gather{GatherOpts: opts, Body: body})
}

func (t *GatherTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}

func (t *GatherTwiml) Play(opts PlayOpts, urls ...string) GatherBody {
	addPlay(t, &opts, urls)
	return t
}

func (t *GatherTwiml) Say(opts SayOpts, lines ...string) GatherBody {
	addSay(t, &opts, lines)
	return t
}

func (t *GatherTwiml) Pause(length int) GatherBody {
	addPause(t, length)
	return t
}
