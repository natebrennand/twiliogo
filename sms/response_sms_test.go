package sms

import (
	"github.com/natebrennand/twiliogo/common"
	"testing"
	"time"
)

func TestJsonEncodeSuccessful(t *testing.T) {
	var msg Response
	err := Unmarshal([]byte(testSmsResponseFixtureString), &msg)
	if err != nil {
		t.Errorf("Json failed to marshal with error => %s\n", err.Error())
	}

	if msg.DateSent != (common.JsonTime{time.Time{}}) {
		t.Errorf("Unmarshal tried to assign a time to DateSent")
	}

	dateCreated, _ := time.Parse(twilioTimeFormat, "Wed, 18 Aug 2010 20:01:40 +0000")
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
