package twiml_test

import (
	"fmt"

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
