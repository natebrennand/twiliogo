package twiml

type say struct {
	XMLName int `xml:"Say"`
	*SayOpts
	Sentence string `xml:",chardata"`
}

type SayOpts struct {
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
}

func addSay(t twimlResponse, opts *SayOpts, lines []string) {
	for _, line := range lines {
		t.appendContents(&say{0, opts, line})
	}
}
