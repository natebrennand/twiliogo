package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
)

func main() {
	act := twiliogo.NewAccountFromEnv()

	fmt.Println("Sending MMS")
	msg, err := act.Sms.Send(sms.Post{
		From:     "+1{ Your Twilio MMS enabled number }",
		To:       "+1{ Destinatino number }",
		Body:     "Yo",
		MediaURL: "http://i.imgur.com/XcMxly3.jpg", // Cute otter
	})
	if err != nil {
		fmt.Println("Error sending mms: ", err.Error())
		return
	}
	fmt.Printf("%#v\n", msg)

	fmt.Println("Getting MMS media list")
	mediaList, err := act.Sms.GetMediaList(msg.Sid)
	if err != nil {
		fmt.Println("Error getting mms: ", err.Error())
		return
	}
	fmt.Printf("%#v\n", *mediaList.MediaList)
}
