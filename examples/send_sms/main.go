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
		From:           "+{ Your Twilio Number }",
		To:             "+{ Your Destination Number }",
		Body:           "Yo",
		StatusCallback: "{ Your callback endpoint }",
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
	http.ListenAndServe(":500", nil)
}
