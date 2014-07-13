package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
)

func main() {

	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Transcriptions.List()
	if err != nil {
		fmt.Println("Error getting recording: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
	for _, t := range *resp.Transcriptions {
		fmt.Printf("%#v\n", t)
	}
}
