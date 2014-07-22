package numbers

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var testResponseFixture = Number{
	Sid:         "PN2a0747eba6abf96b7e3c3ff0b4530f6e",
	DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
	DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
	PhoneNumber: "+15105647903",
	Capabilities: capabilities{
		true,
		true,
		false,
	},
	URI:        "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/IncomingPhoneNumbers/PN2a0747eba6abf96b7e3c3ff0b4530f6e.json",
	AccountSid: "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
	core: core{
		FriendlyName:        "My Company Line",
		APIVersion:          "2010-04-01",
		VoiceURL:            "http://demo.twilio.com/docs/voice.xml",
		VoiceMethod:         "POST",
		VoiceFallbackMethod: "POST",
		SmsURL:              "http://demo.twilio.com/docs/sms.xml",
		SmsMethod:           "POST",
		SmsFallbackMethod:   "GET",
	},
}

var testResponseFixtureString = `{
    "sid": "PN2a0747eba6abf96b7e3c3ff0b4530f6e",
    "account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
    "friendly_name": "My Company Line",
    "phone_number": "+15105647903",
    "voice_url": "http://demo.twilio.com/docs/voice.xml",
    "voice_method": "POST",
    "voice_fallback_url": null,
    "voice_fallback_method": "POST",
    "voice_caller_id_lookup": null,
    "voice_application_sid": null,
    "date_created": "Mon, 16 Aug 2010 23:00:23 +0000",
    "date_updated": "Mon, 16 Aug 2010 23:00:23 +0000",
    "sms_url": "http://demo.twilio.com/docs/sms.xml",
    "sms_method": "POST",
    "sms_fallback_url": null,
    "sms_fallback_method": "GET",
    "sms_application_sid": null,
    "capabilities": {
        "voice": true,
        "sms": true,
        "mms": false
    },
    "status_callback": null,
    "status_callback_method": null,
    "api_version": "2010-04-01",
    "uri": "\/2010-04-01\/Accounts\/AC5116d5d4df9f61ceae2f0732e1ea9f1b\/IncomingPhoneNumbers\/PN2a0747eba6abf96b7e3c3ff0b4530f6e.json"
}`

var testListFixture = NumberList{
	ListResponseCore: common.ListResponseCore{
		Start:           0,
		Total:           261,
		NumPages:        6,
		Page:            0,
		PageSize:        50,
		End:             49,
		URI:             "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json",
		FirstPageURI:    "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json?Page=0&PageSize=50",
		LastPageURI:     "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json?Page=5&PageSize=50",
		NextPageURI:     "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json?Page=1&PageSize=50",
		PreviousPageURI: "",
	},
	IncomingPhoneNumbers: &[]Number{
		Number{
			Sid:         "PN2a0747eba6abf96b7e3c3ff0b4530f6e",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			PhoneNumber: "+15105647903",
			Capabilities: capabilities{
				Voice: true,
				SMS:   true,
				MMS:   false,
			},
			URI: "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/IncomingPhoneNumbers/PN2a0747eba6abf96b7e3c3ff0b4530f6e.json",
			core: core{
				FriendlyName:        "My Company Line",
				APIVersion:          "2010-04-01",
				VoiceURL:            "http://demo.twilio.com/docs/voice.xml",
				VoiceMethod:         "POST",
				VoiceFallbackMethod: "POST",
				SmsURL:              "http://demo.twilio.com/docs/sms.xml",
				SmsMethod:           "POST",
				SmsFallbackMethod:   "GET",
			},
			AccountSid: "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
		},
	},
}

var testListFixtureString = `{
	"start": 0,
	"total": 261,
	"num_pages": 6,
	"page": 0,
	"page_size": 50,
	"end": 49,
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Calls.json",
	"first_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json?Page=0&PageSize=50",
	"last_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json?Page=5&PageSize=50",
	"next_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls.json?Page=1&PageSize=50",
	"previous_page_uri": null,
	"recordings": [
		{
			"account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
			"api_version": "2010-04-01",
			"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
			"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
			"direction": "outbound-api",
			"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Calls/CAd88cd5b804dbcfb0ae6e0ccbbca67b83.json",
			"sid": "REc8009ae243230394bc84437e1a9f4650",
			"call_sid": "CAd88cd5b804dbcfb0ae6e0ccbbca67b83",
			"duration": "45"
		}
	]
}`
