package twiml

type play struct {
	XMLName int `xml:"Play"`
	*PlayOpts
	Url string `xml:",chardata"`
}

type PlayOpts struct {
	Loop   int `xml:"loop,attr,omitempty"`
	Digits int `xml:"digits,attr,omitempty"`
}

func addPlay(t twimlResponse, opts *PlayOpts, urls []string) {
	for _, url := range urls {
		t.appendContents(&play{0, opts, url})
	}
}
