package twiml

type sms struct {
	XMLName int `xml:"Sms"`
	*SmsOpts
	Text *string `xml:",chardata"`
}

// SmsOpts sets the attributes of the SMS message.
//
// https://www.twilio.com/docs/api/twiml/sms
type SmsOpts struct {
	To             string `xml:"to,attr,omitempty"`
	From           string `xml:"from,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	StatusCallback string `xml:"statusCallback,attr,omitempty"`
}

func addSms(t twimlResponse, opts *SmsOpts, text *string) {
	t.appendContents(&sms{SmsOpts: opts, Text: text})
}
