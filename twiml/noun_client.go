package twiml

type client struct {
	XMLName int `xml:"Client"`
	*ClientOpts
	Name string `xml:",chardata"`
}

// ClientOpts allows setting of a Twiml to be called upon the Client entering the call.
//
// https://www.twilio.com/docs/api/twiml/client
type ClientOpts struct {
	URL                  string `xml:"url,attr,omitempty"`
	Method               string `xml:"method,attr,omitempty"`
	StatusCallbackEvent  string `xml:"statusCallbackEvent,attr,omitempty"`
	StatusCallback       string `xml:"statusCallback,attr,omitempty"`
	StatusCallbackMethod string `xml:"statusCallbackMethod,attr,omitempty"`
}

func addClient(t twimlResponse, opts *ClientOpts, clients []string) {
	for _, name := range clients {
		t.appendContents(&client{ClientOpts: opts, Name: name})
	}
}
