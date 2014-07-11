package twiml

type media struct {
	XMLName int    `xml:"Media"`
	Source  string `xml:",chardata"`
}

func addMedia(t twimlResponse, sources []string) {
	for _, source := range sources {
		t.appendContents(&media{Source: source})
	}
}
