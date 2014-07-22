package twiliogo

import (
	"fmt"
	"github.com/natebrennand/twiliogo/applications"
	"github.com/natebrennand/twiliogo/common"
	"github.com/natebrennand/twiliogo/conference"
	"github.com/natebrennand/twiliogo/notifications"
	"github.com/natebrennand/twiliogo/numbers"
	"github.com/natebrennand/twiliogo/recording"
	"github.com/natebrennand/twiliogo/sms"
	"github.com/natebrennand/twiliogo/sms/shortcodes"
	"github.com/natebrennand/twiliogo/transcription"
	"github.com/natebrennand/twiliogo/usage"
	"github.com/natebrennand/twiliogo/voice"
	"net/http"
	"os"
	"regexp"
)

const (
	twilioAccount = "TWILIO_ACCOUNT"
	twilioToken   = "TWILIO_TOKEN"
)

var validateAccountSid = regexp.MustCompile("^AC[a-z0-9]{32}$").MatchString

// Account is a catch-all account object that holds references to all Twilio resources.
type Account struct {
	AccountSid string
	Token      string
	// redundancy for usability
	Sms            sms.Account
	ShortCode      shortcodes.Account
	Voice          voice.Account
	Recordings     recording.Account
	Transcriptions transcription.Account
	Conferences    conference.Account
	Applications   applications.Account
	Notifications  notifications.Account
	Usage          usage.Account
	Numbers        numbers.Account
}

// NewAccount builds an Account resource from a sid & token.
func NewAccount(sid, token string) Account {
	if !validateAccountSid(sid) {
		panic("Invalid Account Sid")
	}

	a := common.Account{
		AccountSid: sid,
		Token:      token,
		Client:     http.Client{},
	}

	return Account{
		AccountSid:     sid,
		Token:          token,
		Sms:            sms.Account{Account: a},
		ShortCode:      shortcodes.Account{Account: a},
		Voice:          voice.Account{Account: a},
		Usage:          usage.Account{Account: a},
		Recordings:     recording.Account{Account: a},
		Transcriptions: transcription.Account{Account: a},
		Applications:   applications.Account{Account: a},
		Conferences:    conference.Account{Account: a},
		Notifications:  notifications.Account{Account: a},
		Numbers:        numbers.Account{Account: a},
	}
}

// NewAccountFromEnv builds an Account resource from environment variables.
func NewAccountFromEnv() Account {
	sid := os.Getenv(twilioAccount)
	token := os.Getenv(twilioToken)
	if sid == "" {
		panic(fmt.Sprintf("You must set the environment variable %s.", twilioAccount))
	}
	if token == "" {
		panic(fmt.Sprintf("You must set the environment variable %s.", twilioToken))
	}

	return NewAccount(sid, token)
}
