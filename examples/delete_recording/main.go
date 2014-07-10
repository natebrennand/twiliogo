package main

import (
	"bufio"
	"fmt"
	"github.com/natebrennand/twiliogo"
	"os"
)

func main() {
	fmt.Println("Waiting for sid to delete")
	sid, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Sending a DELETE for {%s}\n", sid[0:34])

	act := twiliogo.NewAccountFromEnv()
	err := act.Voice.Delete(sid[0:34])
	if err != nil {
		fmt.Println("Error getting recording: ", err.Error())
	}
}
