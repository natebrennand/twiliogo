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
