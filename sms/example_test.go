package sms_test

import (
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
)

func Example_sendSms() {
	act := twiliogo.NewAccount("AC1234567890abcdefghik1234567890ab", "token")
	act.Sms.Send(sms.Post{
		From: "+{Your twilio number}",
		To:   "+{Destination number}",
		Body: "Yo",
	})
}
