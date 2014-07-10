package sms

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var (
	testNumber1                      = "+15558675309"
	testNumber2                      = "+14158141829"
	testSmsResponseFixtureAccountSid = "AC5ef8732a3c49700934481addd5ce1659"
	testSmsPostFixtureString         = `{
		"body":"Jenny please?! I love you <3",
		"to":"+15558675309",
		"from":"+14158141829",
		"media_url":"http://www.example.com/hearts.png"
	}`
	testSmsPostFixture = Post{
		Body:     "Jenny please?! I love you <3",
		To:       "+15558675309",
		From:     "+14158141829",
		MediaUrl: "http://www.example.com/hearts.png",
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
			ApiVersion:   "2010-04-01",
			Body:         "Jenny please?! I love you <3",
			Direction:    "outbound-api",
			ErrorCode:    "",
			ErrorMessage: "",
			From:         "+14158141829",
			Sid:          "MM90c6fc909d8504d45ecdb3a3d5b3556e",
			Status:       "queued",
			To:           "+15558675309",
			Uri:          "/2010-04-01/Accounts/AC5ef8732a3c49700934481addd5ce1659/Messages/MM90c6fc909d8504d45ecdb3a3d5b3556e.json",
		},
		DateCreated: common.JsonTime{time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
		DateUpdated: common.JsonTime{time.Date(2010, time.August, 18, 20, 1, 40, 0, &time.Location{})},
		DateSent:    common.JsonTime{time.Time{}},
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
		MediaList:     []Media{},
		StandardRequest: common.StandardRequest{
			AccountSid: "AC381707b751dbe4c74b15c5697ba67afd",
			From:       "+14248004123",
			To:         "+13605847116",
			Location:   nil,
		},
	}
	testSmsCallbackFixtureFormString        = `MessageSid=SMa2ff4e37c7cb43b49a820f2d7e3ee135&SmsSid=SMa2ff4e37c7cb43b49a820f2d7eee135&Body=Yo&MessageStatus=sent&AccountSid=AC381707b751dbe4c74b15c5697ba67afd&From=+14248004123&To=+13605847116`
	testSmsCallbackFixtureFormStringFailure = `MessageSid=SMa2ff4e37c7cb43b49a820f2d7e3ee135&SmsSid=SMa2ff4e37c7cb43b49a820f2d7eee135&Body=Yo&MessageStatus=sent&AccountSid=AC381707b751dbe4c74b15c5697ba67afd&From=+14248004123&To=+13605847116&NumMedia=a`
)
