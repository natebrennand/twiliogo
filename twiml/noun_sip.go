package twiml

type sip struct {
	XMLName int `xml:"Sip"`
	*SipOpts
	Uri string `xml:",chardata"`
}

type SipOpts struct {
	Url      string `xml:"url,attr,omitempty"`
	Method   string `xml:"method,attr,omitempty"`
	Username string `xml:"username,attr,omitempty"`
	Password string `xml:"password,attr,omitempty"`
}

func addSip(t twimlResponse, opts *SipOpts, sips []string) {
	for _, uri := range sips {
		t.appendContents(&sip{0, opts, uri})
	}
}
