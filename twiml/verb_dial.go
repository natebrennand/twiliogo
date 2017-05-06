package twiml

import "encoding/xml"

type dial struct {
	XMLName int `xml:"Dial"`
	*DialOpts
	Body *dialBody
}

// DialOpts configures the behavoir of the call
//
// https://www.twilio.com/docs/api/twiml/dial
type DialOpts struct {
	Action                        string `xml:"action,attr,omitempty"`
	Method                        string `xml:"method,attr,omitempty"`
	Timeout                       int    `xml:"timeout,attr,omitempty"`
	HangupOnStar                  bool   `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit                     int    `xml:"timeLimit,attr,omitempty"`
	CallerID                      string `xml:"callerId,attr,omitempty"`
	Record                        string `xml:"record,attr,omitempty"`
	Trim                          string `xml:"trim,attr,omitempty"`
	RecordingStatusCallback       string `xml:"recordingStatusCallback,attr,omitempty"`
	RecordingStatusCallbackMethod string `xml:"recordingStatusCallbackMethod,attr,omitempty"`
	RingTone                      string `xml:"ringTone,attr,omitempty"`
}

type dialBody interface {
	xml.Marshaler
	Number(NumberOpts, ...string) dialBody
	Sip(SipOpts, ...string) dialBody
	Client(ClientOpts, ...string) dialBody
	Conference(ConferenceOpts, string) dialBody
	Queue(QueueOpts, string) dialBody
}

// DialTwiml is the container for the Dial operations.
type DialTwiml struct {
	baseTwiml
}

func addDial(t twimlResponse, opts *DialOpts, body *dialBody) {
	t.appendContents(&dial{DialOpts: opts, Body: body})
}

// MarshalXML implements an interface for proper rendering.
func (t *DialTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}

// Number adds a number to be called.
func (t *DialTwiml) Number(opts NumberOpts, numbers ...string) dialBody {
	addNumber(t, &opts, numbers)
	return t
}

// Sip adds a client to be connected via SIP
func (t *DialTwiml) Sip(opts SipOpts, sips ...string) dialBody {
	addSip(t, &opts, sips)
	return t
}

// Client adds a pre-set client to be connected
func (t *DialTwiml) Client(opts ClientOpts, clients ...string) dialBody {
	addClient(t, &opts, clients)
	return t
}

// Conference sets a conference for >1 clients to connect.
func (t *DialTwiml) Conference(opts ConferenceOpts, name string) dialBody {
	addConference(t, &opts, &name)
	return t
}

// Queue identifies a queue that the call should be connected to.
func (t *DialTwiml) Queue(opts QueueOpts, name string) dialBody {
	addQueue(t, &opts, &name)
	return t
}
