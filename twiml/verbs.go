package twiml

type say struct {
	XMLName  int    `xml:"Say"`
	Sentence string `xml:",chardata"`
	*SayOpts
}

type SayOpts struct {
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
}

type play struct {
	XMLName int    `xml:"Play"`
	Url     string `xml:",chardata"`
	*PlayOpts
}

type PlayOpts struct {
	Loop   int `xml:"loop,attr,omitempty"`
	Digits int `xml:"digits,attr,omitempty"`
}

type record struct {
	XMLName int    `xml:"Record"`
	Action  string `xml:"action,attr,omitempty"`
	*RecordOpts
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
