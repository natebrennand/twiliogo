package twiml

import "encoding/xml"

func marshalTwiml(t twimlResponse, e *xml.Encoder, start *xml.StartElement) error {
	return e.Encode(t.getContents())
}

func addPlay(t twimlResponse, opts PlayOpts, urls []string) {
	for _, url := range urls {
		t.appendContents(&play{0, url, &opts})
	}
}

func addSay(t twimlResponse, opts SayOpts, lines []string) {
	for _, line := range lines {
		t.appendContents(&say{0, line, &opts})
	}
}

func addPause(t twimlResponse, secs int) {
	t.appendContents(&pause{0, secs})
}
