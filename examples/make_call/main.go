package main

import (
	"bufio"
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/voice"
	"net/http"
	"os"
)

func makeCall() {
	fmt.Println("Waiting...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Here we gooo")
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Voice.Call(voice.Post{
		From: "+ {Your ###} ",
		To:   "+ {Their ###<3 }",
		Url:  "http://twimlbin.com/558a498f",
	})

	if err != nil {
		fmt.Println("Error making call: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
}

func printStatus(cbChan chan voice.Callback) {
	var cb voice.Callback
	for {
		cb = <-cbChan
		fmt.Printf("%#v\n", cb)
	}
}

func main() {
	go makeCall()
	cb := make(chan voice.Callback)
	go printStatus(cb)

	http.Handle("/", voice.CallbackHandler(cb))
	http.ListenAndServe("0.0.0.0:8000", nil)
}
