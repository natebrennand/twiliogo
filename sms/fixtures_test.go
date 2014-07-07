package sms

var (
	testNumber1                      = "+15558675309"
	testNumber2                      = "+14158141829"
	testSmsResponseFixtureAccountSid = "AC5ef8732a3c49700934481addd5ce1659"
)
var testSmsPostFixture = `{
	"body":"Jenny please?! I love you <3",
	"to":"+15558675309",
	"from":"+14158141829",
	"media_url":"http://www.example.com/hearts.png"
}`

var testSmsResponseFixture = `{
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
