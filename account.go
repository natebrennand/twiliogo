package twiliogo

import (
	"fmt"
	"github.com/natebrennand/twiliogo/sms"
	"github.com/natebrennand/twiliogo/voice"
	"os"
	"regexp"
	"strings"
)

const (
	twilioAccount = "TWILIO_ACCOUNT"
	twilioToken   = "TWILIO_TOKEN"
)

type Account struct {
	AccountSid string
	Token      string
	Sms        sms.SmsAccount // redundancy for usability
	Voice      voice.VoiceAccount
}

func NewAccount(sid, token string) Account {
	match, err := regexp.MatchString("AC[a-z0-9]{32}", sid)
	if err != nil {
		panic("Error while checking AccountSid validity")
	}
	if match != true {
		panic("Invalid Account Sid")
	}
	return Account{
		AccountSid: sid,
		Token:      token,
		Sms:        sms.SmsAccount{sid, token},
		Voice:      voice.VoiceAccount{sid, token},
	}
}

func NewAccountFromEnv() Account {
	var sid, token string
	envVars := os.Environ()
	for _, x := range envVars {
		splits := strings.Split(x, "=")
		key, value := splits[0], splits[1]
		if key == twilioAccount {
			sid = value
		} else if key == twilioToken {
			token = value
		}
	}
	if sid == "" {
		panic(fmt.Sprintf("You must set the environment variable %s.", twilioAccount))
	}
	if token == "" {
		panic(fmt.Sprintf("You must set the environment variable %s.", twilioToken))
	}

	return NewAccount(sid, token)
}
