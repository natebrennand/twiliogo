package twiml

type reject struct {
	XMLName int     `xml:"Reject"`
	Reason  *string `xml:"reason,attr,omitempty"`
}

func addReject(t twimlResponse, reason *string) {
	t.appendContents(&reject{Reason: reason})
}
