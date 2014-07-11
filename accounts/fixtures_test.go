package account

var (
	testAccount = `{
	"sid": "AC381707b751dbe4c74b15c5697ba67afd",
	"friendly_name": "Do you like my friendly name?",
	"type": "Full",
	"status": "active",
	"date_created": "Wed, 04 Aug 2010 21:37:41 +0000",
	"date_updated": "Fri, 06 Aug 2010 01:15:02 +0000",
	"auth_token": "redacted",
	"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd.json",
	"subresource_uris": {
		"available_phone_numbers": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/AvailablePhoneNumbers.json",
		"calls": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Calls.json",
		"conferences": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Conferences.json",
		"incoming_phone_numbers": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/IncomingPhoneNumbers.json",
		"notifications": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications.json",
		"outgoing_caller_ids": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/OutgoingCallerIds.json",
		"recordings": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Recordings.json",
		"sandbox": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Sandbox.json",
		"sms_messages": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/SMS\/Messages.json",
		"transcriptions": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Transcriptions.json"
	}
}`
	testListAccount = `{
	"page": 0,
	"num_pages": 1,
	"page_size": 50,
	"total": 1,
	"start": 0,
	"end": 1,
	"uri": "\/2010-04-01\/Accounts.json",
	"first_page_uri": "\/2010-04-01\/Accounts.json?Page=0&PageSize=50",
	"previous_page_uri": null,
	"next_page_uri": null,
	"last_page_uri": "\/2010-04-01\/Accounts.json?Page=0&PageSize=50",
	"accounts": [{
		"sid":"AC381707b751dbe4c74b15c5697ba67afd",
		"friendly_name": "Chieftain",
		"status": "active",
		"auth_token": "redacted",
		"date_created": "Tue, 12 Jan 2010 04:41:09 +0000",
		"date_updated": "Tue, 25 Jan 2011 07:24:36 +0000",
		"type": "Full",
		"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd.json",
		"subresource_uris": {
			"available_phone_numbers": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/AvailablePhoneNumbers.json",
			"calls": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Calls.json",
			"conferences": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Conferences.json",
			"incoming_phone_numbers": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/IncomingPhoneNumbers.json",
			"notifications": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications.json",
			"outgoing_caller_ids": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/OutgoingCallerIds.json",
			"recordings": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Recordings.json",
			"sandbox": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Sandbox.json",
			"sms_messages": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/SMS\/Messages.json",
			"transcriptions": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Transcriptions.json"}
		}
	]
}`
)
