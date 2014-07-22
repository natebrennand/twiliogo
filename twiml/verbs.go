package twiml

// Say some text during a phone call.
//
// https://www.twilio.com/docs/api/twiml/say
func (t *Response) Say(opts SayOpts, lines ...string) twimlInterface {
	addSay(t, &opts, lines)
	return t
}

// Play an audio file during a phone call.
//
// https://www.twilio.com/docs/api/twiml/play
func (t *Response) Play(opts PlayOpts, urls ...string) twimlInterface {
	addPlay(t, &opts, urls)
	return t
}

// Record audio during a phone call.
//
// https://www.twilio.com/docs/api/twiml/record
func (t *Response) Record(opts RecordOpts, action string) twimlInterface {
	addRecord(t, &opts, &action)
	return t
}

// Pause for some number of seconds during a call
//
// https://www.twilio.com/docs/api/twiml/pause
func (t *Response) Pause(length int) twimlInterface {
	addPause(t, length)
	return t
}

// Gather collects digits entered by a caller. Pass a GatherBody interface to use Say,
// Play, and Pause verbs during the Gather.
//
// https://www.twilio.com/docs/api/twiml/gather
func (t *Response) Gather(opts GatherOpts, nested gatherBody) twimlInterface {
	addGather(t, &opts, &nested)
	return t
}

// Dial various things. Pass a DialBody interface to use Number, Queue,
// Conference, Sip, and Client nouns inside Dial.
// Play, and Pause verbs during the Gather.
//
// https://www.twilio.com/docs/api/twiml/gather
func (t *Response) Dial(opts DialOpts, nested dialBody) twimlInterface {
	addDial(t, &opts, &nested)
	return t
}

// Sms sends an Sms message during a phone call
//
// https://www.twilio.com/docs/api/twiml/sms
func (t *Response) Sms(opts SmsOpts, text string) twimlInterface {
	addSms(t, &opts, &text)
	return t
}

// Enqueue moves a caller into a queue.
//
// https://www.twilio.com/docs/api/twiml/enqueue
func (t *Response) Enqueue(opts EnqueueOpts, queue string) twimlInterface {
	addEnqueue(t, &opts, &queue)
	return t
}

// Leave removes a caller from a queue and return control to the previously executing
// TwiML.
//
// https://www.twilio.com/docs/api/twiml/leave
func (t *Response) Leave() twimlInterface {
	addLeave(t)
	return t
}

// Hangup during a phone call
//
// https://www.twilio.com/docs/api/twiml/hangup
func (t *Response) Hangup() twimlInterface {
	addHangup(t)
	return t
}

// Redirect TwiML flow to another page.
//
// https://www.twilio.com/docs/api/twiml/redirect
func (t *Response) Redirect(opts RedirectOpts, target string) twimlInterface {
	addRedirect(t, &opts, &target)
	return t
}

// Reject an incoming call.
//
// https://www.twilio.com/docs/api/twiml/reject
func (t *Response) Reject(reason string) twimlInterface {
	addReject(t, &reason)
	return t
}

// MessageMedia sends a message with text and / or media.
//
// https://www.twilio.com/docs/api/twiml/sms/message
func (t *Response) MessageMedia(opts MessageOpts, body messageBody) twimlInterface {
	addMessage(t, &opts, body)
	return t
}

// Message sends an sms message.
//
// https://www.twilio.com/docs/api/twiml/sms/message
func (t *Response) Message(opts MessageOpts, text ...string) twimlInterface {
	addMessageText(t, &opts, text)
	return t
}
