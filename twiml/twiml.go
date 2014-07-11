package twiml

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

var (
	preTwiml  = []byte(xml.Header + "\n<Response>\n")
	postTwiml = []byte("\n</Response>\n")
)

type twimlResponse interface {
	xml.Marshaler
	getContents() []interface{}
	appendContents(interface{})
}

type TwimlInterface interface {
	xml.Marshaler
	Render() ([]byte, error)
	RenderReader() (io.Reader, error)
	Say(SayOpts, ...string) TwimlInterface
	Play(PlayOpts, ...string) TwimlInterface
	Dial(DialOpts, DialBody) TwimlInterface
	Record(RecordOpts, string) TwimlInterface
	Gather(GatherOpts, GatherBody) TwimlInterface
	Sms(SmsOpts, string) TwimlInterface
	Enqueue(EnqueueOpts, string) TwimlInterface
	Leave() TwimlInterface
	Hangup() TwimlInterface
	Redirect(RedirectOpts, string) TwimlInterface
	Pause(int) TwimlInterface
	Reject(string) TwimlInterface
	Message(MessageOpts, ...string) TwimlInterface
	MessageMedia(MessageOpts, MessageBody) TwimlInterface
}

type Response struct {
	baseTwiml
	cache struct {
		xml   []byte
		valid bool
	}
}

func (t *Response) appendContents(v interface{}) {
	t.cache.valid = false
	t.contents = append(t.contents, v)
}

// Forcibly clear the internal cache. See Render for an explanation.
func (t *Response) ClearCache() {
	t.cache.valid = false
}

// Returns a TwiML representation of the previous calls on the struct as a byte
// slice.
// TODO: the caching implemented here will fail if someone retroactively updates
// one of the nested structs in Dial, Gather, or Message. This should be fixed
// by either finding a way to update the cache flag when that happens,
// recursivley caching, or forcing the user to pass whole objects instead of
// pointers so they loose the reference. For now any user who wants to do this
// must use ClearCache.
func (t *Response) Render() (result []byte, err error) {
	if !t.cache.valid {
		result, err = xml.MarshalIndent(t.contents, "\t", "\t")
		if err != nil {
			return
		}
		result = append(preTwiml, result...)
		result = append(result, postTwiml...)
		t.cache.xml = make([]byte, len(result))
		copy(t.cache.xml, result)
		t.cache.valid = true
		return
	}
	result = make([]byte, len(t.cache.xml))
	copy(result, t.cache.xml)
	return
}

// Returns a TwiML representation of the previous calls on the struct, enclosed
// in a Reader interface.
func (t *Response) RenderReader() (io.Reader, error) {
	result, err := t.Render()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(result), nil
}

// Returns an implements http.Handler so a Response object can be passed
// directly to http.Handle
func (t *Response) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	output, err := t.Render()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Write(output)
}
