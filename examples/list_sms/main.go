package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
)

func main() {
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Sms.List(sms.Filter{
		From: "+{ Your phone # }",
	})
	if err != nil {
		fmt.Println("Error sending sms: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
	for _, m := range *resp.Messages {
		fmt.Printf("%#v\n", m)
	}
}
