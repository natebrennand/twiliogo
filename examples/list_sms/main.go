package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
)

func main() {
	act := twiliogo.NewAccountFromEnv()
	messageList, err := act.Sms.List(sms.Filter{})
	if err != nil {
		fmt.Println("Error retrieving sms: ", err.Error())
		return
	}
	fmt.Printf("%#v\n", messageList)
	for _, m := range *messageList.Messages {
		fmt.Println(m.Sid)
	}

	for {
		// keep trying to get the next page of results
		err = messageList.Next()
		if err != nil {
			fmt.Println("Error retrieving sms: ", err.Error())
			return
		}
		fmt.Printf("%#v\n", messageList)
		for _, m := range *messageList.Messages {
			fmt.Println(m.Sid)
		}
	}
}
