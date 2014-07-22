package twiml

type play struct {
	XMLName int `xml:"Play"`
	*PlayOpts
	URL string `xml:",chardata"`
}

// PlayOpts sets the configuration for playing audio files.
type PlayOpts struct {
	Loop   int `xml:"loop,attr,omitempty"`
	Digits int `xml:"digits,attr,omitempty"`
}

func addPlay(t twimlResponse, opts *PlayOpts, urls []string) {
	for _, url := range urls {
		t.appendContents(&play{PlayOpts: opts, URL: url})
	}
}
