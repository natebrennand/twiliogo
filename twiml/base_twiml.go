package twiml

import "encoding/xml"

type baseTwiml struct {
	contents []interface{}
}

func (t *baseTwiml) getContents() []interface{} {
	return t.contents
}

func (t *baseTwiml) appendContents(v interface{}) {
	t.contents = append(t.contents, v)
}

func (t *baseTwiml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalTwiml(t, e, &start)
}
