package notifications

var (
	testNotificationString = `{
	"sid": "NO5a7a84730f529f0a76b3e30c01315d1a",
	"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
	"call_sid": "CAa8857b0dcc71b4909aced594f7f87453",
	"log": "0",
	"error_code": "11205",
	"more_info": "http:\/\/www.twilio.com\/docs\/errors\/11205",
	"message_text": "EmailNotification=false&LogLevel=ERROR&sourceComponent=13400&Msg=HTTP+Connection+Failure+-+Read+timed+out&ErrorCode=11205&msg=HTTP+Connection+Failure+-+Read+timed+out&url=4min19secs.mp3",
	"message_date": "Tue, 09 Feb 2010 01:23:53 +0000",
	"response_body": "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<Response>\n\t<Play>4min19secs.mp3<\/Play>\n<\/Response>\n",
	"request_method": "GET",
	"request_url": "http:\/\/demo.twilio.com\/welcome",
	"request_variables": "AccountSid=AC381707b751dbe4c74b15c5697ba67afd&CallStatus=in-progress&Called=4152374451&CallerCountry=US&CalledZip=94937&CallerCity=&Caller=4150000000&CalledCity=INVERNESS&CalledCountry=US&DialStatus=answered&CallerState=California&CallSid=CAa8857b0dcc71b4909aced594f7f87453&CalledState=CA&CallerZip=",
	"response_headers": "Date=Tue%2C+09+Feb+2010+01%3A23%3A38+GMT&Vary=Accept-Encoding&Content-Length=91&Content-Type=text%2Fxml&Accept-Ranges=bytes&Server=Apache%2F2.2.3+%28CentOS%29&X-Powered-By=PHP%2F5.1.6",
	"date_created": "Tue, 09 Feb 2010 01:23:53 +0000",
	"api_version": "2008-08-01",
	"date_updated": "Tue, 09 Feb 2010 01:23:53 +0000",
	"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications\/NO5a7a84730f529f0a76b3e30c01315d1a.json"
}`

	testNotificatinListString = `{
	"page": 0,
	"num_pages": 25,
	"page_size": 50,
	"total": 1224,
	"start": 0,
	"end": 49,
	"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications.json",
	"first_page_uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications.json?Page=0&PageSize=50",
	"previous_page_uri": null,
	"next_page_uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications.json?Page=1&PageSize=50",
	"last_page_uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications.json?Page=24&PageSize=50",
	"notifications": [
		{
			"sid": "NO5a7a84730f529f0a76b3e30c01315d1a",
			"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
			"call_sid": "CAa8857b0dcc71b4909aced594f7f87453",
			"log": "0",
			"error_code": "11205",
			"more_info": "http:\/\/www.twilio.com\/docs\/errors\/11205",
			"message_text": "EmailNotification=false&LogLevel=ERROR&sourceComponent=13400&Msg=HTTP+Connection+Failure+-+Read+timed+out&ErrorCode=11205&msg=HTTP+Connection+Failure+-+Read+timed+out&url=4min19secs.mp3",
			"message_date": "Tue, 09 Feb 2010 01:23:53 +0000",
			"request_method": "POST",
			"request_url": "http:\/\/demo.twilio.com\/welcome",
			"date_created": "Tue, 09 Feb 2010 01:23:53 +0000",
			"api_version": "2008-08-01",
			"date_updated": "Tue, 09 Feb 2010 01:23:53 +0000",
			"uri": "\/2010-04-01\/Accounts\/AC381707b751dbe4c74b15c5697ba67afd\/Notifications\/NO5a7a84730f529f0a76b3e30c01315d1a.json"
		}
	]
}`
)
