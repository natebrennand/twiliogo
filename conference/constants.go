package conference

import ()

const (
	getURL          = "https://api.twilio.com/2010-04-01/Accounts/%s/Conferences/%s.json"                 // takes account sid, conference sid
	participantsURL = "https://api.twilio.com/2010-04-01/Accounts/%s/Conferences/%s/Participants.json"    // takes account sid, conference sid
	listURL         = "https://api.twilio.com/2010-04-01/Accounts/%s/Conferences.json"                    // takes account sid
	participantURL  = "https://api.twilio.com/2010-04-01/Accounts/%s/Conferences/%s/Participants/%s.json" // takes account sid, conference sid, callsid

)
