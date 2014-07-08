package main

import (
	"bufio"
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
	"net/http"
	"os"
)

func sendSms() {
	fmt.Println("Waiting for ENTER to send")
	bufio.NewReader(os.Stdin).ReadString('\n') // wait to let server catching callback start
	fmt.Println("Sending")
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Sms.Send(sms.Post{
		From: "+14248004123",
		To:   "+13605847116",
		Body: "Yo",
		// StatusCallback: "http://172.16.32.138:8001/",
		StatusCallback: "http://1c756787.ngrok.com/",
	})
	if err != nil {
		fmt.Println("Error sending sms: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
}

func printStatus(cbChan chan sms.Callback) {
	var cb sms.Callback
	for {
		cb = <-cbChan
		fmt.Printf("%#v\n", cb)
	}
}

func main() {
	go sendSms()
	cp := make(chan sms.Callback)
	go printStatus(cp)

	http.Handle("/", sms.CallbackHandler(cp))
	http.ListenAndServe("0.0.0.0:8001", nil)
}
