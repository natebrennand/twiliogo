package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/applications"
)

func main() {
	act := twiliogo.NewAccountFromEnv()

	appList, err := act.Applications.List(applications.ListFilter{})
	if err != nil {
		fmt.Println("Error sending sms: ", err.Error())
	}

	fmt.Println("Current applications:")
	for _, app := range *appList.Applications {
		fmt.Printf("%#v\n", app)
	}

	fmt.Println("Waiting for sid to send request")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n') // wait to let server catching callback start
	name = strings.TrimRight(name, "\n")
	fmt.Printf("Creating a new app called %s\n", name)

	newApp, err := act.Applications.Create(applications.Application{FriendlyName: name})
	if err != nil {
		fmt.Println("Error creating app: ", err.Error())
	}
	fmt.Printf("%#v\n", newApp)
}
