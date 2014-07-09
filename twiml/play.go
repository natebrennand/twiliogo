package twiml

type play struct {
	XMLName int    `xml:"Play"`
	Url     string `xml:",chardata"`
	*PlayOpts
}

type PlayOpts struct {
	Loop   int `xml:"loop,attr,omitempty"`
	Digits int `xml:"digits,attr,omitempty"`
}
