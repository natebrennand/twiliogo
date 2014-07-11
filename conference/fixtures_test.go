package conference

var testFixtureString = `{
	"sid": "CFd03387e99bf959daa1e4810cc945708d",
	"friendly_name": "Go Team GoHelper",
	"status": "in-progress",
	"date_created": "Mon, 16 Aug 2010 03:45:01 +0000",
	"date_updated": "Mon, 16 Aug 2010 03:45:03 +0000",
	"account_sid": "ACd03387e99bf959daa1e4810cc945708d",
	"subresource_uris": {
		"participants": "/2010-04-01/Accounts/ACd03387e99bf959daa1e4810cc945708d/Conferences/CFd03387e99bf959daa1e4810cc945708d/Participants.json"
	}
	"uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39.json"
}`

var testFixtureListString = ` {
    "end": 49,
    "first_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences.json?Page=0&PageSize=50",
    "last_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences.json?Page=9&PageSize=50",
    "next_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences.json?Page=1&PageSize=50",
    "num_pages": 10,
    "page": 0,
    "page_size": 50,
    "previous_page_uri": null,
    "start": 0,
    "total": 462,
    "uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences.json",
    "conferences": [
        {
            "account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
            "api_version": "2010-04-01",
            "date_created": "Wed, 18 Aug 2010 20:20:06 +0000",
            "date_updated": "Wed, 18 Aug 2010 20:24:32 +0000",
            "friendly_name": "Party Line",
            "sid": "CFbbe46ff1274e283f7e3ac1df0072ab39",
            "status": "completed",
            "subresource_uris": {
                "participants": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants.json",
            },
            "uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39.json"
        }
    ]
}`

var testParticipantFixtureString = `{
    "account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
    "call_sid": "CA386025c9bf5d6052a1d1ea42b4d16662",
    "conference_sid": "CFbbe46ff1274e283f7e3ac1df0072ab39",
    "date_created": "Wed, 18 Aug 2010 20:20:10 +0000",
    "date_updated": "Wed, 18 Aug 2010 20:20:10 +0000",
    "end_conference_on_exit": true,
    "muted": false,
    "start_conference_on_enter": true,
    "uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants/CA386025c9bf5d6052a1d1ea42b4d16662.json"
}`

var testParticipantListFixtureString = ` {
    "num_pages": 1,
    "page": 0,
    "page_size": 50,
    "start": 0,
    "total": 2,
    "end": 1,
    "first_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants.json?Page=0&PageSize=50",
    "last_page_uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants.json?Page=0&PageSize=50",
    "next_page_uri": null,
    "previous_page_uri": null,
    "uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants.json",
    "participants": [
        {
            "account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
            "call_sid": "CA386025c9bf5d6052a1d1ea42b4d16662",
            "conference_sid": "CFbbe46ff1274e283f7e3ac1df0072ab39",
            "date_created": "Wed, 18 Aug 2010 20:20:10 +0000",
            "date_updated": "Wed, 18 Aug 2010 20:20:10 +0000",
            "end_conference_on_exit": true,
            "muted": false,
            "start_conference_on_enter": true,
            "uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants/CA386025c9bf5d6052a1d1ea42b4d16662.json"
        }
    ]
}`

var testParticipantMuteFixture = `{
    "account_sid": "AC5116d5d4df9f61ceae2f0732e1ea9f1b",
    "call_sid": "CA386025c9bf5d6052a1d1ea42b4d16662",
    "conference_sid": "CFbbe46ff1274e283f7e3ac1df0072ab39",
    "date_created": "Wed, 18 Aug 2010 20:20:10 +0000",
    "date_updated": "Wed, 18 Aug 2010 20:23:06 +0000",
    "end_conference_on_exit": true,
    "muted": true,
    "start_conference_on_enter": true,
    "uri": "/2010-04-01/Accounts/AC5116d5d4df9f61ceae2f0732e1ea9f1b/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants/CA386025c9bf5d6052a1d1ea42b4d16662.json"
}`
