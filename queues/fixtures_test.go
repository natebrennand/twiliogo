package queues

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var (
	testUpdateFixture = Update{
		MaxSize:      300,
		FriendlyName: "newname",
	}
	testQueueFixtureString = `{
	    "sid": "QU5ef8732a3c49700934481addd5ce1659",
	    "friendly_name": "persistent_queue1",
	    "current_size": 0,
	    "average_wait_time": 0,
	    "max_size": 10,
	    "date_created": "Mon, 26 Mar 2012 22:00:14 +0000",
	    "date_updated": "Mon, 26 Mar 2012 22:00:14 +0000",
	    "uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues/QU5ef8732a3c49700934481addd5ce1659.json"
	}`

	testQueueFixture = Queue{
		Sid:             "QU5ef8732a3c49700934481addd5ce1659",
		FriendlyName:    "persistent_queue1",
		CurrentSize:     0,
		AverageWaitTime: 0,
		MaxSize:         10,
		DateCreated:     common.JSONTime{Time: time.Date(2012, time.March, 26, 22, 0, 40, 14, &time.Location{})},
		DateUpdated:     common.JSONTime{Time: time.Date(2012, time.March, 26, 22, 0, 40, 14, &time.Location{})},
		URI:             "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues/QU5ef8732a3c49700934481addd5ce1659.json",
	}

	testQueueListFixtureString = `{ 
	    "start": 0,
	    "total": 2,
	    "uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues.json",
	    "end": 49,
	    "first_page_uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues.json?Page=0&PageSize=50",
	    "last_page_uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues.json?Page=0&PageSize=50",
	    "next_page_uri": "",
	    "num_pages": 0,
	    "page": 0,
	    "page_size": 50,
	    "previous_page_uri": "",
	    "queues": [
	        {
	            "average_wait_time": 0,
	            "current_size": 0,
	            "date_created": "Thu, 17 May 2012 23:50:39 +0000",
	            "date_updated": "Thu, 17 May 2012 23:50:39 +0000",
	            "friendly_name": "persistent_queue1",
	            "max_size": 1000,
	            "sid": "QU5ef8732a3c49700934481addd5ce1659",
	            "uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues/QU5ef8732a3c49700934481addd5ce1659.json"
	        },
	        { 
	            "average_wait_time": 0,
	            "current_size": 0,
	            "date_created": "Thu, 26 Apr 2012 20:12:45 +0000",
	            "date_updated": "Thu, 26 Apr 2012 20:12:45 +0000",
	            "friendly_name": "persistent_queue2",
	            "max_size": 100,
	            "sid": "QU5ef8732a3c49700934481addd5ce1660",
	            "uri": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues/QU5ef8732a3c49700934481addd5ce1660.json"
	        }
	    ]
	}`

	testQueueListFixture = QueueList{
		ListResponseCore: common.ListResponseCore{
			Start:           0,
			Total:           2,
			NumPages:        0,
			Page:            0,
			PageSize:        50,
			End:             49,
			URI:             "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues.json",
			FirstPageURI:    "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues.json?Page=0&PageSize=50",
			LastPageURI:     "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues.json?Page=0&PageSize=50",
			NextPageURI:     "",
			PreviousPageURI: "",
		},
		Queues: &[]Queue{
			Queue{
				Sid:             "QU5ef8732a3c49700934481addd5ce1659",
				FriendlyName:    "persistent_queue1",
				CurrentSize:     0,
				AverageWaitTime: 0,
				MaxSize:         1000,
				DateCreated:     common.JSONTime{Time: time.Date(2012, time.May, 17, 23, 50, 39, 0, &time.Location{})},
				DateUpdated:     common.JSONTime{Time: time.Date(2012, time.May, 17, 23, 50, 39, 0, &time.Location{})},
				URI:             "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues/QU5ef8732a3c49700934481addd5ce1659.json",
			},
			Queue{
				Sid:             "QU5ef8732a3c49700934481addd5ce1660",
				FriendlyName:    "persistent_queue2",
				CurrentSize:     0,
				AverageWaitTime: 0,
				MaxSize:         100,
				DateCreated:     common.JSONTime{Time: time.Date(2012, time.April, 26, 20, 12, 45, 0, &time.Location{})},
				DateUpdated:     common.JSONTime{Time: time.Date(2012, time.April, 26, 20, 12, 45, 0, &time.Location{})},
				URI:             "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Queues/QU5ef8732a3c49700934481addd5ce1660.json",
			},
		},
	}
)
