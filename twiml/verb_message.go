package twiml

import "encoding/xml"

type message struct {
	XMLName int `xml:"Message"`
	*MessageOpts
	Body MessageBody
}

type MessageOpts struct {
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}

type MessageBody interface {
	xml.Marshaler
	Body(...string) MessageBody
	Media(...string) MessageBody
}

type MessageTwiml struct {
	baseTwiml
}

func addMessage(t twimlResponse, opts *MessageOpts, body MessageBody) {
	t.appendContents(&message{MessageOpts: opts, Body: body})
}

func addMessageText(t twimlResponse, opts *MessageOpts, text []string) {
	inner := new(MessageTwiml)
	addBody(inner, text)
	t.appendContents(&message{MessageOpts: opts, Body: inner})
}

func (t *MessageTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}

func (t *MessageTwiml) Body(text ...string) MessageBody {
	addBody(t, text)
	return t
}

func (t *MessageTwiml) Media(source ...string) MessageBody {
	addMedia(t, source)
	return t
}
