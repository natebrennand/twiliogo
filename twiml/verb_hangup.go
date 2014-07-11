package twiml

type hangup struct {
	XMLName int `xml:"Hangup"`
}

func addHangup(t twimlResponse) {
	t.appendContents(&hangup{})
}
