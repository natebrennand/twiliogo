package voice

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var (
	testNumber1                   = "+15558675309"
	testNumber2                   = "+14158141829"
	testResponseFixtureAccountSid = "ACd03387e99bf959daa1e4810cc945708d"
)

var testPostFixtureString = `{
		"url":"http://twimlbin.com/558a498f",
		"to":"+15558675309",
		"from":"+14158141829",
	}`

var testPostFixture = Post{
	URL:  "http://twimlbin.com/558a498f",
	To:   "+15558675309",
	From: "+14158141829",
}

var testResponseFixtureString = `{
	"sid": "CA7383500ec70ce66bd3a7ac5d2fbbd6a9",
	"date_created": null,
	"date_updated": null,
	"parent_call_sid": null,
	"account_sid": "ACd03387e99bf959daa1e4810cc945708d",
	"to": "+16164601267",
	"to_formatted": "(616) 460-1267",
	"from": "+13139202596",
	"from_formatted": "(313) 920-2596",
	"phone_number_sid": "PN722b65e5e68bbd7f428f23708c4f47ee",
	"status": "queued",
	"start_time": null,
	"end_time": null,
	"duration": null,
	"price": null,
	"price_unit": "USD",
	"direction": "outbound-api",
	"answered_by": null,
	"api_version": "2010-04-01",
	"annotation": null,
	"forwarded_from": null,
	"group_sid": null,
	"caller_name": null,
	"uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Calls/CA7383500ec70ce66bd3a7ac5d2fbbd6a9.json",
	"subresource_uris": {
		"notifications": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Calls/CA7383500ec70ce66bd3a7ac5d2fbbd6a9/Notifications.json",
		"recordings": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Calls/CA7383500ec70ce66bd3a7ac5d2fbbd6a9/Recordings.json"
	}
}`

var testResponseFixture = Call{
	ResponseCore: common.ResponseCore{
		AccountSid:   "ACd03387e99bf959daa1e4810cc945708d",
		APIVersion:   "2010-04-01",
		Direction:    "outbound-api",
		ErrorCode:    "",
		ErrorMessage: "",
		From:         "+13139202596",
		To:           "+16164601267",
		URI:          "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Calls/CA7383500ec70ce66bd3a7ac5d2fbbd6a9.json",
		DateCreated:  common.JSONTime{time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
		DateUpdated:  common.JSONTime{time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
	},

	Price: 0.0,
}

var testCallbackFixture = Callback{
	CallDuration:      4,
	RecordingURL:      "",
	RecordingSid:      "recordingsid",
	RecordingDuration: 60,
	CallSid:           "callsid",
	CallStatus:        "completed",
	APIVersion:        "2010-04-01",
	Direction:         "outbound-api",
	ForwardedFrom:     "NianCat",
	CallerName:        "OwlMonkey",
	StandardRequest: common.StandardRequest{
		AccountSid: "ACd03387e99bf959daa1e4810cc945708d",
		From:       "+13139202596",
		To:         "+16164601267",
		Location:   nil,
	},
}

var testCallbackFixtureFormString = `CallDuration=4&RecordingSid=recordingsid&RecordingDuration=60&CallSid=callsid&CallStatus=completed&APIVersion=2010-04-01&Direction=outbound-api&ForwardedFrom=NianCat&CallerName=OwlMonkey&AccountSid=ACd03387e99bf959daa1e4810cc945708d&From=+13139202596&To=+16164601267`
var testCallbackFixtureFormFailureString = `CallDuration=4&RecordingURL=4&RecordingSid=recordingsid&RecordingDuration=60&CallSid=callsid&CallStatus=completed&APIVersion=2010-04-01&Direction=outbound-api&ForwardedFrom=NianCat&CallerName=OwlMonkey&AccountSid=ACd03387e99bf959daa1e4810cc945708d&From=+13139202596&To=+16164601267`

var testListFixture = CallList{
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
	Calls: &[]Call{
		Call{
			ResponseCore: common.ResponseCore{
				AccountSid:   "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
				APIVersion:   "2010-04-01",
				DateCreated:  common.JSONTime{time.Date(2010, time.August, 16, 3, 45, 01, 0, &time.Location{})},
				DateUpdated:  common.JSONTime{time.Date(2010, time.August, 16, 3, 45, 03, 0, &time.Location{})},
				Direction:    "outbound-api",
				ErrorCode:    "",
				ErrorMessage: "",
				From:         "+16164601267",
				Sid:          "CAd88cd5b804dbcfb0ae6e0ccbbca67b83",
				Status:       "sent",
				To:           "+13139202596",
				URI:          "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls/CAd88cd5b804dbcfb0ae6e0ccbbca67b83.json",
			},
			Price: 0.0,
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
	"calls": [
		{
			"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
			"api_version": "2010-04-01",
			"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
			"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
			"direction": "outbound-api",
			"error_code": null,
			"error_message": null,
			"from": "+16164601267",
			"price": "-0.02000",
			"sid": "CAd88cd5b804dbcfb0ae6e0ccbbca67b83",
			"status": "complete",
			"to": "+13139202596",
			"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Calls/CAd88cd5b804dbcfb0ae6e0ccbbca67b83.json"
		}
	]
}`
