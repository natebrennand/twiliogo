package media

import (
	"github.com/natebrennand/twiliogo/common"
	"time"
)

var (
	testMediaFixture = Media{
		Sid:         "ME85ebf7e12cb821f84b319340424dcb02",
		AccountSid:  "AC381707b751dbe4c74b15c5697ba67afd",
		ContentType: "image/png",
		DateCreated: common.JsonTime{time.Date(2013, time.September, 25, 22, 47, 40, 18, &time.Location{})},
		DateUpdated: common.JsonTime{time.Date(2013, time.September, 25, 22, 47, 40, 19, &time.Location{})},
		Uri:         "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json",
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
		/*
			PreviousPageUri: "",
			End:             1,
			Uri:             "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
			PageSize:        50,
			Start:           0,
			NextPageUri:     null,
			NumPages:        1,
			Total:           2,
			Page:            0,
			FirstPageUri:    "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
			LastPageUri:     "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media.json?PageSize=50&Page=0",
		*/
		MediaList: &[]Media{
			Media{
				Sid:         "ME85ebf7e12cb821f84b319340424dcb02",
				AccountSid:  "AC381707b751dbe4c74b15c5697ba67afd",
				ParentSid:   "MM800f449d0399ed014aae2bcc0cc2f2ec",
				ContentType: "image/png",
				DateCreated: common.JsonTime{time.Date(2013, time.September, 25, 22, 47, 40, 18, &time.Location{})},
				DateUpdated: common.JsonTime{time.Date(2013, time.September, 25, 22, 47, 40, 19, &time.Location{})},
				Uri:         "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json",
			},
			Media{
				Sid:         "ME8d8f717e2d6e5383055b3cd150ac5f54",
				AccountSid:  "AC381707b751dbe4c74b15c5697ba67afd",
				ParentSid:   "MM800f449d0399ed014aae2bcc0cc2f2ec",
				ContentType: "image/png",
				DateCreated: common.JsonTime{time.Date(2013, time.September, 25, 22, 47, 40, 18, &time.Location{})},
				DateUpdated: common.JsonTime{time.Date(2013, time.September, 25, 22, 47, 40, 19, &time.Location{})},
				Uri:         "/2010-04-01/Accounts/AC381707b751dbe4c74b15c5697ba67afd/Messages/MM800f449d0399ed014aae2bcc0cc2f2ec/Media/ME85ebf7e12cb821f84b319340424dcb02.json",
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
