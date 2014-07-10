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

// Collect digits entered by a caller. Pass a GatherBody interface to use Say,
// Play, and Pause verbs during the Gather.
//
// https://www.twilio.com/docs/api/twiml/gather
func (t *Response) Gather(opts GatherOpts, nested GatherBody) TwimlInterface {
	addGather(t, &opts, &nested)
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

// Move a caller into a queue.
//
// https://www.twilio.com/docs/api/twiml/enqueue
func (t *Response) Enqueue(opts EnqueueOpts, queue string) TwimlInterface {
	addEnqueue(t, &opts, &queue)
	return t
}

// Remove a caller from a queue and return control to the previously executing
// TwiML.
//
// https://www.twilio.com/docs/api/twiml/leave
func (t *Response) Leave() TwimlInterface {
	addLeave(t)
	return t
}

// Hangup during a phone call
//
// https://www.twilio.com/docs/api/twiml/hangup
func (t *Response) Hangup() TwimlInterface {
	addHangup(t)
	return t
}

// Redirect TwiML flow to another page.
//
// https://www.twilio.com/docs/api/twiml/redirect
func (t *Response) Redirect(opts RedirectOpts, target string) TwimlInterface {
	addRedirect(t, &opts, &target)
	return t
}
