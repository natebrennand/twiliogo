package twiml

import "encoding/xml"

type dial struct {
	XMLName int `xml:"Dial"`
	*DialOpts
	Body *DialBody
}

type DialOpts struct {
	Action       string `xml:"action,attr,omitempty"`
	Method       string `xml:"method,attr,omitempty"`
	Timeout      int    `xml:"timeout,attr,omitempty"`
	HangupOnStar bool   `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit    int    `xml:"timeLimit,attr,omitempty"`
	CallerId     string `xml:"callerId,attr,omitempty"`
	Record       string `xml:"record,attr,omitempty"`
	Trim         string `xml:"trim,attr,omitempty"`
}

type DialBody interface {
	xml.Marshaler
	Number(NumberOpts, ...string) DialBody
	Sip(SipOpts, ...string) DialBody
	Client(ClientOpts, ...string) DialBody
	Conference(ConferenceOpts, string) DialBody
	Queue(QueueOpts, string) DialBody
}

type DialTwiml struct {
	baseTwiml
}

func addDial(t twimlResponse, opts *DialOpts, body *DialBody) {
	t.appendContents(&dial{DialOpts: opts, Body: body})
}

func (t *DialTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}

func (t *DialTwiml) Number(opts NumberOpts, numbers ...string) DialBody {
	addNumber(t, &opts, numbers)
	return t
}

func (t *DialTwiml) Sip(opts SipOpts, sips ...string) DialBody {
	addSip(t, &opts, sips)
	return t
}

func (t *DialTwiml) Client(opts ClientOpts, clients ...string) DialBody {
	addClient(t, &opts, clients)
	return t
}

func (t *DialTwiml) Conference(opts ConferenceOpts, name string) DialBody {
	addConference(t, &opts, &name)
	return t
}

func (t *DialTwiml) Queue(opts QueueOpts, name string) DialBody {
	addQueue(t, &opts, &name)
	return t
}
