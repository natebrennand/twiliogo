package twiml

type sip struct {
	XMLName int `xml:"Sip"`
	*SipOpts
	URI string `xml:",chardata"`
}

// SipOpts sets the callback for a SIP client.
//
// https://www.twilio.com/docs/api/twiml/sip
type SipOpts struct {
	URL      string `xml:"url,attr,omitempty"`
	Method   string `xml:"method,attr,omitempty"`
	Username string `xml:"username,attr,omitempty"`
	Password string `xml:"password,attr,omitempty"`
}

func addSip(t twimlResponse, opts *SipOpts, sips []string) {
	for _, uri := range sips {
		t.appendContents(&sip{SipOpts: opts, URI: uri})
	}
}
