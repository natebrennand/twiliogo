package recording

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var testResponseFixtureString = `{
	"account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
	"api_version": "2010-04-01",
	"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
	"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
	"direction": "outbound-api",
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Calls/CAd88cd5b804dbcfb0ae6e0ccbbca67b83.json",
	"sid": "REc8009ae243230394bc84437e1a9f4650",
	"call_sid": "CAd88cd5b804dbcfb0ae6e0ccbbca67b83",
	"duration": "45"
}`

var testResponseFixture = Recording{
	Sid:         "REc8009ae243230394bc84437e1a9f4650",
	DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
	DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
	AccountSid:  "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
	CallSid:     "CAd88cd5b804dbcfb0ae6e0ccbbca67b83",
	Duration:    0,
	APIVersion:  "2010-04-01",
	URI:         "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls/CAd88cd5b804dbcfb0ae6e0ccbbca67b83.json",
}

var testListFixture = RecordingList{
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
	Recordings: &[]Recording{
		Recording{
			Sid:         "REc8009ae243230394bc84437e1a9f4650",
			DateCreated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			AccountSid:  "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
			CallSid:     "CAd88cd5b804dbcfb0ae6e0ccbbca67b83",
			Duration:    0,
			APIVersion:  "2010-04-01",
			URI:         "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Calls/CAd88cd5b804dbcfb0ae6e0ccbbca67b83.json",
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
