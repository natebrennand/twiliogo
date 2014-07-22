package twiml

type enqueue struct {
	XMLName int `xml:"Enqueue"`
	*EnqueueOpts
	Queue *string `xml:",chardata"`
}

// EnqueueOpts sets the Twiml instructions
type EnqueueOpts struct {
	Action        string `xml:"action,attr,omitempty"`
	Method        string `xml:"method,attr,omitempty"`
	WaitURL       string `xml:"waitUrl,attr,omitempty"`
	WaitURLMethod string `xml:"waitUrlMethod,attr,omitempty"`
}

func addEnqueue(t twimlResponse, opts *EnqueueOpts, queue *string) {
	t.appendContents(&enqueue{EnqueueOpts: opts, Queue: queue})
}
