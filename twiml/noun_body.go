package twiml

type body struct {
	XMLName int    `xml:"Body"`
	Text    string `xml:",chardata"`
}

func addBody(t twimlResponse, lines []string) {
	for _, text := range lines {
		t.appendContents(&body{Text: text})
	}
}
