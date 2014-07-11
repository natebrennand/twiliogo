package twiml_test

import (
	"fmt"
	"net/http"

	"github.com/natebrennand/twiliogo/twiml"
)

func Example() {
	response := new(twiml.Response)
	response.Say(twiml.SayOpts{Voice: "alice"}, "My hands are typing words", "Haaaaaaaaaaaaands")
	output, _ := response.Render()
	fmt.Println(string(output))
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	//
	// <Response>
	// 	<Say voice="alice">My hands are typing words</Say>
	// 	<Say voice="alice">Haaaaaaaaaaaaands</Say>
	// </Response>
}

// twiml.Response implements the http.Handler interface, so you can pass
// your prepared response directly to http to handle a route.
func ExampleResponse() {
	response := new(twiml.Response)
	response.Say(twiml.SayOpts{Voice: "alice"}, "Hi there!", "I'm responding to an HTTP request.")
	http.Handle("/callback/voice/", response)
	http.ListenAndServe(":8080", nil)
}
