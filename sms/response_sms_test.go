package sms

import (
	"testing"
	"time"
)

var smsResponseFixture = `
{
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

func TestJsonEncode(t *testing.T) {
	var msg SmsResponse
	err := Unmarshal([]byte(smsResponseFixture), &msg)
	if err != nil {
		t.Errorf("Json failed to marshal with error => %s\n", err.Error())
	}

	if msg.DateSent != (time.Time{}) {
		t.Errorf("Unmarshal tried to assign a time to DateSent")
	}

	dateCreated, _ := time.Parse(twilioTimeFormat, "Wed, 18 Aug 2010 20:01:40 +0000")
	if msg.DateCreated.Second() != dateCreated.Second() { // take seconds to avoid diff in loc pointer addresses
		t.Errorf("Unmarshal improperly parsed DateCreated")
	}

	if msg.NumSegments != 1 || msg.NumMedia != 1 || msg.Price != 0.0 {
		t.Errorf("Unmarshal improperly converted strings to numbers")
	}
}
