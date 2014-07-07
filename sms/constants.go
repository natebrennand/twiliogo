package sms

import (
	"time"
)

const (
	twilioTimeFormat = time.RFC1123Z
	postUrl          = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json"
)

var errorCode map[int]string = map[int]string{
	30001: "Queue Overflow",
	30002: "Account Suspended",
	30003: "Unreachable destination handset",
	30004: "Message blocked",
	30005: "Unknown destination handset",
	30006: "Landline or unreachable carrier",
	30007: "Carrier Violation",
	30008: "Unknown error",
}