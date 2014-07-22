package twiml

type number struct {
	XMLName int `xml:"Number"`
	*NumberOpts
	Num string `xml:",chardata"`
}

// NumberOpts allows setting behavior for after a call is established with a number.
//
// https://www.twilio.com/docs/api/twiml/number
type NumberOpts struct {
	SendDigits string `xml:"sendDigits,attr,omitempty"`
	URL        string `xml:"url,attr,omitempty"`
	Method     string `xml:"method,attr,omitempty"`
}

func addNumber(t twimlResponse, opts *NumberOpts, numbers []string) {
	for _, num := range numbers {
		t.appendContents(&number{NumberOpts: opts, Num: num})
	}
}
