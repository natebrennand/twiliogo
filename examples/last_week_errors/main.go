package main

import (
	"fmt"
	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/common"
	"github.com/natebrennand/twiliogo/notifications"
	"time"
)

func main() {
	act := twiliogo.NewAccountFromEnv()

	afterDate := time.Now().AddDate(0, 0, -7)
	notificationList, err := act.Notifications.List(notifications.Filter{
		AfterMessageDate: &afterDate,
	})
	if err != nil {
		fmt.Println("Errored while getting notifications")
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("All %d notifications from after %s:\n", len(*notificationList.Notifications), afterDate.Format(common.GMTTimeLayout))
	for _, n := range *notificationList.Notifications {
		fmt.Printf("%#v\n", n)
	}
}
