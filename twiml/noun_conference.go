package twiml

type conference struct {
	XMLName int `xml:"Conference"`
	*ConferenceOpts
	Name *string `xml:",chardata"`
}

type ConferenceOpts struct {
	Muted                  bool   `xml:"muted,attr,omitempty"`
	Beep                   string `xml:"beep,attr,omitempty"`
	StartConferenceOnEnter *bool  `xml:"startConferenceOnEnter,attr,omitempty"`
	EndConferenceOnExit    bool   `xml:"endConferenceOnExit,attr,omitempty"`
	WaitUrl                string `xml:"waitUrl,attr,omitempty"`
	WaitMethod             string `xml:"waitMethod,attr,omitempty"`
	MaxParticipants        int    `xml:"maxParticipants,attr,omitempty"`
	Record                 string `xml:"record,attr,omitempty"`
	Trim                   string `xml:"trim,attr,omitempty"`
	EventCallbackUrl       string `xml:"eventCallbackUrl,attr,omitempty"`
}

func addConference(t twimlResponse, opts *ConferenceOpts, name *string) {
	t.appendContents(&conference{ConferenceOpts: opts, Name: name})
}
