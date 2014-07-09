package twiml

type queue struct {
	XMLName int `xml:"Queue"`
	*QueueOpts
	Name *string `xml:",chardata"`
}

type QueueOpts struct {
	Url    string `xml:"url,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
}

func addQueue(t twimlResponse, opts *QueueOpts, name *string) {
	t.appendContents(&queue{QueueOpts: opts, Name: name})
}
