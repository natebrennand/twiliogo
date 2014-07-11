package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/voice"
)

func main() {
	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Voice.List(voice.ListFilter{
		From: "+16162882901",
	})
	if err != nil {
		fmt.Println("Error getting call list: ", err.Error())
	}

	fmt.Printf("%#v\n", resp)
	for _, m := range *resp.Calls {
		fmt.Printf("%#v\n", m)
	}
}
