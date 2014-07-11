package common

import (
	"encoding/json"
	"testing"
	"time"
)

type testPriceStruct struct {
	Price *JSONFloat `json:"price"`
}

func TestPriceDecode(t *testing.T) {
	var p JSONFloat = 0.01
	var priceHolder testPriceStruct
	err := json.Unmarshal([]byte(`{"price":null}`), &priceHolder)
	if err != nil || priceHolder.Price != nil {
		t.Error("Price should be set to nil on null")
	}

	err = json.Unmarshal([]byte(`{"price":0.01}`), &priceHolder)
	if err != nil {
		t.Fatalf("Error decoding => %s", err.Error())
	} else if priceHolder.Price == nil {
		t.Fatal("Price should be set")
	} else if *(priceHolder.Price) != p {
		t.Errorf("Price should be set to $.01, %f", *(priceHolder.Price))
	}
}

func TestJSONDecode(t *testing.T) {
	var msg testMessageWithDate
	err := json.Unmarshal([]byte(testSmsResponseFixtureString), &msg)
	if err != nil {
		t.Errorf("JSON failed to marshal with error => %s\n", err.Error())
	}

	if msg.DateSent != (JSONTime{time.Time{}}) {
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
