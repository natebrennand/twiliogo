package twiml

type sms struct {
	XMLName int    `xml:"Sms"`
	Text    string `xml:",chardata"`
	*SmsOpts
}

type SmsOpts struct {
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}
