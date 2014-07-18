package twiml

type queue struct {
	XMLName int `xml:"Queue"`
	*QueueOpts
	Name *string `xml:",chardata"`
}

// QueueOpts sets the Twiml callback for a queue
//
// https://www.twilio.com/docs/api/twiml/queue
type QueueOpts struct {
	URL    string `xml:"url,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
}

func addQueue(t twimlResponse, opts *QueueOpts, name *string) {
	t.appendContents(&queue{QueueOpts: opts, Name: name})
}
