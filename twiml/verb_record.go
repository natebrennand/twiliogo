package twiml

type record struct {
	XMLName int `xml:"Record"`
	*RecordOpts
	Action *string `xml:"action,attr,omitempty"`
}

type RecordOpts struct {
	Method             string `xml:"method,attr,omitempty"`
	Timeout            int    `xml:"timeout,attr,omitempty"`
	FinishOnKey        string `xml:"finishOnKey,attr,omitempty"`
	MaxLength          int    `xml:"maxLength,attr,omitempty"`
	Transcribe         bool   `xml:"transcribe,attr,omitempty"`
	TranscribeCallback string `xml:"transcribeCallback,attr,omitempty"`
	PlayBeep           *bool  `xml:"playBeep,attr,omitempty"`
	Trim               string `xml:"trim,attr,omitempty"`
}

func addRecord(t twimlResponse, opts *RecordOpts, action *string) {
	t.appendContents(&record{RecordOpts: opts, Action: action})
}
