package common

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPriceDecode(t *testing.T) {
	var price *JsonPrice
	price.UnmarshalJSON([]byte("null"))
	if price != nil {
		t.Error("Price should be set to nil on null")
	}
}

func TestJsonDecode(t *testing.T) {
	var msg testMessageWithDate
	err := json.Unmarshal([]byte(testSmsResponseFixtureString), &msg)
	if err != nil {
		t.Errorf("Json failed to marshal with error => %s\n", err.Error())
	}

	if msg.DateSent != (JsonTime{time.Time{}}) {
		t.Errorf("Unmarshal tried to assign a time to DateSent")
	}

	dateCreated, _ := time.Parse(TwilioTimeFormat, "Wed, 18 Aug 2010 20:01:40 +0000")
	if msg.DateCreated.Second() != dateCreated.Second() { // take seconds to avoid diff in loc pointer addresses
		t.Errorf("Unmarshal improperly parsed DateCreated")
	}

	if msg.NumSegments != 1 || msg.NumMedia != 1 || msg.Price != 0.0 {
		t.Errorf("Unmarshal improperly converted strings to numbers")
	}

	if msg.AccountSid != testSmsResponseFixtureAccountSid {
		t.Errorf("Improperly parsed AccountSid")
	}
}
