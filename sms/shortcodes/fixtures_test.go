package shortcodes

var (
	testRecordString = `{
	"sid": "SC6b20cb705c1e8f00210049b20b70fce2",
	"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
	"friendly_name": "67898",
	"short_code": "67898",
	"date_created": "Sat, 09 Jul 2011 22:36:22 +0000",
	"date_updated": "Sat, 09 Jul 2011 22:36:22 +0000",
	"sms_url": "http://demo.twilio.com/docs/sms.xml",
	"sms_method": "POST",
	"sms_fallback_url": "http://example.com/fallback",
	"sms_fallback_method": "GET",
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/SMS/ShortCodes/SC6b20cb705c1e8f00210049b20b70fce2.json"
}`
	testRecordListString = `{
	"page": 0,
	"num_pages": 1,
	"page_size": 50,
	"total": 10,
	"start": 0,
	"end": 9,
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/SMS/ShortCodes.json",
	"first_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/SMS/ShortCodes.json?Page=0&PageSize=50",
	"previous_page_uri": null,
	"next_page_uri": null,
	"last_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/SMS/ShortCodes.json?Page=0&PageSize=50",
	"short_codes": [
		{
			"sid": "SC6b20cb705c1e8f00210049b20b70fce2",
			"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
			"friendly_name": "67898",
			"short_code": "67898",
			"date_created": "Sat, 09 Jul 2011 22:36:22 +0000",
			"date_updated": "Wed, 13 Jul 2011 02:07:05 +0000",
			"sms_url": "http://demo.twilio.com/docs/sms.xml",
			"sms_method": "POST",
			"sms_fallback_url": "http://smsapp.com/fallback",
			"sms_fallback_method": "GET",
			"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/SMS/ShortCodes/SC6b20cb705c1e8f00210049b20b70fce2.json"
		}
	]
}`
)
