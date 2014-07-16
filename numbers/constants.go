package numbers

const (
	getURL      = "https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers/%s.json"       // takes act sid and incoming phone # sid
	listURL     = "https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers.json"          // takes act sid
	localURL    = "https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers/Local.json"    // takes act sid
	tollFreeURL = "https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers/TollFree.json" // takes act sid
	mobileURL   = "https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers/Mobile.json"   // takes act sid
)
