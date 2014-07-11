package main

import (
	"bufio"
	"fmt"
	"github.com/natebrennand/twiliogo"
	"os"
)

func main() {
	fmt.Println("Waiting for sid to send request")
	sid, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Sending a GET for {%s}\n", sid[0:34])

	act := twiliogo.NewAccountFromEnv()
	resp, err := act.Recordings.Get(sid[0:34])
	if err != nil {
		fmt.Println("Error getting recording: ", err.Error())
	}
	fmt.Printf("%#v\n", resp)
}
