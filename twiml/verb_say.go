package twiml

type say struct {
	XMLName int `xml:"Say"`
	*SayOpts
	Sentence string `xml:",chardata"`
}

// SayOpts alters the attributes of the automatic voice.
//
// https://www.twilio.com/docs/api/twiml/say
type SayOpts struct {
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
}

func addSay(t twimlResponse, opts *SayOpts, lines []string) {
	for _, line := range lines {
		t.appendContents(&say{SayOpts: opts, Sentence: line})
	}
}
