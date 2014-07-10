package twiml

type leave struct {
	XMLName int `xml:"Leave"`
}

func addLeave(t twimlResponse) {
	t.appendContents(&leave{})
}
