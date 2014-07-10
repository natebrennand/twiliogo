package twiml

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const endToEndStr = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say voice="alice">My hands are typing words</Say>
	<Say voice="alice">Haaaaaaaaaaaaands</Say>
</Response>
`

var testTwiml *Response

func TestTwimlSatisfiesXmlInterface(t *testing.T) {
	assert.Implements(t, (*xml.Marshaler)(nil), new(Response))
}

// func TestGatherSatisfiesXmlInterface(t *testing.T) {
// 	assert.Implements(t, (*xml.Marshaler)(nil), new(gather))
// }

func TestEmptyResponse(t *testing.T) {
	testTwiml = &Response{}
	output, err := testTwiml.Render()
	str := string(output)
	assert.NoError(t, err)
	assert.Contains(t, str, xml.Header)
	assert.Contains(t, str, "<Response>")
	assert.Contains(t, str, "</Response>")
	t.Log(string(str))
}

func TestEndToEnd(t *testing.T) {
	testTwiml = new(Response)
	output, err := testTwiml.Say(SayOpts{Voice: "alice"}, "My hands are typing words", "Haaaaaaaaaaaaands").Render()
	assert.NoError(t, err)
	expected := strings.TrimSpace(strings.Replace(endToEndStr, "\n", "", -1))
	actual := strings.TrimSpace(strings.Replace(string(output), "\n", "", -1))
	assert.Exactly(t, expected, actual)
}

func TestSay(t *testing.T) {
	testTwiml = &Response{}
	opts := SayOpts{Voice: "alice", Language: "english", Loop: 2}
	_, err := testTwiml.Say(opts, "line two").Render()
	assert.NoError(t, err)
}

func TestPlay(t *testing.T) {
	testTwiml = &Response{}
	_, err := testTwiml.Play(PlayOpts{}, "http://demo.kevinwhinnery.com/audio/zelda.mp3", "http://somesite.com/leroyjenkins.mp3").Render()
	assert.NoError(t, err)
}

func TestRecord(t *testing.T) {
	testTwiml = &Response{}
	recOpts := RecordOpts{Method: "POST"}
	_, err := testTwiml.Record(recOpts, "http://demo.kevinwhinnery.com/audio/zelda.mp3").Render()
	assert.NoError(t, err)
}

func TestPause(t *testing.T) {
	testTwiml = new(Response)
	_, err := testTwiml.Pause(5).Render()
	assert.NoError(t, err)
}

func TestSms(t *testing.T) {
	testTwiml = new(Response)
	_, err := testTwiml.Sms(SmsOpts{}, "Welcome to catfacts...").Render()
	assert.NoError(t, err)
}

func TestGather(t *testing.T) {
	testTwiml = &Response{}
	testTwiml.Gather(GatherOpts{Timeout: 10}, new(GatherTwiml).Play(PlayOpts{}, "stuff"))
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	// Make sure the gather struct doesn't render it's Inner field to xml.
	assert.NotContains(t, str, "Inner")
}
