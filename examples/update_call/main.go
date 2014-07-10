package main

import (
	"bufio"
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/voice"
	"os"
)

func makeCall() string {
	fmt.Println("Waiting...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Here we gooo")
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Voice.Call(voice.Post{
		From: "+16162882901",
		To:   "+16164601267",
		URL:  "http://twimlbin.com/558a498f",
	})

	if err != nil {
		fmt.Println("Error making call: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
	return resp.Sid
}

func updateCall(sid string) {
	act := twiliogo.NewAccountFromEnv()
	fmt.Println("Send an update?")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Posting update for ", sid)

	resp, err := act.Voice.Update(voice.Update{
		URL:    "http://twimlbin.com/701ed8a7",
		Method: "POST",
	}, sid)

	if err != nil {
		fmt.Println("Error making call: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
}

func main() {
	sid := makeCall()
	updateCall(sid)
}
