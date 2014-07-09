package twiml

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTwiml TwimlInterface

func TestEmptyResponse(t *testing.T) {
	testTwiml = &Twiml{}
	output, err := testTwiml.Render()
	str := string(output)
	assert.NoError(t, err)
	assert.Contains(t, str, xml.Header)
	assert.Contains(t, str, "<Response>")
	assert.Contains(t, str, "</Response>")
	t.Log(string(str))
}

func TestSay(t *testing.T) {
	testTwiml = &Twiml{}
	opts := SayOpts{Voice: "alice", Language: "english", Loop: 2}
	_, err := testTwiml.Say(opts, "line two").Render()
	assert.NoError(t, err)
}

func TestPlay(t *testing.T) {
	testTwiml = &Twiml{}
	_, err := testTwiml.Play(PlayOpts{}, "http://demo.kevinwhinnery.com/audio/zelda.mp3", "http://somesite.com/leroyjenkins.mp3").Render()
	assert.NoError(t, err)
}

func TestRecord(t *testing.T) {
	testTwiml = &Twiml{}
	recOpts := RecordOpts{Method: "POST"}
	_, err := testTwiml.Record(recOpts, "http://demo.kevinwhinnery.com/audio/zelda.mp3").Render()
	assert.NoError(t, err)
}
