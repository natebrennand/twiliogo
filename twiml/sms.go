package twiml

type sms struct {
	XMLName int `xml:"Sms"`
	*SmsOpts
	Text *string `xml:",chardata"`
}

type SmsOpts struct {
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}

func addSms(t twimlResponse, opts *SmsOpts, text *string) {
	t.appendContents(&sms{0, opts, text})
}
