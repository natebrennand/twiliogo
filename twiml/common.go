package twiml

import "encoding/xml"

func marshalTwiml(t twimlResponse, e *xml.Encoder, start *xml.StartElement) error {
	return e.Encode(t.getContents())
}
