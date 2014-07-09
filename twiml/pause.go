package twiml

type pause struct {
	XMLName int `xml:"Pause"`
	Length  int `xml:"length,attr,omitempty"`
}
