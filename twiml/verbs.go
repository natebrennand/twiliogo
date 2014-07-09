package twiml

// Say some text during a phone call.
//
// https://www.twilio.com/docs/api/twiml/say
func (t *Response) Say(opts SayOpts, lines ...string) TwimlInterface {
	addSay(t, &opts, lines)
	return t
}

// Play an audio file during a phone call.
//
// https://www.twilio.com/docs/api/twiml/play
func (t *Response) Play(opts PlayOpts, urls ...string) TwimlInterface {
	addPlay(t, &opts, urls)
	return t
}

// Record audio during a phone call.
//
// https://www.twilio.com/docs/api/twiml/record
func (t *Response) Record(opts RecordOpts, action string) TwimlInterface {
	addRecord(t, &opts, &action)
	return t
}

// Wait for some number of seconds during a call
//
// https://www.twilio.com/docs/api/twiml/pause
func (t *Response) Pause(length int) TwimlInterface {
	addPause(t, length)
	return t
}

// Collect digits entered by a caller. Pass anoter TwimlInterface to use Say,
// Play, and Pause verbs during the Gather.
//
// https://www.twilio.com/docs/api/twiml/gather
func (t *Response) Gather(opts GatherOpts, nested GatherBody) TwimlInterface {
	newGather := &gather{0, &opts, nested}
	t.contents = append(t.contents, &newGather)

	return t
}

// Dial various things. Pass a DialBody interface to use Number, Queue,
// Conference, Sip, and Client nouns inside Dial.
// Play, and Pause verbs during the Gather.
//
// https://www.twilio.com/docs/api/twiml/gather
func (t *Response) Dial(opts DialOpts, nested DialBody) TwimlInterface {
	addDial(t, &opts, &nested)
	return t
}

// Send an Sms message during a phone call
//
// https://www.twilio.com/docs/api/twiml/sms
func (t *Response) Sms(opts SmsOpts, text string) TwimlInterface {
	addSms(t, &opts, &text)
	return t
}
