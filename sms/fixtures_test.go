package sms

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var (
	/********************************************************
	*                     SMS
	********************************************************/
	testNumber1                      = "+15558675309"
	testNumber2                      = "+14158141829"
	testSmsResponseFixtureAccountSid = "AC5ef8732a3c49700934481addd5ce1659"
	testSmsPostFixtureString         = `{
		"body":"Jenny please?! I love you <3",
		"to":"+15558675309",
		"from":"+14158141829",
		"media_url":"http://www.example.com/hearts.png",
		"status_callback":"foobar.com",
		"application_sid":"AP5ef8732a3c49700934481addd5ce1659"
	}`
	testSmsPostFixture = Post{
		Body:           "Jenny please?! I love you <3",
		To:             "+15558675309",
		From:           "+14158141829",
		MediaURL:       "http://www.example.com/hearts.png",
		StatusCallback: "foobar.com",
		ApplicationSid: "AP5ef8732a3c49700934481addd5ce1659",
	}
	testSmsResponseFixtureString = `{
		"account_sid": "AC5ef8732a3c49700934481addd5ce1659",
		"api_version": "2010-04-01",
		"body": "Jenny please?! I love you <3",
		"num_segments": "1",
		"num_media": "1",
		"date_created": "Wed, 18 Aug 2010 20:01:40 +0000",
		"date_sent": null,
		"date_updated": "Wed, 18 Aug 2010 20:01:40 +0000",
		"direction": "outbound-api",
		"error_code": null,
		"error_message": null,
		"from": "+14158141829",
		"price": null,
		"sid": "MM90c6fc909d8504d45ecdb3a3d5b3556e",
		"status": "queued",
		"to": "+15558675309",
		"uri": "/2010-04-01/Accounts/AC5ef8732a3c49700934481addd5ce1659/Messages/MM90c6fc909d8504d45ecdb3a3d5b3556e.json"
	}`
	testSmsResponseFixture = Message{
		ResponseCore: common.ResponseCore{
			AccountSid:   "AC5ef8732a3c49700934481addd5ce1659",
			APIVersion:   "2010-04-01",
			Direction:    "outbound-api",
			ErrorCode:    "",
			ErrorMessage: "",
			From:         "+14158141829",
			Sid:          "MM90c6fc909d8504d45ecdb3a3d5b3556e",
			Status:       "queued",
			To:           "+15558675309",
			URI:          "/2010-04-01/Accounts/AC5ef8732a3c49700934481addd5ce1659/Messages/MM90c6fc909d8504d45ecdb3a3d5b3556e.json",
			DateCreated:  common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
			DateUpdated:  common.JSONTime{Time: time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
		},
		Body:        "Jenny please?! I love you <3",
		DateSent:    common.JSONTime{Time: time.Time{}},
		Price:       0.0,
		NumSegments: 1,
		NumMedia:    1,
	}
	testSmsCallbackFixture = Callback{
		MessageSid:    "SMa2ff4e37c7cb43b49a820f2d7e3ee135",
		SmsSid:        "SMa2ff4e37c7cb43b49a820f2d7e3ee135",
		Body:          "Yo",
		NumMedia:      0,
		MessageStatus: "sent",
		ErrorCode:     "",
		MediaList: []MediaReference{
			{
				ContentType: "XXX",
				URL:         "foobar.com",
			},
		},
		StandardRequest: common.StandardRequest{
			AccountSid: "AC381707b751dbe4c74b15c5697ba67afd",
			From:       "+14248004123",
			To:         "+13605847116",
			Location:   nil,
		},
	}
	testSmsCallbackFixtureFormString        = `MessageSid=SMa2ff4e37c7cb43b49a820f2d7e3ee135&SmsSid=SMa2ff4e37c7cb43b49a820f2d7eee135&Body=Yo&MessageStatus=sent&AccountSid=AC381707b751dbe4c74b15c5697ba67afd&From=+14248004123&To=+13605847116&NumMedia=1&MediaUrl1=foobar.com&MediaContentType1=XXX`
	testSmsCallbackFixtureFormStringFailure = `MessageSid=SMa2ff4e37c7cb43b49a820f2d7e3ee135&SmsSid=SMa2ff4e37c7cb43b49a820f2d7eee135&Body=Yo&MessageStatus=sent&AccountSid=AC381707b751dbe4c74b15c5697ba67afd&From=+14248004123&To=+13605847116&NumMedia=a`
	testSmsListFixture                      = MessageList{
		ListResponseCore: common.ListResponseCore{
			Start:           0,
			Total:           261,
			NumPages:        6,
			Page:            0,
			PageSize:        50,
			End:             49,
			URI:             "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json",
			FirstPageURI:    "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json?Page=0&PageSize=50",
			LastPageURI:     "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json?Page=5&PageSize=50",
			NextPageURI:     "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json?Page=1&PageSize=50",
			PreviousPageURI: "",
		},
		Messages: &[]Message{
			Message{
				ResponseCore: common.ResponseCore{
					AccountSid:   "AC381707b751dbe4c74b15c5697ba67afd",
					APIVersion:   "2010-04-01",
					DateCreated:  common.JSONTime{Time: time.Date(2010, time.August, 16, 3, 45, 01, 0, &time.Location{})},
					DateUpdated:  common.JSONTime{Time: time.Date(2010, time.August, 16, 3, 45, 03, 0, &time.Location{})},
					Direction:    "outbound-api",
					ErrorCode:    "",
					ErrorMessage: "",
					From:         "+14158141829",
					Sid:          "SM800f449d0399ed014aae2bcc0cc2f2ec",
					Status:       "sent",
					To:           "+15558675309",
					URI:          "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec.json",
				},
				Body:        "Hey Jenny why aren't you returning my calls?",
				DateSent:    common.JSONTime{Time: time.Date(2010, time.August, 16, 3, 45, 03, 0, &time.Location{})},
				NumSegments: 1,
				NumMedia:    0,
				Price:       -0.02000,
			},
		},
	}
	testSmsListFixtureString = `{
	"start": 0,
	"total": 261,
	"num_pages": 6,
	"page": 0,
	"page_size": 50,
	"end": 49,
	"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json",
	"first_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json?Page=0&PageSize=50",
	"last_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json?Page=5&PageSize=50",
	"next_page_uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages.json?Page=1&PageSize=50",
	"previous_page_uri": null,
	"messages": [
		{
			"account_sid": "AC381707b751dbe4c74b15c5697ba67afd",
			"api_version": "2010-04-01",
			"body": "Hey Jenny why aren't you returning my calls?",
			"num_segments": "1",
			"num_media": "0",
			"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
			"date_sent": "Mon, 16 Aug 2010 03:45:03 +0000",
			"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
			"direction": "outbound-api",
			"error_code": null,
			"error_message": null,
			"from": "+14158141829",
			"price": "-0.02000",
			"sid": "SM800f449d0399ed014aae2bcc0cc2f2ec",
			"status": "sent",
			"to": "+15558675309",
			"uri": "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec.json"
		}
	]
}`
	/********************************************************
	*                     MEDIA
	********************************************************/
	testMediaFixture = Media{
		ResourceInfo: common.ResourceInfo{
			Sid:         "ME85ebf7e12cb821f84b319340424dcb02",
			AccountSid:  "AC381707b751dbe4c74b15c5697ba67afd",
			DateCreated: common.JSONTime{Time: time.Date(2013, time.September, 25, 22, 47, 40, 18, &time.Location{})},
			DateUpdated: common.JSONTime{Time: time.Date(2013, time.September, 25, 22, 47, 40, 19, &time.Location{})},
			URI:         "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json",
		},
		ContentType: "image/png",
	}
	testMediaFixtureString = `{
	"sid":"ME85ebf7e12cb821f84b319340424dcb02",
	"account_sid":"AC381707b751dbe4c74b15c5697ba67afd",
	"parent_sid":"MM800f449d0399ed014aae2bcc0cc2f2ec",
	"content_type":"image/png",
	"date_created":"Wed, 25 Sep 2013 22:47:18 +0000",
	"date_updated":"Wed, 25 Sep 2013 22:47:19 +0000",
	"uri":"/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json"
}`
	testMediaListFixture = MediaList{
		ListResponseCore: common.ListResponseCore{
			PreviousPageURI: "",
			End:             1,
			URI:             "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
			PageSize:        50,
			Start:           0,
			NextPageURI:     "",
			NumPages:        1,
			Total:           2,
			Page:            0,
			FirstPageURI:    "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
			LastPageURI:     "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
		},
		MediaList: &[]Media{
			Media{
				ParentSid:   "MM800f449d0399ed014aae2bcc0cc2f2ec",
				ContentType: "image/png",
				ResourceInfo: common.ResourceInfo{
					Sid:         "ME85ebf7e12cb821f84b319340424dcb02",
					AccountSid:  "AC381707b751dbe4c74b15c5697ba67afd",
					DateCreated: common.JSONTime{Time: time.Date(2013, time.September, 25, 22, 47, 40, 18, &time.Location{})},
					DateUpdated: common.JSONTime{Time: time.Date(2013, time.September, 25, 22, 47, 40, 19, &time.Location{})},
					URI:         "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json",
				},
			},
			Media{
				ParentSid:   "MM800f449d0399ed014aae2bcc0cc2f2ec",
				ContentType: "image/png",
				ResourceInfo: common.ResourceInfo{
					Sid:         "ME8d8f717e2d6e5383055b3cd150ac5f54",
					AccountSid:  "AC381707b751dbe4c74b15c5697ba67afd",
					DateCreated: common.JSONTime{Time: time.Date(2013, time.September, 25, 22, 47, 40, 18, &time.Location{})},
					DateUpdated: common.JSONTime{Time: time.Date(2013, time.September, 25, 22, 47, 40, 19, &time.Location{})},
					URI:         "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json",
				},
			},
		},
	}
	testMediaListFixtureString = `{
	"first_page_uri":"/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
	"last_page_uri":"/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
	"media_list":[
		{
			"sid":"ME85ebf7e12cb821f84b319340424dcb02",
			"account_sid":"AC381707b751dbe4c74b15c5697ba67afd",
			"parent_sid":"MM800f449d0399ed014aae2bcc0cc2f2ec",
			"content_type":"image/png",
			"date_created":"Wed, 25 Sep 2013 22:47:18 +0000",
			"date_updated":"Wed, 25 Sep 2013 22:47:19 +0000",
			"uri":"/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json"
		},
		{
			"sid":"ME8d8f717e2d6e5383055b3cd150ac5f54",
			"account_sid":"AC381707b751dbe4c74b15c5697ba67afd",
			"parent_sid":"MM800f449d0399ed014aae2bcc0cc2f2ec",
			"content_type":"image/png",
			"date_created":"Wed, 25 Sep 2013 22:47:18 +0000",
			"date_updated":"Wed, 25 Sep 2013 22:47:19 +0000",
			"uri":"/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME8d8f717e2d6e5383055b3cd150ac5f54.json"
		}
	],
	"previous_page_uri":null,
	"end":1,
	"uri":"/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
	"page_size":50,
	"start":0,
	"next_page_uri":null,
	"num_pages":1,
	"total":2,
	"page":0
}`
)
