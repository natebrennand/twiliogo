package twiml

type enqueue struct {
	XMLName int `xml:"Enqueue"`
	*EnqueueOpts
	Url *string `xml:",chardata"`
}

type EnqueueOpts struct {
	Action        string `xml:"action,attr,omitempty"`
	Method        string `xml:"method,attr,omitempty"`
	WaitUrl       string `xml:"waitUrl,attr,omitempty"`
	WaitUrlMethod string `xml:"waitUrlMethod,attr,omitempty"`
}

func addEnqueue(t twimlResponse, opts *EnqueueOpts, queue *string) {
	t.appendContents(&enqueue{0, opts, queue})
}
