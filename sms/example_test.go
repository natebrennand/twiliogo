package sms_test

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
)

func Example_sendSms() {
	act := twiliogo.NewAccount("AC1234567890abcdefghik1234567890ab", "token")
	resp, err := act.Sms.Send(sms.Post{
		From: "+{Your twilio number}",
		To:   "+{Destination number}",
		Body: "Yo",
	})
	if err != nil {
	}
	fmt.Println(resp.Status)
	// Output:
}
