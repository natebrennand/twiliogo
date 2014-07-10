package twiml

type client struct {
	XMLName int `xml:"Client"`
	*ClientOpts
	Name string `xml:",chardata"`
}

type ClientOpts struct {
	Url    string `xml:"url,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
}

func addClient(t twimlResponse, opts *ClientOpts, clients []string) {
	for _, name := range clients {
		t.appendContents(&client{ClientOpts: opts, Name: name})
	}
}
