package twiml

type number struct {
	XMLName int `xml:"Number"`
	*NumberOpts
	Num string `xml:",chardata"`
}

type NumberOpts struct {
	SendDigits string `xml:"sendDigits,attr,omitempty"`
	Url        string `xml:"url,attr,omitempty"`
	Method     string `xml:"method,attr,omitempty"`
}

func addNumber(t twimlResponse, opts *NumberOpts, numbers []string) {
	for _, num := range numbers {
		t.appendContents(&number{0, opts, num})
	}
}
