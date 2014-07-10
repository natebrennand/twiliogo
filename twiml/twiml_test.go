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

func TestMessageTwimlSatisfiesMessageBody(t *testing.T) {
	assert.Implements(t, (*MessageBody)(nil), new(MessageTwiml))
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
	// Make sure the gather struct doesn't render it's Body field to xml.
	assert.NotContains(t, str, "Body")
}

func TestDial(t *testing.T) {
	testTwiml = &Response{}
	innerDial := new(DialTwiml)
	innerDial.
		Number(NumberOpts{}, "0123456789", "9876543210").
		Sip(SipOpts{}, "sip:address").
		Queue(QueueOpts{}, "myQueue").
		Client(ClientOpts{}, "clientName").
		Conference(ConferenceOpts{}, "myConference")

	testTwiml.Dial(DialOpts{Timeout: 30}, innerDial)
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Dial")
	assert.Contains(t, str, "Number")
	assert.Contains(t, str, "Queue")
	assert.Contains(t, str, "Client")
	assert.Contains(t, str, "Conference")
}

func TestEnqueue(t *testing.T) {
	testTwiml = &Response{}
	testTwiml.Enqueue(EnqueueOpts{Method: "POST"}, "infiniteHold")
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Enqueue")
	assert.Contains(t, str, `method="POST"`)
	assert.Contains(t, str, "infiniteHold")
}

func TestLeave(t *testing.T) {
	testTwiml = &Response{}
	testTwiml.Leave()
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Leave")
}

func TestHangup(t *testing.T) {
	testTwiml = &Response{}
	testTwiml.Hangup()
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Hangup")
}

func TestRedirect(t *testing.T) {
	testTwiml = &Response{}
	testTwiml.Redirect(RedirectOpts{Method: "POST"}, "http://mysite/otherTwiml")
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Redirect")
	assert.Contains(t, str, "POST")
}

func TestReject(t *testing.T) {
	testTwiml = &Response{}
	testTwiml.Reject("busy")
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Reject")
	assert.Contains(t, str, `reason="busy"`)
}

func TestMessageMedia(t *testing.T) {
	testTwiml = new(Response)
	innerTwiml := new(MessageTwiml)
	innerTwiml.Body("Welcome to owl facts").Media("https://demo.twilio.com/owl.png")
	testTwiml.MessageMedia(MessageOpts{Method: "POST"}, innerTwiml)
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Message")
	assert.Contains(t, str, `method="POST"`)
	assert.Contains(t, str, "Media")
	assert.Contains(t, str, "Body")
}

func TestMessage(t *testing.T) {
	testTwiml = new(Response)
	testTwiml.Message(MessageOpts{Method: "POST"}, "Welcome to owl facts", `Text "hoot" to unsibscribe`)
	output, err := testTwiml.Render()
	assert.NoError(t, err)
	str := string(output)
	assert.Contains(t, str, "Message")
	assert.Contains(t, str, `method="POST"`)
	assert.Contains(t, str, "Body")
	t.Log(str)
}
