package applications

var (
	testApplicationString = `{
	"sid": "AP2a0747eba6abf96b7e3c3ff0b4530f6e",
	"date_created": "Mon, 16 Aug 2010 23:00:23 +0000",
	"date_updated": "Mon, 16 Aug 2010 23:00:23 +0000",
	"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
	"friendly_name": "Phone Me",
	"api_version": "2010-04-01",
	"voice_url": "http://demo.twilio.com/docs/voice.xml",
	"voice_method": "POST",
	"voice_fallback_url": null,
	"voice_fallback_method": "POST",
	"status_callback": null,
	"status_callback_method": null,
	"voice_caller_id_lookup": null,
	"sms_url": null,
	"sms_method": "POST",
	"sms_fallback_url": null,
	"sms_fallback_method": "GET",
	"sms_status_callback": null,
	"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Applications\/AP2a0747eba6abf96b7e3c3ff0b4530f6e.json"
}`
	testListApplicationString = `{
	"page": 0,
	"num_pages": 1,
	"page_size": 50,
	"total": 6,
	"start": 0,
	"end": 5,
	"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Applications.json",
	"first_page_uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Applications.json?Page=0&PageSize=50",
	"previous_page_uri": null,
	"next_page_uri": null,
	"last_page_uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Applications.json?Page=0&PageSize=50",
	"applications": [
		{
			"sid": "AP3f94c94562ac88dccf16f8859a1a8b25",
			"date_created": "Thu, 13 Nov 2008 07:56:24 +0000",
			"date_updated": "Thu, 13 Nov 2008 08:45:58 +0000",
			"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
			"friendly_name": "Long Play",
			"api_version": "2010-04-01",
			"voice_url": "http:\/\/demo.twilio.com\/docs/voice.xml",
			"voice_method": "GET",
			"voice_fallback_url": null,
			"voice_fallback_method": null,
			"status_callback": null,
			"status_callback_method": null,
			"voice_caller_id_lookup": null,
			"sms_url": null,
			"sms_method": null,
			"sms_fallback_url": null,
			"sms_fallback_method": null,
			"sms_status_callback": null,
			"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Applications\/AP3f94c94562ac88dccf16f8859a1a8b25.json"
		}
	]
}`
)
