package twiml

type pause struct {
	XMLName int `xml:"Pause"`
	Length  int `xml:"length,attr,omitempty"`
}

func addPause(t twimlResponse, secs int) {
	t.appendContents(&pause{0, secs})
}
