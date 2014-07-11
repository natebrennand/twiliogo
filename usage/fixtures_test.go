package usage

var (
	testTriggerString = `{
	"usage_record_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Records.json?Category=calls",
	"date_updated": "Sat, 29 Sep 2012 19:47:54 +0000",
	"date_fired": null,
	"friendly_name": "Trigger for calls at usage of 500",
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Triggers/UT33c6aeeba34e48f38d6899ea5b765ad4.json",
	"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
	"callback_method": "POST",
	"trigger_by": "usage",
	"sid": "UT33c6aeeba34e48f38d6899ea5b765ad4",
	"current_value": "21",
	"date_created": "Sat, 29 Sep 2012 19:45:43 +0000",
	"callback_url": "http://www.example.com/",
	"recurring": null,
	"usage_category": "calls",
	"trigger_value": "500.000000"
}`
	testTriggerListString = `{
	"first_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Triggers/.json?UsageCategory=calls&Recurring=daily&Page=0&PageSize=50",
	"end": 0,
	"previous_page_uri": null,
	"usage_triggers": [
		{
			"usage_record_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Records/Today.json?Category=calls",
			"date_updated": "Sat, 29 Sep 2012 19:42:57 +0000",
			"date_fired": null,
			"friendly_name": "a trigger",
			"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Triggers//UTc2db285b0cbf4c60a2f1a8db237a5fba.json",
			"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
			"callback_method": "POST",
			"trigger_by": "count",
			"sid": "UTc2db285b0cbf4c60a2f1a8db237a5fba",
			"current_value": "0",
			"date_created": "Sun, 23 Sep 2012 23:07:29 +0000",
			"callback_url": "http://www.google.com",
			"recurring": "daily",
			"usage_category": "calls",
			"trigger_value": "0.000000"
		}
	],
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Triggers.json?UsageCategory=calls&Recurring=daily",
	"page_size": 50,
	"start": 0,
	"next_page_uri": null,
	"num_pages": 1,
	"total": 1,
	"last_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Usage/Triggers.json?UsageCategory=calls&Recurring=daily&Page=0&PageSize=50",
	"page": 0
}`
)
