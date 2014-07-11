package main

import (
	"bufio"
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/notifications"
	"os"
)

func main() {
	act := twiliogo.NewAccountFromEnv()

	notificationList, err := act.Notifications.List(notifications.Filter{})
	if err != nil {
		fmt.Println("Errored while getting notifications")
		fmt.Println(err.Error())
		return
	}
	for _, n := range *notificationList.Notifications {
		fmt.Printf("%#v\n", n)
	}

	fmt.Println("Enter a sid to delete a notifcation")
	sid, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if len(sid) != 35 {
		fmt.Println("Exiting")
		return
	}
	fmt.Printf("Sending a DELETE for {%s}\n", sid[0:34])

	err = act.Notifications.Delete(sid[0:34])
	if err != nil {
		fmt.Println("Error deleting notification: ", err.Error())
		return
	}
	fmt.Printf("{%s} deleted\n", sid[0:34])
}
