package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
)

func main() {
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Sms.Send(sms.Post{
		From: "+{Your twilio number}",
		To:   "+{Destination number}",
		Body: "Yo",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", resp)
}
