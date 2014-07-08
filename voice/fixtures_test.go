package voice

var (
	testNumber1                   = "+15558675309"
	testNumber2                   = "+14158141829"
	testResponseFixtureAccountSid = "ACd03387e99bf959daa1e4810cc945708d"
)
var testPostFixture = `{
	"url":"http://twimlbin.com/558a498f",
	"to":"+15558675309",
	"from":"+14158141829"
}`

var testResponseFixture = `{
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
