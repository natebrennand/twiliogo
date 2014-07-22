package twiml

import "encoding/xml"

type message struct {
	XMLName int `xml:"Message"`
	*MessageOpts
	Body messageBody
}

// MessageOpts sets the message attributes.
type MessageOpts struct {
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}

type messageBody interface {
	xml.Marshaler
	Body(...string) messageBody
	Media(...string) messageBody
}

// MessageTwiml containst the TwiML Message block
type MessageTwiml struct {
	baseTwiml
}

func addMessage(t twimlResponse, opts *MessageOpts, body messageBody) {
	t.appendContents(&message{MessageOpts: opts, Body: body})
}

func addMessageText(t twimlResponse, opts *MessageOpts, text []string) {
	inner := new(MessageTwiml)
	addBody(inner, text)
	t.appendContents(&message{MessageOpts: opts, Body: inner})
}

// MarshalXML implements the XML interface for proper rendering
func (t *MessageTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}

// Body sets the body of the message
func (t *MessageTwiml) Body(text ...string) messageBody {
	addBody(t, text)
	return t
}

// Media adds media to the message
func (t *MessageTwiml) Media(source ...string) messageBody {
	addMedia(t, source)
	return t
}
